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

// AESCBCTransformer — AES в режиме CBC с PKCS7 паддингом.
// Формат выходных данных: IV (aes.BlockSize == 16) || ciphertext
type AESCBCTransformer struct {
	key []byte
}

func NewAESCBCTransformer() (*AESCBCTransformer, error) {
	k := []byte(os.Getenv("SECRET_KEY_16"))
	switch len(k) {
	case 16, 24, 32:
		// ok
	default:
		return nil, fmt.Errorf("SECRET_KEY must be 16, 24, or 32 bytes for AES")
	}
	return &AESCBCTransformer{key: k}, nil
}

func (encryptor *AESCBCTransformer) Name() string { return "aes-cbc" }

func (encryptor *AESCBCTransformer) Encode(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(encryptor.key)
	if err != nil {
		return nil, err
	}
	blockSize := aes.BlockSize
	padded := pkcs7Pad(plaintext, blockSize)

	iv := make([]byte, blockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	ct := make([]byte, len(padded))
	mode.CryptBlocks(ct, padded)
	return append(iv, ct...), nil
}

func (encryptor *AESCBCTransformer) Decode(payload []byte) ([]byte, error) {
	block, err := aes.NewCipher(encryptor.key)
	if err != nil {
		return nil, err
	}
	blockSize := aes.BlockSize
	if len(payload) < blockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := payload[:blockSize]
	ct := payload[blockSize:]
	if len(ct)%blockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plainPadded := make([]byte, len(ct))
	mode.CryptBlocks(plainPadded, ct)

	return pkcs7Unpad(plainPadded, blockSize)
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	if blockSize <= 0 {
		return data
	}
	pad := blockSize - (len(data) % blockSize)
	if pad == 0 {
		pad = blockSize
	}
	padding := bytes.Repeat([]byte{byte(pad)}, pad)
	return append(data, padding...)
}

func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 || len(data)%blockSize != 0 {
		return nil, fmt.Errorf("invalid padded data")
	}
	pad := int(data[len(data)-1])
	if pad == 0 || pad > blockSize {
		return nil, fmt.Errorf("invalid padding")
	}
	for i := 1; i <= pad; i++ {
		if data[len(data)-i] != byte(pad) {
			return nil, fmt.Errorf("invalid padding")
		}
	}
	return data[:len(data)-pad], nil
}
