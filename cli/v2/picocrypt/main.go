package main

import (
	"archive/zip"
	"bytes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"flag"
	"fmt"
	"hash"
	"io"
	"math"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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

var f *bool
var k *bool
var p *bool
var r *bool
var mode string

func parse() int {
	flag.Usage = func() {
		fmt.Println("Usage: picocrypt <item1> [<item2> ...]")
		fmt.Println("Items: can be files, folders, or globs")
		fmt.Println("Flags:")
		flag.PrintDefaults()
		os.Exit(1)
	}
	f = flag.Bool("f", false, "(decryption) attempt to fix corruption")
	k = flag.Bool("k", false, "(decryption) keep output even if corrupted")
	p = flag.Bool("p", false, "(encryption) use paranoid mode")
	r = flag.Bool("r", false, "(encryption) encode with Reed-Solomon")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
	}
	for _, v := range flag.Args() {
		if strings.HasPrefix(v, "-") {
			fmt.Println("Flags must be provided before arguments!")
			return 1
		}
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
				fmt.Println("Multiple items must not contain volumes.")
				return 1
			}
		}
	}

	return 0
}

var password []byte
var confirmp []byte
var err error

func auth() int {
	if mode == "encrypt" {
		fmt.Print("Password: ")
		password, err = term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading password!")
			return 1
		}
		fmt.Print(strings.Repeat("*", len(password)), " | Confirm: ")
		confirmp, err = term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading password!")
			return 1
		}
		fmt.Println(strings.Repeat("*", len(confirmp)))
		if !bytes.Equal(password, confirmp) {
			fmt.Println("Passwords don't match!")
			return 1
		}
	} else {
		fmt.Print("Password: ")
		password, err = term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading password!")
			return 1
		}
		fmt.Println(strings.Repeat("*", len(password)))
	}

	return 0
}

var pin string
var pout string
var pzip string
var file *os.File
var writer *zip.Writer
var files []string
var interrupted bool

func prepare() int {
	if mode == "decrypt" {
		pin = flag.Arg(0)
		pout = strings.TrimSuffix(pin, ".pcv")
	} else {
		stat, err := os.Stat(flag.Arg(0))
		if flag.NArg() == 1 && err == nil && !stat.IsDir() {
			pin = flag.Arg(0)
			pout = pin + ".pcv"
		} else {
			items := []string{}
			for _, v := range flag.Args() {
				if strings.Contains(v, "../") || strings.HasPrefix(v, "/") {
					fmt.Println("Cannot encrypt outside of current directory.")
					return 1
				}
				matches, err := filepath.Glob(v)
				if err != nil {
					fmt.Println("Invalid glob pattern:", v)
					return 1
				}
				items = append(items, matches...)
			}
			for _, v := range items {
				stat, err := os.Stat(v)
				if err != nil {
					fmt.Println("Cannot access input:", v)
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
				fmt.Println("Nothing to encrypt!")
				return 1
			}
		}
	}

	return 0
}

func compress() int {
	if files == nil {
		return 0
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Cannot get current working directory!")
		return 1
	}
	dir = filepath.ToSlash(dir)
	file, err = os.CreateTemp("", "picocrypt-cli-v2-*.tmp")
	if err != nil {
		fmt.Println("Cannot create temporary file!")
		return 1
	}
	pzip = file.Name()
	writer = zip.NewWriter(file)

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
		header.Name = strings.TrimPrefix(abs, dir)
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
			fmt.Println("Read access to input denied:", path)
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
			if interrupted {
				time.Sleep(1 * time.Second)
			} else {
				writer.Close()
				file.Close()
				fmt.Println("Insufficient disk space!")
				return 1
			}
		}
	}
	writer.Close()
	file.Close()
	pin = file.Name()
	pout = "encrypted-" + strconv.Itoa(int(time.Now().Unix())) + ".zip.pcv"

	return 0
}

var fin *os.File
var fout *os.File
var padded bool
var salt []byte
var hkdfSalt []byte
var serpentIV []byte
var nonce []byte
var keyHash []byte
var keyHashRef []byte
var authTag []byte
var key []byte
var mac hash.Hash
var MiB = 1 << 20
var GiB = 1 << 30
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
	fin, err = os.Open(pin)
	if err != nil {
		fmt.Println("Error accessing input file:", pin)
		return 1
	}
	_, err = os.Stat(pout)
	if err == nil {
		fmt.Println("Output file already exists!")
		return 1
	}
	fout, err = os.Create(pout)
	if err != nil {
		fmt.Println("Error creating output file:", pout)
		return 1
	}
	stat, err := os.Stat(pin)
	if err != nil {
		fmt.Println("Error accessing input file:", pin)
		return 1
	}
	total := stat.Size()
	if mode == "decrypt" {
		total -= 789
	}

	if mode == "encrypt" {
		errs := make([]error, 10)
		salt = make([]byte, 16)
		hkdfSalt = make([]byte, 32)
		serpentIV = make([]byte, 16)
		nonce = make([]byte, 24)
		_, errs[0] = fout.Write(rsEncode(rs5, []byte("v1.34")))
		_, errs[1] = fout.Write(rsEncode(rs5, []byte("00000")))
		flags := make([]byte, 5)
		if *p {
			flags[0] = 1
		}
		if *r {
			flags[3] = 1
		}
		if total%int64(MiB) >= int64(MiB)-128 {
			flags[4] = 1
		}
		_, errs[2] = fout.Write(rsEncode(rs5, flags))
		rand.Read(salt)
		rand.Read(hkdfSalt)
		rand.Read(serpentIV)
		rand.Read(nonce)
		_, errs[3] = fout.Write(rsEncode(rs16, salt))
		_, errs[4] = fout.Write(rsEncode(rs32, hkdfSalt))
		_, errs[5] = fout.Write(rsEncode(rs16, serpentIV))
		_, errs[6] = fout.Write(rsEncode(rs24, nonce))
		_, errs[7] = fout.Write(make([]byte, 192))
		_, errs[8] = fout.Write(make([]byte, 96))
		_, errs[9] = fout.Write(make([]byte, 192))
		for _, err := range errs {
			if err != nil {
				fin.Close()
				fout.Close()
				fmt.Println("Insufficient disk space!")
				return 1
			}
		}
	} else {
		errs := make([]error, 9)
		version := make([]byte, 15)
		fin.Read(version)
		_, errs[0] = rsDecode(rs5, version, !(*f))
		tmp := make([]byte, 15)
		fin.Read(tmp)
		tmp, errs[1] = rsDecode(rs5, tmp, !(*f))
		comments, _ := strconv.Atoi(string(tmp))
		fin.Read(make([]byte, comments*3))
		total -= int64(comments) * 3
		flags := make([]byte, 15)
		fin.Read(flags)
		flags, errs[2] = rsDecode(rs5, flags, !(*f))
		*p = flags[0] == 1
		*r = flags[3] == 1
		padded = flags[4] == 1
		if flags[1] == 1 {
			fin.Close()
			fout.Close()
			fmt.Println("Keyfiles are not supported!")
			return 1
		}
		salt = make([]byte, 48)
		fin.Read(salt)
		salt, errs[3] = rsDecode(rs16, salt, !(*f))
		hkdfSalt = make([]byte, 96)
		fin.Read(hkdfSalt)
		hkdfSalt, errs[4] = rsDecode(rs32, hkdfSalt, !(*f))
		serpentIV = make([]byte, 48)
		fin.Read(serpentIV)
		serpentIV, errs[5] = rsDecode(rs16, serpentIV, !(*f))
		nonce = make([]byte, 72)
		fin.Read(nonce)
		nonce, errs[6] = rsDecode(rs24, nonce, !(*f))
		keyHashRef = make([]byte, 192)
		fin.Read(keyHashRef)
		keyHashRef, errs[7] = rsDecode(rs64, keyHashRef, !(*f))
		fin.Read(make([]byte, 96))
		authTag = make([]byte, 192)
		fin.Read(authTag)
		authTag, errs[8] = rsDecode(rs64, authTag, !(*f))
		for _, err := range errs {
			if err != nil {
				fin.Close()
				fout.Close()
				fmt.Println("The volume header is irrecoverably damaged!")
				return 1
			}
		}
	}

	if *p {
		key = argon2.IDKey(password, salt, 8, 1<<20, 8, 32)
	} else {
		key = argon2.IDKey(password, salt, 4, 1<<20, 4, 32)
	}
	tmp := sha3.New512()
	tmp.Write(key)
	keyHash = tmp.Sum(nil)
	if mode == "decrypt" {
		if !bytes.Equal(keyHash, keyHashRef) {
			fin.Close()
			fout.Close()
			fmt.Println("Incorrect password!")
			return 1
		}
	}

	done, counter := 0, 0
	chacha, _ := chacha20.NewUnauthenticatedCipher(key, nonce)
	subkey := make([]byte, 32)
	hkdf := hkdf.New(sha3.New256, key, hkdfSalt, nil)
	hkdf.Read(subkey)
	if *p {
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
		if mode == "decrypt" && *r {
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
		bar.Write(src)

		if mode == "encrypt" {
			if *p {
				serpent.XORKeyStream(dst, src)
				copy(src, dst)
			}
			chacha.XORKeyStream(dst, src)
			mac.Write(dst)
			if *r {
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
			if *r {
				copy(dst, src)
				src = nil
				if len(dst) == MiB/128*136 {
					for i := 0; i < MiB/128*136; i += 136 {
						tmp, err := rsDecode(rs128, dst[i:i+136], !(*f))
						if err != nil {
							fin.Close()
							fout.Close()
							fmt.Println("\nThe input file is irrecoverably damaged.")
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
						tmp, err := rsDecode(rs128, dst[i*136:(i+1)*136], !(*f))
						if err != nil {
							fin.Close()
							fout.Close()
							fmt.Println("\nThe input file is irrecoverably damaged.")
							return 1
						}
						src = append(src, tmp...)
					}
					tmp, err := rsDecode(rs128, dst[int(chunks)*136:], !(*f))
					if err != nil {
						fin.Close()
						fout.Close()
						fmt.Println("\nThe input file is irrecoverably damaged.")
						return 1
					}
					src = append(src, unpad(tmp)...)
				}
				dst = make([]byte, len(src))
			}
			mac.Write(src)
			chacha.XORKeyStream(dst, src)
			if *p {
				copy(src, dst)
				serpent.XORKeyStream(dst, src)
			}
		}

		_, err = fout.Write(dst)
		if err != nil {
			if interrupted {
				time.Sleep(1 * time.Second)
			} else {
				fin.Close()
				fout.Close()
				fmt.Println("\nInsufficient disk space!")
				return 1
			}
		}
		if mode == "decrypt" && *r {
			done += MiB / 128 * 136
		} else {
			done += MiB
		}

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

	if mode == "encrypt" {
		fout.Seek(309, 0)
		fout.Write(rsEncode(rs64, keyHash))
		fout.Write(rsEncode(rs32, make([]byte, 32)))
		fout.Write(rsEncode(rs64, mac.Sum(nil)))
	} else {
		if !bytes.Equal(mac.Sum(nil), authTag) {
			fin.Close()
			fout.Close()
			if *k {
				fmt.Println("\nThe modified output has been kept.")
				return 0
			} else {
				fmt.Println("\nThe input volume is damaged or modified!")
				if *r {
					fmt.Println("Fortunately, this volume is encoded with Reed-Solomon.")
					fmt.Println("Try again using the '-f' flag to repair the corruption.")
				}
				return 1
			}
		}
	}

	fin.Close()
	fout.Close()
	fmt.Println("Completed ->", fout.Name())
	return 0
}

func main() {
	if parse() == 1 {
		os.Exit(1)
	}
	if auth() == 1 {
		os.Exit(1)
	}
	if prepare() == 1 {
		os.Exit(1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		interrupted = true
		if fin != nil {
			fin.Close()
		}
		if fout != nil {
			fmt.Print("\nSystem interrupt detected, cleaning up incomplete output: ")
			fout.Close()
			if err := os.Remove(fout.Name()); err == nil {
				fmt.Print("Success.")
			} else {
				fmt.Print("Failure.")
			}
		}
		if pzip != "" {
			fmt.Print("\nSystem interrupt detected, cleaning up temporary files: ")
			writer.Close()
			file.Close()
			if err := os.Remove(pzip); err == nil {
				fmt.Print("Success.")
			} else {
				fmt.Print("Failure.")
			}
		}
		fmt.Println()
		os.Exit(1)
	}()

	if compress() == 1 {
		os.Remove(pzip)
		os.Exit(1)
	} else {
		defer os.Remove(pzip)
	}
	if work() == 1 {
		if pzip != "" {
			os.Remove(pzip)
		}
		if fout != nil {
			os.Remove(fout.Name())
		}
		os.Exit(1)
	}
}
