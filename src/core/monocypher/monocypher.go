package monocypher

/*
#include "monocypher.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func Lock(plaintext,nonce,key []byte) (mac,ciphertext []byte){
	// void crypto_lock(
	// 	uint8_t        mac[16],
	// 	uint8_t       *ciphertext,
	// 	const uint8_t  key[32],
	// 	const uint8_t  nonce[24],
	// 	const uint8_t *plaintext,
	//	size_t         text_size
	// );

	CSize := (C.size_t)(len(plaintext))
	CPlain := (*C.uint8_t)(unsafe.Pointer(C.CBytes([]uint8(plaintext))))
	defer C.free(unsafe.Pointer(CPlain))
	CKey := (*C.uint8_t)(unsafe.Pointer(C.CBytes([]uint8(key[:32]))))
	defer C.free(unsafe.Pointer(CKey))
	CNonce := (*C.uint8_t)(unsafe.Pointer(C.CBytes([]uint8(nonce[:24]))))
	defer C.free(unsafe.Pointer(CNonce))
	CMac := (*C.uint8_t)(unsafe.Pointer(C.CBytes(make([]uint8,16))))
	defer C.free(unsafe.Pointer(CMac))
	CCipher := (*C.uint8_t)(unsafe.Pointer(C.CBytes(make([]uint8,len(plaintext)))))
	defer C.free(unsafe.Pointer(CCipher))

	C.crypto_lock(CMac,CCipher,CKey,CNonce,CPlain,CSize)
	var GCipher []byte = C.GoBytes(unsafe.Pointer(CCipher),C.int(len(plaintext)))
	var GMac []byte = C.GoBytes(unsafe.Pointer(CMac),C.int(16))
	return GMac,GCipher
}

func Unlock(ciphertext,nonce,key,mac []byte) (plaintext []byte,authentic bool){
	// int crypto_unlock(
	// 	uint8_t       *plaintext,
	// 	const uint8_t  key[32],
	// 	const uint8_t  nonce[24],
	// 	const uint8_t  mac[16],
	// 	const uint8_t *ciphertext,
	// 	size_t         text_size
	// );

	CSize := (C.size_t)(len(ciphertext))
	CCipher := (*C.uint8_t)(unsafe.Pointer(C.CBytes(ciphertext)))
	defer C.free(unsafe.Pointer(CCipher))
	CKey := (*C.uint8_t)(unsafe.Pointer(C.CBytes([]uint8(key[:32]))))
	defer C.free(unsafe.Pointer(CKey))
	CNonce := (*C.uint8_t)(unsafe.Pointer(C.CBytes([]uint8(nonce[:24]))))
	defer C.free(unsafe.Pointer(CNonce))
	CMac := (*C.uint8_t)(unsafe.Pointer(C.CBytes(mac)))
	defer C.free(unsafe.Pointer(CMac))
	CPlain := (*C.uint8_t)(unsafe.Pointer(C.CBytes(make([]uint8,len(ciphertext)))))
	defer C.free(unsafe.Pointer(CPlain))

	valid := int(C.crypto_unlock(CPlain,CKey,CNonce,CMac,CCipher,CSize))
	var GPlain []byte = C.GoBytes(unsafe.Pointer(CPlain),C.int(len(ciphertext)))
	if valid==-1{
		return GPlain,false
	}else{
		return GPlain,true
	}
}
