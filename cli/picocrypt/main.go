package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/HACKERALERT/crypto/argon2"
	"github.com/HACKERALERT/crypto/blake2b"
	"github.com/HACKERALERT/crypto/chacha20"
	"github.com/HACKERALERT/crypto/hkdf"
	"github.com/HACKERALERT/crypto/sha3"
	"github.com/HACKERALERT/infectious"
)

var MiB = 1 << 20
var GiB = 1 << 30
var rs5, _ = infectious.NewFEC(5, 15)
var rs16, _ = infectious.NewFEC(16, 48)
var rs24, _ = infectious.NewFEC(24, 72)
var rs32, _ = infectious.NewFEC(32, 96)
var rs64, _ = infectious.NewFEC(64, 192)

func work(filename string, password string) int {
	var salt []byte
	var hkdfSalt []byte
	var nonce []byte
	var keyHash []byte
	var keyHashRef []byte
	var authTag []byte

	fin, err := os.Open(filename)
	if err != nil {
		fmt.Println("Couldn't open input file.")
		return 1
	}
	defer fin.Close()

	var fout *os.File
	if strings.HasSuffix(filename, ".pcv") {
		fout, err = os.Create(strings.TrimSuffix(filename, ".pcv"))
	} else {
		fout, err = os.Create(filename + ".pcv")
	}
	if err != nil {
		fmt.Println("Couldn't create output file.")
		return 1
	}
	defer fout.Close()

	if !strings.HasSuffix(filename, ".pcv") {
		salt = make([]byte, 16)
		hkdfSalt = make([]byte, 32)
		nonce = make([]byte, 24)
		rand.Read(salt)
		rand.Read(hkdfSalt)
		rand.Read(nonce)
		fout.Write(rsEncode(rs5, []byte("v1.32")))
		fout.Write(rsEncode(rs5, []byte("00000")))
		fout.Write(rsEncode(rs5, make([]byte, 5)))
		fout.Write(rsEncode(rs16, salt))
		fout.Write(rsEncode(rs32, hkdfSalt))
		fout.Write(rsEncode(rs16, make([]byte, 16)))
		fout.Write(rsEncode(rs24, nonce))
		fout.Write(make([]byte, 480))
	} else {
		errs := make([]error, 7)
		comments := make([]byte, 30)
		fin.Read(comments)
		comments, errs[0] = rsDecode(rs5, comments[15:])
		length, _ := strconv.Atoi(string(comments))
		fin.Read(make([]byte, length*3))
		flags := make([]byte, 15)
		fin.Read(flags)
		flags, errs[1] = rsDecode(rs5, flags)
		salt = make([]byte, 48)
		fin.Read(salt)
		salt, errs[2] = rsDecode(rs16, salt)
		hkdfSalt = make([]byte, 96)
		fin.Read(hkdfSalt)
		hkdfSalt, errs[3] = rsDecode(rs32, hkdfSalt)
		fin.Read(make([]byte, 48))
		nonce = make([]byte, 72)
		fin.Read(nonce)
		nonce, errs[4] = rsDecode(rs24, nonce)
		keyHashRef = make([]byte, 192)
		fin.Read(keyHashRef)
		keyHashRef, errs[5] = rsDecode(rs64, keyHashRef)
		fin.Read(make([]byte, 96))
		authTag = make([]byte, 192)
		fin.Read(authTag)
		authTag, errs[6] = rsDecode(rs64, authTag)
		for _, err := range errs {
			if err != nil {
				fmt.Println("The header is corrupted.")
				return 1
			}
		}
		if flags[0]+flags[1]+flags[3] > 0 {
			fmt.Println("Unsupported volume.")
			return 1
		}
	}

	key := argon2.IDKey([]byte(password), salt, 4, 1<<20, 4, 32)
	tmp := sha3.New512()
	tmp.Write(key)
	keyHash = tmp.Sum(nil)
	if strings.HasSuffix(filename, ".pcv") && !bytes.Equal(keyHash, keyHashRef) {
		fmt.Println("Incorrect password.")
		return 1
	}

	counter := 0
	chacha, _ := chacha20.NewUnauthenticatedCipher(key, nonce)
	subkey := make([]byte, 32)
	hkdf := hkdf.New(sha3.New256, key, hkdfSalt, nil)
	hkdf.Read(subkey)
	mac, _ := blake2b.New512(subkey)
	hkdf.Read(make([]byte, 32))

	for {
		src := make([]byte, MiB)
		size, err := fin.Read(src)
		if err != nil {
			break
		}
		src = src[:size]
		dst := make([]byte, len(src))

		if !strings.HasSuffix(filename, ".pcv") {
			chacha.XORKeyStream(dst, src)
			mac.Write(dst)
		} else {
			mac.Write(src)
			chacha.XORKeyStream(dst, src)
		}
		fout.Write(dst)

		counter += MiB
		if counter >= 60*GiB {
			nonce = make([]byte, 24)
			hkdf.Read(nonce)
			chacha, _ = chacha20.NewUnauthenticatedCipher(key, nonce)
			hkdf.Read(make([]byte, 16))
			counter = 0
		}
	}

	if !strings.HasSuffix(filename, ".pcv") {
		fout.Seek(309, 0)
		fout.Write(rsEncode(rs64, keyHash))
		fout.Write(rsEncode(rs32, make([]byte, 32)))
		fout.Write(rsEncode(rs64, mac.Sum(nil)))
	} else {
		if !bytes.Equal(mac.Sum(nil), authTag) {
			fmt.Println("The file has been modified.")
			return 1
		}
	}

	fmt.Println("Operation successful.")
	return 0
}

func rsEncode(rs *infectious.FEC, data []byte) []byte {
	res := make([]byte, rs.Total())
	rs.Encode(data, func(s infectious.Share) {
		res[s.Number] = s.Data[0]
	})
	return res
}

func rsDecode(rs *infectious.FEC, data []byte) ([]byte, error) {
	tmp := make([]infectious.Share, rs.Total())
	for i := 0; i < rs.Total(); i++ {
		tmp[i].Number = i
		tmp[i].Data = []byte{data[i]}
	}
	res, err := rs.Decode(nil, tmp)
	if err != nil {
		return data[:rs.Total()/3], err
	}
	return res, nil
}

func main() {
	flag.Usage = func() { fmt.Println("Usage: picocrypt -p password <file>") }
	password := flag.String("p", "", "")
	flag.Parse()
	filename := flag.Arg(0)

	if filename == "" || *password == "" || flag.Arg(1) != "" {
		flag.Usage()
		os.Exit(1)
	}
	if _, err := os.Stat(filename); err != nil {
		fmt.Println("Input file not found.")
		os.Exit(1)
	}
	if stat, _ := os.Stat(filename); stat.IsDir() {
		fmt.Println("Directories are not supported.")
		os.Exit(1)
	}
	if !strings.HasSuffix(filename, ".pcv") {
		if _, err := os.Stat(filename + ".pcv"); err == nil {
			fmt.Println("Output already exists.")
			os.Exit(1)
		}
	} else {
		if _, err := os.Stat(strings.TrimSuffix(filename, ".pcv")); err == nil {
			fmt.Println("Output already exists.")
			os.Exit(1)
		}
	}

	if work(filename, *password) != 0 {
		if !strings.HasSuffix(filename, ".pcv") {
			os.Remove(filename + ".pcv")
		} else {
			os.Remove(strings.TrimSuffix(filename, ".pcv"))
		}
		os.Exit(1)
	}
}
