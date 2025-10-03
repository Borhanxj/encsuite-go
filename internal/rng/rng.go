package rng

import (
	"crypto/rand"
	"math/big"
)

// Randomness source
type Source interface {
	RandIntN(n int) (int, error)
}

// Source : Crypto
type CryptoSource struct{}

func (CryptoSource) RandIntN(n int) (int, error) {
	x, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		return 0, err
	}
	return int(x.Int64()), nil
}
