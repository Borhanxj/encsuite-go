# Encryption Suite in Go

A structured, step-by-step (micro-step) implementation of classic and modern cryptographic techniques in Go.

---

## Features

### Modern Crypto (Core)
- [ ] **Symmetric (AEAD)**
  - [ ] AES-GCM
  - [ ] ChaCha20-Poly1305
  - [ ] XChaCha20-Poly1305
- [ ] **Asymmetric**
  - [ ] RSA-OAEP (encryption)
  - [ ] RSA-PSS (signatures)
  - [ ] X25519 (key exchange)
  - [ ] Ed25519 (signatures)
  - [ ] P-256 ECDH/ECDSA (interop)
- [ ] **KDFs & Passwords**
  - [ ] HKDF
  - [ ] Argon2id
  - [ ] scrypt
  - [ ] PBKDF2-HMAC (interop)
- [ ] **Hashes & MAC**
  - [ ] SHA-2: SHA-224/256/384/512, SHA-512/224, SHA-512/256
  - [ ] SHA-3: SHA3-256/512, SHAKE128/256
  - [ ] Keccak-256 (Ethereum)
  - [ ] BLAKE2b / BLAKE2s
  - [ ] BLAKE3
  - [ ] RIPEMD-160 (Bitcoin)
  - [ ] Whirlpool (optional)
  - [ ] HMAC (with any hash)
  - [ ] Bitcoin-style `SHA256d` (double SHA-256)

### Encoders / Decoders (Crypto Ecosystem)
- [ ] Hex (with/without `0x`)
- [ ] Base64
- [ ] Base64URL
- [ ] Base32 (RFC 4648)
- [ ] Base58 + Base58Check (Bitcoin)
- [ ] Bech32 / Bech32m (BTC SegWit)
- [ ] Varint (Bitcoin-style)
- [ ] Length-prefixed encodings
- [ ] Checksums: Base58Check, Bech32 polymod, CRC32C
- [ ] Ethereum: EIP-55 checksum casing
- [ ] PEM/DER key IO (PKCS#8, SPKI)
- [ ] Streaming encoders (Reader/Writer wrappers)

### Classical / Historical Ciphers (Educational)
- [ ] Substitution: Caesar/ROT13, Affine, Monoalphabetic
- [ ] Polyalphabetic: Vigenère, Beaufort, Autokey
- [ ] Polygraph: Playfair, Hill (n×n)
- [ ] Transposition: Rail Fence, Columnar
- [ ] Extras: ADFGX / ADFGVX, Enigma simulator
- [ ] Analysis helpers: frequency tables, Index of Coincidence, chi-square, Kasiski test  
  > 🚫 **Not secure.** For learning and demos only.

### Legacy (Disabled by Default — For Compatibility)
- [ ] AES-CBC / CTR / CFB (legacy, foot-gun)
- [ ] DES / 3DES
- [ ] RC4
- [ ] MD5, SHA-1 (legacy validation only; blocked in APIs)

---

## Future Work

### Advanced Elliptic Curve Cryptography
- [ ] Ed448 signatures
- [ ] Curve448 (X448 key exchange)
- [ ] Pairing-friendly curves (BLS12-381 for aggregate signatures)

### Post-Quantum Cryptography
- [ ] Lattice-based KEMs: Kyber, Saber
- [ ] Lattice-based signatures: Dilithium, Falcon
- [ ] Code-based: Classic McEliece
- [ ] Hash-based: SPHINCS+

### Fully Homomorphic Encryption (FHE) & Advanced Schemes
- [ ] Toy FHE schemes (BFV / BGV / CKKS) in Go (educational)
- [ ] Homomorphic operations demo (encrypted add/mul)
- [ ] Secure multiparty computation (SMPC) toy protocols
- [ ] Zero-Knowledge Proof (ZKP) frameworks (zk-SNARKs, Bulletproofs)

### Applied Cryptography Extensions
- [ ] Noise Protocol Framework patterns (IK/XX)
- [ ] Signal-style double ratchet demo
- [ ] TLS-like handshake mini-protocol
- [ ] JWT / JOSE helpers
- [ ] Blockchain utilities (transaction signing, Merkle proofs)

---

## Vision

The long-term vision is to make **`encsuite-go`** a practical learning toolkit for cryptography in Go.