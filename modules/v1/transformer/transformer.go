package transformer

// Transformer is used for coding or decoding files or archives
type Transformer interface {
	Name() string
	Encode(input []byte) ([]byte, error)
	Decode(input []byte) ([]byte, error)
}
