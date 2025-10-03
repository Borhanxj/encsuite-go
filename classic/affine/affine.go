package affine

import "github.com/borhanxj/encsuite-go/internal/modmath"

// Encrypt applies the Affine cipher: (a*x + b) mod 26
func Encrypt(plaintext string, key_a, key_b int) string {
	ciphertext := ""

	for i := 0; i < len(plaintext); i++ {
		ch := plaintext[i]

		if ch >= 'A' && ch <= 'Z' {
			x := int(ch - 'A')
			enc := (key_a*x + key_b + 26) % 26
			ciphertext += string('A' + byte(enc))
		} else if ch >= 'a' && ch <= 'z' {
			x := int(ch - 'a')
			enc := (key_a*x + key_b + 26) % 26
			ciphertext += string('a' + byte(enc))
		} else {
			ciphertext += string(ch)
		}
	}
	return ciphertext
}

// Decrypt calculates the modular inverse using EEA
func Decrypt(ciphertext string, key_a, key_b int) (string, bool) {
	a_Inverse, ok := modmath.ModInverse(key_a, 26)

	if ok {

		plaintext := ""

		for i := 0; i < len(ciphertext); i++ {
			char := ciphertext[i]
			if char >= 'A' && char <= 'Z' {
				y := int(char - 'A')
				dec := (a_Inverse * (y - key_b + 26)) % 26
				plaintext += string('A' + byte(dec))
			} else if char >= 'a' && char <= 'z' {
				y := int(char - 'a')
				dec := (a_Inverse * (y - key_b + 26)) % 26
				plaintext += string('a' + byte(dec))
			} else {
				plaintext += string(char)
			}
		}

		return plaintext, true

	} else {
		return "Error", false
	}
}
