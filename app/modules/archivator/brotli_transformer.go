package archivator

import (
	"bytes"
	"io"

	"github.com/andybalholm/brotli"
)

type BrotliTransformer struct{}

func NewBrotliTransformer() *BrotliTransformer {
	return &BrotliTransformer{}
}

func (archivator *BrotliTransformer) Name() string { return "brotli" }

func (archivator *BrotliTransformer) Encode(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := brotli.NewWriter(&buf)
	if _, err := w.Write(input); err != nil {
		return nil, err
	}
	w.Close()
	return buf.Bytes(), nil
}

func (archivator *BrotliTransformer) Decode(input []byte) ([]byte, error) {
	r := brotli.NewReader(bytes.NewReader(input))
	return io.ReadAll(r)
}
