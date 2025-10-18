package encryptor

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

var aesCFBHeader = []byte("ENCv1:")

type AESTransformer struct {
	key []byte
}

func NewAESTransformer() (*AESTransformer, error) {
	key := os.Getenv("SECRET_KEY_16")
	if key == "" {
		return nil, fmt.Errorf("SECRET_KEY not set in environment")
	}

	k := []byte(key)
	switch len(k) {
	case 16, 24, 32:
		// ok
	default:
		return nil, fmt.Errorf("SECRET_KEY must be 16, 24, or 32 bytes long")
	}

	return &AESTransformer{key: k}, nil
}

func (encryptor *AESTransformer) Name() string { return "aes" }

func (encryptor *AESTransformer) Encode(input []byte) ([]byte, error) {
	block, err := aes.NewCipher(encryptor.key)
	if err != nil {
		return nil, err
	}

	plain := append(aesCFBHeader, input...)

	ciphertext := make([]byte, aes.BlockSize+len(plain))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plain)

	return ciphertext, nil
}

func (encryptor *AESTransformer) Decode(input []byte) ([]byte, error) {
	if len(input) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	block, err := aes.NewCipher(encryptor.key)
	if err != nil {
		return nil, err
	}

	iv := input[:aes.BlockSize]
	ciphertext := input[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	// verify header
	if len(ciphertext) < len(aesCFBHeader) || !bytes.Equal(ciphertext[:len(aesCFBHeader)], aesCFBHeader) {
		return nil, fmt.Errorf("invalid ciphertext or header mismatch")
	}
	plain := ciphertext[len(aesCFBHeader):]
	if plain == nil {
		return []byte{}, nil
	}
	return plain, nil
}
