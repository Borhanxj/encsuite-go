package caesar

// ROT13 applies Caesar cipher with key = 13
// Encrypt and Decrypt is the same function
func ROT13(plaintext string) string {
	return Encrypt(plaintext, 13)
}
