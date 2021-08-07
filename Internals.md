# Internals
If you're wondering about how Picocrypt handles cryptography, you've come to the right place! This page contains the technical details about the cryptographic algorithms and parameters used. This page does not go in-depth, as you can just review the source (there's only one file) to understand the nitty-gritty details.

# Core Cryptography
Picocrypt uses the following cryptographic primitives:
- XChaCha20 (cascaded with Serpent for paranoid mode)
- HMAC-SHA3 for normal mode, keyed-BLAKE2b for fast mode (512 bits)
- HKDF-SHA3 for deriving a subkey used with the MAC above
- Argon2id (8 passes, 1 GiB memory, 8 threads) for normal mode, (4 passes, 128 MiB memory, 4 threads) for fast mode

All algorithms used are from the well-known golang.org/x/crypto module.
