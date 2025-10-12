package archivator

import (
	"github.com/klauspost/compress/zstd"
)

type ZSTDTransformer struct{}

func NewZSTDTransformer() *ZSTDTransformer {
	return &ZSTDTransformer{}
}

func (archivator *ZSTDTransformer) Name() string { return "zstd" }

func (archivator *ZSTDTransformer) Encode(input []byte) ([]byte, error) {
	enc, _ := zstd.NewWriter(nil)
	defer enc.Close()

	out := enc.EncodeAll(input, nil)
	if out == nil {
		return []byte{}, nil
	}
	return out, nil
}

func (archivator *ZSTDTransformer) Decode(input []byte) ([]byte, error) {
	dec, _ := zstd.NewReader(nil)
	defer dec.Close()

	out, err := dec.DecodeAll(input, nil)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return []byte{}, nil
	}
	return out, nil
}
