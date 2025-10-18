package encryptor

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/chacha20"
)

var chacha20Header = []byte("ENCv1:")

type ChaCha20Transformer struct {
	key []byte
}

func NewChaCha20Transformer() (*ChaCha20Transformer, error) {
	k := []byte(os.Getenv("SECRET_KEY_32"))
	if len(k) != chacha20.KeySize {
		return nil, fmt.Errorf("SECRET KEY must be %d bytes for ChaCha20", chacha20.KeySize)
	}
	return &ChaCha20Transformer{key: k}, nil
}

func (encryptor *ChaCha20Transformer) Name() string { return "chacha20" }

func (encryptor *ChaCha20Transformer) Encode(data []byte) ([]byte, error) {
	// generate a new nonce per message
	nonce := make([]byte, chacha20.NonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	plain := append(chacha20Header, data...)

	cipher, err := chacha20.NewUnauthenticatedCipher(encryptor.key, nonce)
	if err != nil {
		return nil, err
	}
	out := make([]byte, len(plain))
	cipher.XORKeyStream(out, plain)
	return append(nonce, out...), nil
}

func (encryptor *ChaCha20Transformer) Decode(data []byte) ([]byte, error) {
	if len(data) < chacha20.NonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	nonce, ct := data[:chacha20.NonceSize], data[chacha20.NonceSize:]
	cipher, err := chacha20.NewUnauthenticatedCipher(encryptor.key, nonce)
	if err != nil {
		return nil, err
	}
	out := make([]byte, len(ct))
	cipher.XORKeyStream(out, ct)

	// verify header
	if len(out) < len(chacha20Header) || !bytes.Equal(out[:len(chacha20Header)], chacha20Header) {
		return nil, fmt.Errorf("invalid ciphertext or header mismatch")
	}
	plain := out[len(chacha20Header):]
	if plain == nil {
		return []byte{}, nil
	}
	return plain, nil
}
