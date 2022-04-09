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

All primitives used are from the well-known [golang.org/x/crypto](https://golang.org/x/crypto) module.

# Header Format
A Picocrypt volume's header is encoded with Reed-Solomon by default, since it is, after all, the most important part of the entire file. An encoded value will take up three times the size of the unencoded value.

**All offsets and sizes below are in bytes.**
| Offset | Encoded size | Decoded size | Description
| ------ | ------------ | ------------ | -----------
| 0      | 15           | 5            | Version number (ex. "v1.15")
| 15     | 15           | 5            | Length of comments, zero-padded to 5 bytes
| 30     | 3C           | C            | Comments with a length of C characters
| 30+3C  | 15           | 5            | Flags (paranoid mode, use keyfiles, etc.)
| 45+3C  | 48           | 16           | Salt for Argon2
| 93+3C  | 96           | 32           | Salt for HKDF-SHA3
| 189+3C | 48           | 16           | Salt for Serpent
| 237+3C | 72           | 24           | Nonce for XChaCha20
| 309+3C | 192          | 64           | SHA3-512 of encryption key
| 501+3C | 96           | 32           | Hash of keyfile key
| 597+3C | 192          | 64           | Authentication tag (BLAKE2b/HMAC-SHA3)
| 789+3C |              |              | Encrypted contents of input data

# Keyfile Design
Picocrypt allows the use of keyfiles as an additional form of authentication. Picocrypt's unique "Require correct order" feature enforces the user to drop keyfiles into the window in the exact same order as they did when encrypting, in order to decrypt the volume successfully. Here's how it works:

If "Require correct order" is not checked, Picocrypt will take the SHA3 hash of each file individually, and XOR the hashes together. Finally, the result is XORed with the master key. Because the XOR operation is both commutative and associative, the order in which the keyfile hashes are XORed with each other doesn't matter - the end result is the same.

If "Require correct order" is checked, Picocrypt will concatenate the files together in the order they were dropped into the window and take the SHA3 hash of the combined keyfiles. If the order is not correct, the keyfiles, when appended to each other, will result in a different file, and therefore a different hash. Thus, the correct order of keyfiles is required to successfully decrypt the volume.

# Reed-Solomon
By default, all Picocrypt volume headers are encoded with Reed-Solomon to improve resiliency to bit rot, etc. The header uses N+2N encoding, where N is the size of a particular header field such as the version number or Argon2 salt. If Reed-Solomon is to be used with the input data itself, the data is encoded using 128+8 encoding, with the data being read in chunks of 1 MiB, and the final set padded to 128 bytes using PKCS#7.
