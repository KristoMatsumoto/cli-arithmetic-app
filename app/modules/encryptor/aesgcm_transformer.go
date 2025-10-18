package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

type AESGCMTransformer struct {
	key []byte
}

func NewAESGCMTransformer() (*AESGCMTransformer, error) {
	k := []byte(os.Getenv("SECRET_KEY_16"))
	switch len(k) {
	case 16, 24, 32:
	default:
		return nil, fmt.Errorf("SECRET_KEY must be 16, 24, or 32 bytes")
	}
	return &AESGCMTransformer{key: k}, nil
}

func (encryptor *AESGCMTransformer) Name() string { return "aes-gcm" }

func (encryptor *AESGCMTransformer) Encode(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(encryptor.key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return append(nonce, aesgcm.Seal(nil, nonce, plaintext, nil)...), nil
}

func (encryptor *AESGCMTransformer) Decode(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(encryptor.key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aesgcm.NonceSize() {
		return nil, fmt.Errorf("ciphertext too short")
	}
	nonce, data := ciphertext[:aesgcm.NonceSize()], ciphertext[aesgcm.NonceSize():]
	plain, err := aesgcm.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, err
	}
	if plain == nil {
		return []byte{}, nil
	}
	return plain, nil
}
