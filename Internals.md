# Internals
If you're wondering about how Picocrypt handles cryptography, you've come to the right place! This page contains the technical details about the cryptographic algorithms and parameters used, as well as how cryptographic values are stored in the header format.

# Core Cryptography
Picocrypt uses the following cryptographic primitives:
- XChaCha20 (cascaded with Serpent in counter mode for paranoid mode)
- Keyed-BLAKE2b for normal mode, HMAC-SHA3 for paranoid mode (256-bit key, 512-bit digest)
- HKDF-SHA3 for deriving a subkey for the MAC above, as well as a key for Serpent
- Argon2id:
    - Normal mode: 4 passes, 1 GiB memory, 4 threads
    - Paranoid mode: 8 passes, 1 GiB memory, 8 threads

All primitives used are from the well-known [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto) module.

# Counter Overflow
Since XChaCha20 has a max message size of 256 GiB, Picocrypt will use the HKDF-SHA3 mentioned above to generate a new nonce for XChaCha20 and a new IV for Serpent if the total encrypted data is more than 60 GiB. While this threshold can be increased up to 256 GiB, Picocrypt uses 60 GiB to prevent any edge cases with blocks or the counter used by Serpent.

# Header Format
A Picocrypt volume's header is encoded with Reed-Solomon by default since it is, after all, the most important part of the entire file. An encoded value will take up three times the size of the unencoded value.

**All offsets and sizes below are in bytes.**
| Offset | Encoded size | Decoded size | Description
| ------ | ------------ | ------------ | -----------
| 0      | 15           | 5            | Version number (ex. "v1.15")
| 15     | 15           | 5            | Length of comments, zero-padded to 5 bytes
| 30     | 3C           | C            | Comments with a length of C characters
| 30+3C  | 15           | 5            | Flags (paranoid mode, use keyfiles, etc.)
| 45+3C  | 48           | 16           | Salt for Argon2
| 93+3C  | 96           | 32           | Salt for HKDF-SHA3
| 189+3C | 48           | 16           | IV for Serpent
| 237+3C | 72           | 24           | Nonce for XChaCha20
| 309+3C | 192          | 64           | SHA3-512 of encryption key
| 501+3C | 96           | 32           | SHA3-256 of keyfile key
| 597+3C | 192          | 64           | Authentication tag (BLAKE2b/HMAC-SHA3)
| 789+3C |              |              | Encrypted contents of input data

# Keyfile Design
Picocrypt allows the use of keyfiles as an additional form of authentication. Picocrypt's unique "Require correct order" feature enforces the user to drop keyfiles into the window in the same order as they did when encrypting in order to decrypt the volume successfully. Here's how it works:

If correct order is not required, Picocrypt will take the SHA3-256 of each keyfile individually and XOR the hashes together. Finally, the result is XORed with the master key. Because the XOR operation is both commutative and associative, the order in which the keyfile hashes are XORed with each other doesn't matter - the end result is the same.

If correct order is required, Picocrypt will concatenate the keyfiles together in the order they were dropped into the window and take the SHA3-256 of the combined keyfiles. If the order is not correct, the keyfiles, when appended to each other, will result in a different file, and thus a different hash. So, the correct order of keyfiles is required to decrypt the volume successfully.

# Reed-Solomon
By default, all Picocrypt volume headers are encoded with Reed-Solomon to improve resiliency against bit rot. The header uses N+2N encoding, where N is the size of a particular header field such as the version number, and 2N is the number of parity bytes added. Using the Berlekamp-Welch algorithm, Picocrypt is able to automatically detect and correct up to 2N/2=N broken bytes.

If Reed-Solomon is to be used with the input data itself, the data will be encoded using 128+8 encoding, with the data being read in 1 MiB chunks and encoded in 128-byte blocks, and the final block padded to 128 bytes using PKCS#7.

To address the edge case where the final 128-byte block happens to be padded so that it completes a full 1 MiB chunk, a flag is used to distinguish whether the last 128-byte block was padded originally or if it is just a full 128-byte block of data.

# Deniability
Plausible deniability in Picocrypt is achieved by simply re-encrypting the volume but without storing any identifiable header data. A new Argon2 salt and XChaCha20 nonce will be generated and stored in the deniable volume, but since both values are random, they don't reveal anything. A deniable volume will look something like this:
```
[argon2 salt][xchacha20 nonce][encrypted stream of bytes]
```

# Just Read the Code
Picocrypt is a very simple tool and only has one source file. The source Go file is just 2K lines and a lot of the code is dealing with the UI. The core cryptography code is only about 1K lines of code, and even so, a lot of that code deals with the UI and other features of Picocrypt. So if you need more information about how Picocrypt works, just read the code. It's not long, and it is well commented and will explain what happens under the hood better than a document can.
