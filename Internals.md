# Internals
If you're wondering about how Picocrypt handles cryptography, you've come to the right place! This page contains the technical details about the cryptographic algorithms and parameters used. This page does not go in-depth, as you can just review the source (there's only one file) to understand the nitty-gritty details.

# Core Cryptography
Picocrypt uses the following cryptographic primitives:
- XChaCha20 (cascaded with Serpent in CTR mode for paranoid mode)
- Keyed-BLAKE2b for normal mode, HMAC-SHA3 for paranoid mode (256-bit key, 512-bit digest)
- HKDF-SHA3 for deriving a subkey used with the MAC above, as well as a key for Serpent
- Argon2id:
    - Normal mode: 4 passes, 1 GiB memory, 4 threads
    - Paranoid mode: 8 passes, 1 GiB memory, 8 threads

All primitives used are from the well-known golang.org/x/crypto module.

# Keyfile Design
Picocrypt allows the use of keyfiles as an additional (or only) form of authentication. Picocrypt's unique "Require correct order" feature enforces the user to drop keyfiles into the window in the exact same order as he/she did when encrypting, in order to decrypt the volume successfully. Here's how it works:

If "Require correct order" is not checked, Picocrypt will take the SHA3 hash of each file individually, and XORs the hashes together. Finally, the result is XORed to the master key. Because the XOR operation is both commutative and associative, the order in which the keyfiles hashes are XORed to each other doesn't matter -- the end result is the same.

If "Require correct order" is checked, Picocrypt will combine (concatenate) the files together in the order they were dropped into the window, and take the SHA3 hash of the combined keyfiles. If the order is not correct, the keyfiles, when appended to each other, will result in a different file, and therefore a different hash. Thus, the correct order of keyfiles is required to successfully decrypt the volume.

# Header Format
Work in progress...
