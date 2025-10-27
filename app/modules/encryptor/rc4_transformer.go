package encryptor

import (
	"bytes"
	"cli-arithmetic-app/app/config"
	"crypto/rc4"
	"fmt"
)

var rc4Header = []byte("ENCv1:")

type RC4Transformer struct {
	key []byte
}

func NewRC4Transformer() (*RC4Transformer, error) {
	config.GetConfig()
	k := []byte(config.C.SecretKeys.S_16)
	if len(k) < 1 || len(k) > 256 {
		return nil, fmt.Errorf("SECRET KEY must be 1-256 bytes for RC4")
	}
	return &RC4Transformer{key: k}, nil
}

func (encryptor *RC4Transformer) Name() string { return "rc4" }

func (encryptor *RC4Transformer) Encode(data []byte) ([]byte, error) {
	c, err := rc4.NewCipher(encryptor.key)
	if err != nil {
		return nil, err
	}
	plain := append(rc4Header, data...)
	out := make([]byte, len(plain))
	c.XORKeyStream(out, plain)
	return out, nil
}

func (encryptor *RC4Transformer) Decode(data []byte) ([]byte, error) {
	c, err := rc4.NewCipher(encryptor.key)
	if err != nil {
		return nil, err
	}
	out := make([]byte, len(data))
	c.XORKeyStream(out, data)

	if len(out) < len(rc4Header) || !bytes.Equal(out[:len(rc4Header)], rc4Header) {
		return nil, fmt.Errorf("invalid ciphertext or header mismatch")
	}
	plain := out[len(rc4Header):]
	if plain == nil {
		return []byte{}, nil
	}
	return plain, nil
}
