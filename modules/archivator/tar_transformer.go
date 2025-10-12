package archivator

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
)

type TARTransformer struct{}

func NewTARTransformer() *TARTransformer {
	return &TARTransformer{}
}

func (archivator *TARTransformer) Name() string { return "tar" }

func (archivator *TARTransformer) Encode(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	hdr := &tar.Header{
		Name: "data",
		Mode: 0600,
		Size: int64(len(input)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return nil, err
	}
	if _, err := tw.Write(input); err != nil {
		return nil, err
	}
	if err := tw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (archivator *TARTransformer) Decode(input []byte) ([]byte, error) {
	tr := tar.NewReader(bytes.NewReader(input))
	hdr, err := tr.Next()
	if err != nil {
		return nil, err
	}
	if hdr == nil {
		return nil, fmt.Errorf("tar archive empty")
	}
	return io.ReadAll(tr)
}
