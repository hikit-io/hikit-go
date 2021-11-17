package main

import (
	"fmt"
	"go.hikit.io/hkcrypto"
	"testing"
)

func Test_example(t *testing.T) {
	code := hkcrypto.NewAesCbcEncode([32]byte{}, [16]byte{}, hkcrypto.Strength256)
	edata := code.Encrypt([]byte("hikit.io"))
	fmt.Println(string(edata))
	ddata := code.Decrypt([]byte(edata))
	fmt.Println(string(ddata))
}
