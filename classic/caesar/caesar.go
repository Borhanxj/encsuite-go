package caesar

// Encrypt shifts the input text by 'key' positions (Caesar cipher).
func Encrypt(plaintext string, key int) string {
	ciphertext := ""

	for i := 0; i < len(plaintext); i++ {
		char := plaintext[i]

		if char >= 'A' && char <= 'Z' {
			shifted := (int(char-'A') + key + 26) % 26
			ciphertext += string('A' + byte(shifted))
		} else if char >= 'a' && char <= 'z' {
			shifted := (int(char-'a') + key + 26) % 26
			ciphertext += string('a' + byte(shifted))
		} else {
			ciphertext += string(char)
		}
	}

	return ciphertext
}

func Decrypt(ciphertext string, key int) string {
	return Encrypt(ciphertext, -key)
}
