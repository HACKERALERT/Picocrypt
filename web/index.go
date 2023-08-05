package main

/*

Picocrypt v1.33 (WebAssembly Version)
Copyright (c) Evan Su
Released under a GNU GPL v3 License
https://github.com/HACKERALERT/Picocrypt

~ In cryptography we trust ~

*/

import (
	"bytes"
	"crypto/rand"
	"strconv"
	"strings"
	"syscall/js"

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
var password string

func work(din []byte, mode string) []byte {
	var salt []byte
	var hkdfSalt []byte
	var nonce []byte
	var keyHash []byte
	var keyHashRef []byte
	var authTag []byte
	var dout []byte

	if mode == "encrypt" {
		salt = make([]byte, 16)
		hkdfSalt = make([]byte, 32)
		nonce = make([]byte, 24)
		rand.Read(salt)
		rand.Read(hkdfSalt)
		rand.Read(nonce)
		dout = append(dout, rsEncode(rs5, []byte("v1.33"))...)
		dout = append(dout, rsEncode(rs5, []byte("00000"))...)
		dout = append(dout, rsEncode(rs5, make([]byte, 5))...)
		dout = append(dout, rsEncode(rs16, salt)...)
		dout = append(dout, rsEncode(rs32, hkdfSalt)...)
		dout = append(dout, rsEncode(rs16, make([]byte, 16))...)
		dout = append(dout, rsEncode(rs24, nonce)...)
		dout = append(dout, make([]byte, 480)...)
	} else {
		errs, tmp := make([]error, 7), make([]byte, 15)
		tmp, din = din[15:30], din[30:]
		tmp, errs[0] = rsDecode(rs5, tmp)
		c, _ := strconv.Atoi(string(tmp))
		tmp, din = din[c*3:c*3+15], din[c*3+15:]
		tmp, errs[1] = rsDecode(rs5, tmp)
		if tmp[0]+tmp[1]+tmp[3] > 0 {
			return []byte{1}
		}
		salt, din = din[:48], din[48:]
		salt, errs[2] = rsDecode(rs16, salt)
		hkdfSalt, din = din[:96], din[96:]
		hkdfSalt, errs[3] = rsDecode(rs32, hkdfSalt)
		nonce, din = din[48:120], din[120:]
		nonce, errs[4] = rsDecode(rs24, nonce)
		keyHashRef, din = din[:192], din[192:]
		keyHashRef, errs[5] = rsDecode(rs64, keyHashRef)
		authTag, din = din[96:288], din[288:]
		authTag, errs[6] = rsDecode(rs64, authTag)
		for _, err := range errs {
			if err != nil {
				return []byte{2}
			}
		}
	}

	key := argon2.IDKey([]byte(password), salt, 4, 1<<20, 4, 32)
	tmp := sha3.New512()
	tmp.Write(key)
	keyHash = tmp.Sum(nil)
	if mode == "decrypt" && !bytes.Equal(keyHash, keyHashRef) {
		return []byte{3}
	}

	counter := 0
	chacha, _ := chacha20.NewUnauthenticatedCipher(key, nonce)
	subkey := make([]byte, 32)
	hkdf := hkdf.New(sha3.New256, key, hkdfSalt, nil)
	hkdf.Read(subkey)
	mac, _ := blake2b.New512(subkey)
	hkdf.Read(make([]byte, 32))

	for {
		var src []byte
		if len(din) == 0 {
			break
		} else if len(din) < MiB {
			src, din = din, nil
		} else {
			src, din = din[:MiB], din[MiB:]
		}
		dst := make([]byte, len(src))

		if mode == "encrypt" {
			chacha.XORKeyStream(dst, src)
			mac.Write(dst)
		} else {
			mac.Write(src)
			chacha.XORKeyStream(dst, src)
		}
		dout = append(dout, dst...)

		counter += MiB
		if counter >= 60*GiB {
			nonce = make([]byte, 24)
			hkdf.Read(nonce)
			chacha, _ = chacha20.NewUnauthenticatedCipher(key, nonce)
			hkdf.Read(make([]byte, 16))
			counter = 0
		}
	}

	if mode == "encrypt" {
		buff := rsEncode(rs64, keyHash)
		buff = append(buff, rsEncode(rs32, make([]byte, 32))...)
		buff = append(buff, rsEncode(rs64, mac.Sum(nil))...)
		for i, v := range buff {
			dout[309+i] = v
		}
	} else {
		if !bytes.Equal(mac.Sum(nil), authTag) {
			return []byte{4}
		}
	}

	return append([]byte{0}, dout...)
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
	quit := make(chan struct{}, 0)
	document := js.Global().Get("document")
	input := document.Call("getElementById", "fin")
	button := document.Call("getElementById", "work")
	callback := js.FuncOf(func(v js.Value, x []js.Value) any {
		var dout []byte
		data := js.Global().Get("Uint8Array").New(x[0])
		din := make([]byte, data.Get("length").Int())
		js.CopyBytesToGo(din, data)
		password = document.Call("getElementById", "password").Get("value").String()
		filename := input.Get("files").Call("item", 0).Get("name").String()
		if strings.HasSuffix(filename, ".pcv") {
			dout = work(din, "decrypt")
		} else {
			dout = work(din, "encrypt")
		}
		arr := js.Global().Get("Uint8Array").New(len(dout))
		js.CopyBytesToJS(arr, dout)
		js.Global().Call("download", arr)
		return nil
	})
	button.Set("onclick", js.FuncOf(func(v js.Value, x []js.Value) any {
		input.Get("files").Call("item", 0).Call("arrayBuffer").Call("then", callback)
		return nil
	}))
	<-quit
}
