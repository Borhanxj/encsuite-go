package main

import (
	"fmt"

	//"github.com/borhanxj/encsuite-go/classic/caesar"
	//"github.com/borhanxj/encsuite-go/classic/affine"
	"github.com/borhanxj/encsuite-go/internal/keygen"
	"github.com/borhanxj/encsuite-go/internal/rng"
)

func main() {
	src := rng.CryptoSource{}
	key,_ := keygen.CaesarKeyGen(src)
	fmt.Println(key.Shift)
}
