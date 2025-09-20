package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

type AESTransformer struct {
	key []byte
}

func NewAESTransformer() (*AESTransformer, error) {
	key := os.Getenv("SECRET_KEY")
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

	ciphertext := make([]byte, aes.BlockSize+len(input))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], input)

	return ciphertext, nil
}

func (encryptor *AESTransformer) Decode(input []byte) ([]byte, error) {
	block, err := aes.NewCipher(encryptor.key)
	if err != nil {
		return nil, err
	}

	if len(input) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := input[:aes.BlockSize]
	ciphertext := input[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}
