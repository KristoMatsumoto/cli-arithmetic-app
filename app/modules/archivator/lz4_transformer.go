package archivator

import (
	"bytes"
	"io"

	"github.com/pierrec/lz4/v4"
)

type LZ4Transformer struct{}

func NewLZ4Transformer() *LZ4Transformer {
	return &LZ4Transformer{}
}

func (archivator *LZ4Transformer) Name() string { return "lz4" }

func (archivator *LZ4Transformer) Encode(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := lz4.NewWriter(&buf)
	if _, err := w.Write(input); err != nil {
		return nil, err
	}
	w.Close()
	return buf.Bytes(), nil
}

func (archivator *LZ4Transformer) Decode(input []byte) ([]byte, error) {
	r := lz4.NewReader(bytes.NewReader(input))
	return io.ReadAll(r)
}
