package archivator

import (
	"bytes"
	"compress/gzip"
	"io"
)

type GZIPTransformer struct{}

func NewGZIPTransformer() *GZIPTransformer {
	return &GZIPTransformer{}
}

func (archivator *GZIPTransformer) Name() string { return "gzip" }

func (archivator *GZIPTransformer) Encode(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	if _, err := w.Write(input); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (archivator *GZIPTransformer) Decode(input []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return io.ReadAll(r)
}
