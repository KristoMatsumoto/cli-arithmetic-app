package encryptor

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

var tripleDESHeader = []byte("ENCv1:")

type TripleDESTransformer struct {
	key []byte
}

func NewTripleDESTransformer() (*TripleDESTransformer, error) {
	k := []byte(os.Getenv("SECRET_KEY_24"))
	if len(k) != 24 {
		return nil, fmt.Errorf("SECRET KEY must be 24 bytes for 3DES")
	}
	return &TripleDESTransformer{key: k}, nil
}

func (encryptor *TripleDESTransformer) Name() string { return "3des" }

func (encryptor *TripleDESTransformer) Encode(data []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(encryptor.key)
	if err != nil {
		return nil, err
	}

	// prepend header to plaintext so Decode can validate input integrity
	plain := append(tripleDESHeader, data...)

	ciphertext := make([]byte, des.BlockSize+len(plain))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[des.BlockSize:], plain)
	return ciphertext, nil
}

func (encryptor *TripleDESTransformer) Decode(data []byte) ([]byte, error) {
	if len(data) < des.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	block, err := des.NewTripleDESCipher(encryptor.key)
	if err != nil {
		return nil, err
	}

	iv := data[:des.BlockSize]
	ct := data[des.BlockSize:]
	// decrypt in-place
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ct, ct)

	// verify header
	if len(ct) < len(tripleDESHeader) || !bytes.Equal(ct[:len(tripleDESHeader)], tripleDESHeader) {
		return nil, fmt.Errorf("invalid ciphertext or header mismatch")
	}
	// return the original plaintext without header
	plain := ct[len(tripleDESHeader):]
	// normalize empty slice (non-nil)
	if plain == nil {
		return []byte{}, nil
	}
	return plain, nil
}
