package encryptor

import (
	"bytes"
	"fmt"
	"os"
)

var xorHeader = []byte("XORv1:")

type XORTransformer struct {
	key []byte
}

func NewXORTransformer() (*XORTransformer, error) {
	k := []byte(os.Getenv("SECRET_KEY_16"))
	if len(k) != 16 {
		return nil, fmt.Errorf("SECRET_KEY_16 must be 16 bytes for XORTransformer")
	}
	return &XORTransformer{key: k}, nil
}

func (encryptor *XORTransformer) Name() string {
	return "xor"
}

func (encryptor *XORTransformer) Encode(data []byte) ([]byte, error) {
	plain := append(xorHeader, data...)
	out := make([]byte, len(plain))
	for i := 0; i < len(plain); i++ {
		out[i] = plain[i] ^ encryptor.key[i%len(encryptor.key)]
	}
	return out, nil
}

func (encryptor *XORTransformer) Decode(data []byte) ([]byte, error) {
	out := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		out[i] = data[i] ^ encryptor.key[i%len(encryptor.key)]
	}

	if len(out) < len(xorHeader) || !bytes.Equal(out[:len(xorHeader)], xorHeader) {
		return nil, fmt.Errorf("invalid ciphertext or header mismatch")
	}

	plain := out[len(xorHeader):]
	if plain == nil {
		return []byte{}, nil
	}
	return plain, nil
}
