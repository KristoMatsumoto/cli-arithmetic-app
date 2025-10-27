package encryptor

import (
	"bytes"
	"cli-arithmetic-app/app/config"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/blowfish"
)

type BlowfishTransformer struct {
	key []byte
}

func NewBlowfishTransformer() (*BlowfishTransformer, error) {
	config.GetConfig()
	k := []byte(config.C.SecretKeys.S_16)
	if len(k) < 1 || len(k) > 56 {
		return nil, fmt.Errorf("SECRET_KEY must be 1-56 bytes for Blowfish")
	}
	return &BlowfishTransformer{key: k}, nil
}

func (encryptor *BlowfishTransformer) Name() string { return "blowfish" }

func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func unpad(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return nil, fmt.Errorf("empty data")
	}
	padding := int(src[len(src)-1])
	if padding > len(src) {
		return nil, fmt.Errorf("invalid padding")
	}
	return src[:len(src)-padding], nil
}

func (t *BlowfishTransformer) Encode(data []byte) ([]byte, error) {
	block, _ := blowfish.NewCipher(t.key)
	padded := pad(data, blowfish.BlockSize)
	ciphertext := make([]byte, len(padded))
	iv := make([]byte, blowfish.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext, padded)
	return append(iv, ciphertext...), nil
}

func (t *BlowfishTransformer) Decode(data []byte) ([]byte, error) {
	if len(data) < blowfish.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := data[:blowfish.BlockSize]
	ct := data[blowfish.BlockSize:]
	block, _ := blowfish.NewCipher(t.key)
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ct, ct)
	return unpad(ct)
}
