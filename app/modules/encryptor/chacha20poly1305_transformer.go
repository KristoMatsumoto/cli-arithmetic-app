package encryptor

import (
	"cli-arithmetic-app/app/config"
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/chacha20poly1305"
)

type ChaCha20Poly1305Transformer struct {
	key []byte
}

func NewChaCha20Poly1305Transformer() (*ChaCha20Poly1305Transformer, error) {
	config.GetConfig()
	k := []byte(config.C.SecretKeys.S_32)
	if len(k) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("SECRET KEY must be %d bytes for ChaCha20-Poly1305", chacha20poly1305.KeySize)
	}
	return &ChaCha20Poly1305Transformer{key: k}, nil
}

func (encryptor *ChaCha20Poly1305Transformer) Name() string { return "chacha20-poly1305" }

func (encryptor *ChaCha20Poly1305Transformer) Encode(data []byte) ([]byte, error) {
	aead, err := chacha20poly1305.New(encryptor.key)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return append(nonce, aead.Seal(nil, nonce, data, nil)...), nil
}

func (encryptor *ChaCha20Poly1305Transformer) Decode(data []byte) ([]byte, error) {
	aead, err := chacha20poly1305.New(encryptor.key)
	if err != nil {
		return nil, err
	}
	if len(data) < aead.NonceSize() {
		return nil, fmt.Errorf("ciphertext too short")
	}
	nonce, ct := data[:aead.NonceSize()], data[aead.NonceSize():]
	plain, err := aead.Open(nil, nonce, ct, nil)
	if err != nil {
		return nil, err
	}
	if plain == nil {
		return []byte{}, nil
	}
	return plain, nil
}
