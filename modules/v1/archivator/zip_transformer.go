package archivator

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
)

type ZIPTransformer struct{}

func NewZIPTransformer() *ZIPTransformer {
	return &ZIPTransformer{}
}

func (archivator *ZIPTransformer) Name() string { return "zip" }

func (archivator *ZIPTransformer) Encode(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := zip.NewWriter(&buf)
	f, err := w.Create("data")
	if err != nil {
		return nil, err
	}
	if _, err := f.Write(input); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (archivator *ZIPTransformer) Decode(input []byte) ([]byte, error) {
	r, err := zip.NewReader(bytes.NewReader(input), int64(len(input)))
	if err != nil {
		return nil, err
	}
	if len(r.File) == 0 {
		return nil, fmt.Errorf("zip archive empty")
	}
	f := r.File[0]
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return io.ReadAll(rc)
}
