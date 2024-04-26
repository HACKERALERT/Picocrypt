package main

import (
	"archive/zip"
	"bytes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/subtle"
	"flag"
	"fmt"
	"hash"
	"io"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/HACKERALERT/infectious"
	"github.com/HACKERALERT/serpent"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/chacha20"
	"golang.org/x/crypto/hkdf"
	"golang.org/x/crypto/sha3"
	"golang.org/x/term"
)

var KiB = 1 << 10
var MiB = 1 << 20
var GiB = 1 << 30
var TiB = 1 << 40
var rs5, _ = infectious.NewFEC(5, 15)
var rs16, _ = infectious.NewFEC(16, 48)
var rs24, _ = infectious.NewFEC(24, 72)
var rs32, _ = infectious.NewFEC(32, 96)
var rs64, _ = infectious.NewFEC(64, 192)
var rs128, _ = infectious.NewFEC(128, 136)

func rsEncode(rs *infectious.FEC, data []byte) []byte {
	res := make([]byte, rs.Total())
	rs.Encode(data, func(s infectious.Share) {
		res[s.Number] = s.Data[0]
	})
	return res
}

func rsDecode(rs *infectious.FEC, data []byte, fast bool) ([]byte, error) {
	if rs.Total() == 136 && fast {
		return data[:128], nil
	}
	tmp := make([]infectious.Share, rs.Total())
	for i := 0; i < rs.Total(); i++ {
		tmp[i].Number = i
		tmp[i].Data = append(tmp[i].Data, data[i])
	}
	res, err := rs.Decode(nil, tmp)
	if err != nil {
		if rs.Total() == 136 {
			return data[:128], err
		}
		return data[:rs.Total()/3], err
	}
	return res, nil
}

func pad(data []byte) []byte {
	padLen := 128 - len(data)%128
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, padding...)
}

func unpad(data []byte) []byte {
	padLen := int(data[127])
	return data[:128-padLen]
}

func work() int {
	flag.Usage = func() {
		fmt.Println("Usage: picocrypt [-p]aranoid [-r]eedsolo <item1> [<item2> ...]")
		fmt.Println("Items: can be a file (cat.png), folder (./src), or glob (*.txt)")
	}
	paranoid := flag.Bool("p", false, "")
	reedsolo := flag.Bool("r", false, "")
	fix := flag.Bool("fix", false, "")
	flag.Parse()

	mode := ""
	if flag.NArg() == 0 {
		flag.Usage()
		return 0
	}
	if flag.NArg() == 1 {
		if strings.HasSuffix(flag.Arg(0), ".pcv") {
			mode = "decrypt"
		} else {
			mode = "encrypt"
		}
	} else {
		mode = "encrypt"
		for _, v := range flag.Args() {
			if strings.HasSuffix(v, ".pcv") {
				fmt.Println("Multiple items cannot contain volumes.")
				return 1
			}
		}
	}
	for _, v := range flag.Args() {
		if v == "-p" || v == "-r" || v == "-fix" {
			fmt.Println("Flags are only accepted before arguments!")
			return 1
		}
	}

	var password, cpassword []byte
	var err error
	if mode == "encrypt" {
		fmt.Print("Password: ")
		password, err = term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading password.")
			return 1
		}
		fmt.Print(strings.Repeat("*", len(password)), " | Confirm: ")
		cpassword, err = term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading password.")
			return 1
		}
		fmt.Println(strings.Repeat("*", len(cpassword)))
		if !bytes.Equal(password, cpassword) {
			fmt.Println("Passwords don't match.")
			return 1
		}
	} else {
		fmt.Print("Password: ")
		password, err = term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading password.")
			return 1
		}
		fmt.Println(strings.Repeat("*", len(password)))
	}

	fin_, fout_ := "", ""
	if mode == "decrypt" {
		fin_ = flag.Arg(0)
		fout_ = strings.TrimSuffix(fin_, ".pcv")
	} else {
		stat, err := os.Stat(flag.Arg(0))
		if flag.NArg() == 1 && err == nil && !stat.IsDir() {
			fin_ = flag.Arg(0)
			fout_ = fin_ + ".pcv"
		} else {
			items := []string{}
			for _, v := range flag.Args() {
				if strings.Contains(v, "../") || strings.HasPrefix(v, "/") {
					fmt.Println("Cannot encrypt outside of current directory.")
					return 1
				}
				matches, err := filepath.Glob(v)
				if err != nil {
					fmt.Println("Invalid glob pattern(s).")
					return 1
				}
				items = append(items, matches...)
			}
			files := []string{}
			for _, v := range items {
				stat, err := os.Stat(v)
				if err != nil {
					fmt.Println("Cannot access input(s).")
					return 1
				}
				if !stat.IsDir() {
					files = append(files, v)
				} else {
					filepath.Walk(v, func(path string, _ os.FileInfo, _ error) error {
						stat, err := os.Stat(path)
						if err == nil && !stat.IsDir() {
							files = append(files, path)
						}
						return nil
					})
				}
			}
			if len(files) == 0 {
				fmt.Println("Nothing to encrypt.")
				return 1
			}

			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("Cannot get current working directory.")
				return 1
			}
			file, err := os.CreateTemp("", "")
			if err != nil {
				fmt.Println("Cannot create temporary file.")
				return 1
			}
			writer := zip.NewWriter(file)
			for i, path := range files {
				stat, err := os.Stat(path)
				if err != nil {
					continue
				}
				header, err := zip.FileInfoHeader(stat)
				if err != nil {
					continue
				}
				abs, err := filepath.Abs(path)
				if err != nil {
					continue
				}
				abs = filepath.ToSlash(abs)
				header.Name = strings.TrimPrefix(abs, filepath.ToSlash(dir))
				header.Name = strings.TrimPrefix(header.Name, "/")
				header.Method = zip.Deflate
				entry, err := writer.CreateHeader(header)
				if err != nil {
					continue
				}
				fin, err := os.Open(path)
				if err != nil {
					writer.Close()
					file.Close()
					os.Remove(file.Name())
					fmt.Println("Read access to input(s) denied.")
					return 1
				}
				bar := progressbar.NewOptions(
					int(stat.Size()),
					progressbar.OptionClearOnFinish(),
					progressbar.OptionFullWidth(),
					progressbar.OptionShowBytes(true),
					progressbar.OptionUseIECUnits(true),
					progressbar.OptionSetDescription(
						fmt.Sprintf("Compressing [%d/%d]:", i+1, len(files)),
					),
				)
				_, err = io.Copy(io.MultiWriter(entry, bar), fin)
				fin.Close()
				if err != nil {
					writer.Close()
					file.Close()
					os.Remove(file.Name())
					fmt.Println("Insufficient disk space.")
					return 1
				}
			}
			writer.Close()
			file.Close()
			fin_ = file.Name()
			fout_ = "encrypted.zip.pcv"
			defer os.Remove(file.Name())
		}
	}

	var padded bool
	var salt []byte       // Argon2 salt, 16 bytes
	var hkdfSalt []byte   // HKDF-SHA3 salt, 32 bytes
	var serpentIV []byte  // Serpent IV, 16 bytes
	var nonce []byte      // 24-byte XChaCha20 nonce
	var keyHash []byte    // SHA3-512 hash of encryption key
	var keyHashRef []byte // Same as 'keyHash', but used for comparison
	var authTag []byte    // 64-byte authentication tag (BLAKE2b or HMAC-SHA3)

	fin, err := os.Open(fin_)
	if err != nil {
		fmt.Println("Error accessing input file.")
		return 1
	}
	_, err = os.Stat(fout_)
	if err == nil {
		fmt.Println("Output file already exists.")
		return 1
	}
	fout, err := os.Create(fout_)
	if err != nil {
		fmt.Println("Error creating output file.")
		return 1
	}

	stat, err := os.Stat(fin_)
	if err != nil {
		fmt.Println("Error accessing input file.")
		return 1
	}
	total := stat.Size()
	if mode == "decrypt" {
		total -= 789
	}
	if mode == "encrypt" {
		errs := make([]error, 11)
		salt = make([]byte, 16)
		hkdfSalt = make([]byte, 32)
		serpentIV = make([]byte, 16)
		nonce = make([]byte, 24)
		_, errs[0] = fout.Write(rsEncode(rs5, []byte("v1.34")))
		commentsLength := []byte(fmt.Sprintf("%05d", 0))
		_, errs[1] = fout.Write(rsEncode(rs5, commentsLength))
		flags := make([]byte, 5)
		if *paranoid {
			flags[0] = 1
		}
		if *reedsolo {
			flags[3] = 1
		}
		if total%int64(MiB) >= int64(MiB)-128 {
			flags[4] = 1
		}
		_, errs[3] = fout.Write(rsEncode(rs5, flags))
		rand.Read(salt)
		rand.Read(hkdfSalt)
		rand.Read(serpentIV)
		rand.Read(nonce)
		_, errs[4] = fout.Write(rsEncode(rs16, salt))
		_, errs[5] = fout.Write(rsEncode(rs32, hkdfSalt))
		_, errs[6] = fout.Write(rsEncode(rs16, serpentIV))
		_, errs[7] = fout.Write(rsEncode(rs24, nonce))
		_, errs[8] = fout.Write(make([]byte, 192))
		_, errs[9] = fout.Write(make([]byte, 96))
		_, errs[10] = fout.Write(make([]byte, 192))
		for _, err := range errs {
			if err != nil {
				fin.Close()
				fout.Close()
				os.Remove(fout_)
				fmt.Println("Insufficient disk space.")
				return 1
			}
		}
	} else {
		errs := make([]error, 10)
		version := make([]byte, 15)
		fin.Read(version)
		_, errs[0] = rsDecode(rs5, version, !(*fix))
		tmp := make([]byte, 15)
		fin.Read(tmp)
		tmp, errs[1] = rsDecode(rs5, tmp, !(*fix))
		commentsLength, _ := strconv.Atoi(string(tmp))
		fin.Read(make([]byte, commentsLength*3))
		total -= int64(commentsLength) * 3
		flags := make([]byte, 15)
		fin.Read(flags)
		flags, errs[2] = rsDecode(rs5, flags, !(*fix))
		*paranoid = flags[0] == 1
		*reedsolo = flags[3] == 1
		padded = flags[4] == 1
		if flags[1] == 1 {
			fin.Close()
			fout.Close()
			os.Remove(fout_)
			fmt.Println("Keyfiles are not supported.")
			return 1
		}
		salt = make([]byte, 48)
		fin.Read(salt)
		salt, errs[3] = rsDecode(rs16, salt, !(*fix))
		hkdfSalt = make([]byte, 96)
		fin.Read(hkdfSalt)
		hkdfSalt, errs[4] = rsDecode(rs32, hkdfSalt, !(*fix))
		serpentIV = make([]byte, 48)
		fin.Read(serpentIV)
		serpentIV, errs[5] = rsDecode(rs16, serpentIV, !(*fix))
		nonce = make([]byte, 72)
		fin.Read(nonce)
		nonce, errs[6] = rsDecode(rs24, nonce, !(*fix))
		keyHashRef = make([]byte, 192)
		fin.Read(keyHashRef)
		keyHashRef, errs[7] = rsDecode(rs64, keyHashRef, !(*fix))
		keyfileHashRef := make([]byte, 96)
		fin.Read(keyfileHashRef)
		_, errs[8] = rsDecode(rs32, keyfileHashRef, !(*fix))
		authTag = make([]byte, 192)
		fin.Read(authTag)
		authTag, errs[9] = rsDecode(rs64, authTag, !(*fix))
		for _, err := range errs {
			if err != nil {
				fin.Close()
				fout.Close()
				os.Remove(fout_)
				fmt.Println("The volume header is damaged.")
				return 1
			}
		}
	}

	var key []byte
	if *paranoid {
		key = argon2.IDKey(
			password,
			salt,
			8,
			1<<20,
			8,
			32,
		)
	} else {
		key = argon2.IDKey(
			password,
			salt,
			4,
			1<<20,
			4,
			32,
		)
	}

	tmp := sha3.New512()
	tmp.Write(key)
	keyHash = tmp.Sum(nil)

	if mode == "decrypt" {
		if subtle.ConstantTimeCompare(keyHash, keyHashRef) != 1 {
			fin.Close()
			fout.Close()
			os.Remove(fout_)
			fmt.Println("Incorrect password.")
			return 1
		}
	}

	done, counter := 0, 0
	chacha, _ := chacha20.NewUnauthenticatedCipher(key, nonce)
	var mac hash.Hash
	subkey := make([]byte, 32)
	hkdf := hkdf.New(sha3.New256, key, hkdfSalt, nil)
	hkdf.Read(subkey)
	if *paranoid {
		mac = hmac.New(sha3.New512, subkey)
	} else {
		mac, _ = blake2b.New512(subkey)
	}
	serpentKey := make([]byte, 32)
	hkdf.Read(serpentKey)
	s, _ := serpent.NewCipher(serpentKey)
	serpent := cipher.NewCTR(s, serpentIV)

	bar := progressbar.NewOptions(
		int(total),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionFullWidth(),
		progressbar.OptionShowBytes(true),
		progressbar.OptionUseIECUnits(true),
		progressbar.OptionSetDescription(
			(func() string {
				if mode == "encrypt" {
					return "Encrypting:"
				}
				return "Decrypting:"
			})(),
		),
	)
	for {
		var src []byte
		if mode == "decrypt" && *reedsolo {
			src = make([]byte, MiB/128*136)
		} else {
			src = make([]byte, MiB)
		}
		size, err := fin.Read(src)
		if err != nil {
			break
		}
		src = src[:size]
		dst := make([]byte, len(src))

		if mode == "encrypt" {
			if *paranoid {
				serpent.XORKeyStream(dst, src)
				copy(src, dst)
			}
			chacha.XORKeyStream(dst, src)
			mac.Write(dst)
			if *reedsolo {
				copy(src, dst)
				dst = nil
				if len(src) == MiB {
					for i := 0; i < MiB; i += 128 {
						dst = append(dst, rsEncode(rs128, src[i:i+128])...)
					}
				} else {
					chunks := math.Floor(float64(len(src)) / 128)
					for i := 0; float64(i) < chunks; i++ {
						dst = append(dst, rsEncode(rs128, src[i*128:(i+1)*128])...)
					}
					dst = append(dst, rsEncode(rs128, pad(src[int(chunks*128):]))...)
				}
			}
		} else {
			if *reedsolo {
				copy(dst, src)
				src = nil
				if len(dst) == MiB/128*136 {
					for i := 0; i < MiB/128*136; i += 136 {
						tmp, err := rsDecode(rs128, dst[i:i+136], !(*fix))
						if err != nil {
							fin.Close()
							fout.Close()
							os.Remove(fout_)
							fmt.Println("The input file is irrecoverably damaged.")
							return 1
						}
						if i == MiB/128*136-136 && done+MiB/128*136 >= int(total) && padded {
							tmp = unpad(tmp)
						}
						src = append(src, tmp...)
					}
				} else {
					chunks := len(dst)/136 - 1
					for i := 0; i < chunks; i++ {
						tmp, err := rsDecode(rs128, dst[i*136:(i+1)*136], !(*fix))
						if err != nil {
							fin.Close()
							fout.Close()
							os.Remove(fout_)
							fmt.Println("The input file is irrecoverably damaged.")
							return 1
						}
						src = append(src, tmp...)
					}
					tmp, err := rsDecode(rs128, dst[int(chunks)*136:], !(*fix))
					if err != nil {
						fin.Close()
						fout.Close()
						os.Remove(fout_)
						fmt.Println("The input file is irrecoverably damaged.")
						return 1
					}
					src = append(src, unpad(tmp)...)
				}
				dst = make([]byte, len(src))
			}
			mac.Write(src)
			chacha.XORKeyStream(dst, src)
			if *paranoid {
				copy(src, dst)
				serpent.XORKeyStream(dst, src)
			}
		}

		_, err = fout.Write(dst)
		if err != nil {
			fin.Close()
			fout.Close()
			os.Remove(fout_)
			fmt.Println("Insufficient disk space.")
			return 1
		}
		if mode == "decrypt" && *reedsolo {
			done += MiB / 128 * 136
		} else {
			done += MiB
		}
		bar.Set(done)

		if counter >= 60*GiB {
			nonce = make([]byte, 24)
			hkdf.Read(nonce)
			chacha, _ = chacha20.NewUnauthenticatedCipher(key, nonce)
			serpentIV = make([]byte, 16)
			hkdf.Read(serpentIV)
			serpent = cipher.NewCTR(s, serpentIV)
			counter = 0
		}
	}
	bar.Set64(total)

	if mode == "encrypt" {
		fout.Seek(int64(309+0*3), 0)
		fout.Write(rsEncode(rs64, keyHash))
		fout.Write(rsEncode(rs32, make([]byte, 32)))
		fout.Write(rsEncode(rs64, mac.Sum(nil)))
	} else {
		if subtle.ConstantTimeCompare(mac.Sum(nil), authTag) != 1 {
			fin.Close()
			fout.Close()
			os.Remove(fout_)
			fmt.Println("The input volume is damaged or modified.")
			if *reedsolo {
				if !(*fix) {
					fmt.Println("Fortunately, this volume is encoded with Reed-Solomon.")
					fmt.Println("Try again using the -fix flag to repair the corruption.")
				} else {
					fmt.Println("The corruption could not be fixed with Reed-Solomon.")
				}
			}
			return 1
		}
	}

	fin.Close()
	fout.Close()
	fmt.Println("\nCompleted.")
	return 0
}

func main() {
	os.Exit(work())
}
