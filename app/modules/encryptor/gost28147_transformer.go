package encryptor

import (
	"cli-arithmetic-app/app/config"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/ddulesov/gogost/gost28147"
)

type GOST28147Transformer struct {
	key []byte
}

func NewGOST28147Transformer() (*GOST28147Transformer, error) {
	config.GetConfig()
	k := []byte(config.C.SecretKeys.S_32)
	if len(k) != 32 {
		return nil, fmt.Errorf("SECRET KEY must be 32 bytes for GOST")
	}
	return &GOST28147Transformer{key: k}, nil
}

func (t *GOST28147Transformer) Name() string { return "gost-28147" }

func (t *GOST28147Transformer) Encode(data []byte) ([]byte, error) {
	block := gost28147.NewCipher(t.key, gost28147.SboxDefault)
	padded := pad(data, 8)
	ciphertext := make([]byte, len(padded))
	iv := make([]byte, 8)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext, padded)
	return append(iv, ciphertext...), nil
}

func (t *GOST28147Transformer) Decode(data []byte) ([]byte, error) {
	if len(data) < 8 {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := data[:8]
	ct := data[8:]
	block := gost28147.NewCipher(t.key, gost28147.SboxDefault)
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ct, ct)
	return unpad(ct)
}
