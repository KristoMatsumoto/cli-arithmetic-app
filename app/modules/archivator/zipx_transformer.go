package archivator

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
)

type ZIPXTransformer struct{}

func NewZIPXTransformer() *ZIPXTransformer {
	return &ZIPXTransformer{}
}

func (archivator *ZIPXTransformer) Name() string { return "zipx" }

func (archivator *ZIPXTransformer) Encode(input []byte) ([]byte, error) {
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

func (archivator *ZIPXTransformer) Decode(input []byte) ([]byte, error) {
	r, err := zip.NewReader(bytes.NewReader(input), int64(len(input)))
	if err != nil {
		return nil, err
	}
	if len(r.File) == 0 {
		return nil, fmt.Errorf("zipx archive empty")
	}
	f := r.File[0]
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return io.ReadAll(rc)
}
