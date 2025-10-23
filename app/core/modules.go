package core

import (
	"cli-arithmetic-app/app/modules/archivator"
	"cli-arithmetic-app/app/modules/encryptor"
	"cli-arithmetic-app/app/modules/parser"
	"cli-arithmetic-app/app/modules/processor"
	"cli-arithmetic-app/app/modules/processor/lib_processor"
	"cli-arithmetic-app/app/modules/processor/naive_processor"
	"cli-arithmetic-app/app/modules/processor/regex_processor"
	"cli-arithmetic-app/app/modules/transformer"

	"fmt"
)

// =========================================================
//                         FACTORIES
// =========================================================

// ----------------------- Processor -----------------------
var processorRegistry = map[string]func() processor.Processor{
	"1": func() processor.Processor { return naive_processor.NewNaiveProcessor() },
	"2": func() processor.Processor { return regex_processor.NewRegexProcessor() },
	"3": func() processor.Processor { return lib_processor.NewLibProcessor() },

	"naive": func() processor.Processor { return naive_processor.NewNaiveProcessor() },
	"regex": func() processor.Processor { return regex_processor.NewRegexProcessor() },
	"lib":   func() processor.Processor { return lib_processor.NewLibProcessor() },
}

func CreateProcessor(mode string) (processor.Processor, error) {
	if factory, ok := processorRegistry[mode]; ok {
		return factory(), nil
	}
	return nil, fmt.Errorf("unsupported processor mode: %s", mode)
}

// ------------------------ Parser -------------------------
var parserRegistry = map[string]func() parser.Parser{
	"text": func() parser.Parser { return parser.NewTextParser() },
	"txt":  func() parser.Parser { return parser.NewTextParser() },
	"json": func() parser.Parser { return parser.NewJSONParser() },
	"xml":  func() parser.Parser { return parser.NewXMLParser() },
	"yaml": func() parser.Parser { return parser.NewYAMLParser() },
	"html": func() parser.Parser { return parser.NewHTMLParser() },
}

func CreateParser(format string) (parser.Parser, error) {
	if factory, ok := parserRegistry[format]; ok {
		return factory(), nil
	}
	return nil, fmt.Errorf("unsupported parser format: %s", format)
}

// -------- Transformers (archivators + encryptors) --------
var transformerRegistry = map[string]func() (transformer.Transformer, error){
	"brotli": func() (transformer.Transformer, error) { return archivator.NewBrotliTransformer(), nil },
	"gzip":   func() (transformer.Transformer, error) { return archivator.NewGZIPTransformer(), nil },
	"lz4":    func() (transformer.Transformer, error) { return archivator.NewLZ4Transformer(), nil },
	"tar":    func() (transformer.Transformer, error) { return archivator.NewTARTransformer(), nil },
	"zip":    func() (transformer.Transformer, error) { return archivator.NewZIPTransformer(), nil },
	"zipx":   func() (transformer.Transformer, error) { return archivator.NewZIPXTransformer(), nil },
	"zstd":   func() (transformer.Transformer, error) { return archivator.NewZSTDTransformer(), nil },

	"3des":              func() (transformer.Transformer, error) { return encryptor.NewTripleDESTransformer() },
	"aes":               func() (transformer.Transformer, error) { return encryptor.NewAESTransformer() },
	"aes-cbc":           func() (transformer.Transformer, error) { return encryptor.NewAESCBCTransformer() },
	"aes-gcm":           func() (transformer.Transformer, error) { return encryptor.NewAESGCMTransformer() },
	"blowfish":          func() (transformer.Transformer, error) { return encryptor.NewBlowfishTransformer() },
	"chacha20":          func() (transformer.Transformer, error) { return encryptor.NewChaCha20Transformer() },
	"chacha20-poly1305": func() (transformer.Transformer, error) { return encryptor.NewChaCha20Poly1305Transformer() },
	"gost-28147":        func() (transformer.Transformer, error) { return encryptor.NewGOST28147Transformer() },
	"rc4":               func() (transformer.Transformer, error) { return encryptor.NewRC4Transformer() },
	"xor":               func() (transformer.Transformer, error) { return encryptor.NewXORTransformer() },
}

func CreateTransformer(name string) (transformer.Transformer, error) {
	if factory, ok := transformerRegistry[name]; ok {
		return factory()
	}
	return nil, fmt.Errorf("unsupported transformer: %s", name)
}

func BuildTransformChain(chain []string) ([]transformer.Transformer, error) {
	var transformers []transformer.Transformer
	for _, name := range chain {
		ctor, ok := transformerRegistry[name]
		if !ok {
			return nil, fmt.Errorf("unknown transformer: %s", name)
		}
		t, err := ctor()
		if err != nil {
			return nil, fmt.Errorf("failed to init transformer %s: %w", name, err)
		}
		transformers = append(transformers, t)
	}
	return transformers, nil
}
