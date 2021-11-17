package hkcrypto

import (
	"crypto/aes"
	"crypto/cipher"
)

type Crypto interface {
	Encrypt(plainText []byte) []byte
	Decrypt(eptData []byte) []byte
}

type StrengthType = uint16

const (
	Strength128 StrengthType = (iota)*64 + 128
	Strength192
	Strength256
)

type aesCbcEncode struct {
	aesKey   [32]byte
	cbcIv    [16]byte
	strength StrengthType
	block    cipher.Block
}

func NewAesCbcEncode(aesKey [32]byte, cbcIv [16]byte, s StrengthType) Crypto {
	key := getKey(aesKey, s)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	return &aesCbcEncode{
		aesKey:   aesKey,
		cbcIv:    cbcIv,
		strength: s,
		block:    block,
	}
}

func getKey(aesKey [32]byte, s StrengthType) []byte {
	switch s {
	case Strength128:
		return aesKey[:16]
	case Strength192:
		return aesKey[:24]
	case Strength256:
		return aesKey[:32]
	}
	return aesKey[:16]
}

func (a *aesCbcEncode) Encrypt(plainText []byte) []byte {
	if v := len(plainText) % a.block.BlockSize(); v != 0 {
		v = a.block.BlockSize() - v
		plainText = append(plainText, make([]byte, v-1)...)
		plainText = append(plainText, byte(v))
	}
	res := make([]byte, len(plainText))
	e := cipher.NewCBCEncrypter(a.block, a.cbcIv[:])
	e.CryptBlocks(res, plainText)
	return res
}

func (a *aesCbcEncode) Decrypt(eptData []byte) []byte {
	originData := make([]byte, len(eptData))
	de := cipher.NewCBCDecrypter(a.block, a.cbcIv[:])
	de.CryptBlocks(originData, eptData)
	d := int(originData[len(originData)-1])
	originData = originData[:len(originData)-d]
	return originData
}
