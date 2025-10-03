package keygen

import "github.com/borhanxj/encsuite-go/internal/rng"

// --------------------------------------------
// Caesar Cipher

type CaesarKey struct {
	Shift int
}

func CaesarKeyGen(src rng.Source) (CaesarKey, error) {
	shift, err := src.RandIntN(26)
	if err != nil {
		return CaesarKey{}, err
	}
	return CaesarKey{Shift: shift}, nil
}

// --------------------------------------------

