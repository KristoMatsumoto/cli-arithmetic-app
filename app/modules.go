package app

import (
	"cli-arithmetic-app/modules/v1/archivator"
	"cli-arithmetic-app/modules/v1/encryptor"
	"cli-arithmetic-app/modules/v1/parser"
	"cli-arithmetic-app/modules/v1/processor"
	"cli-arithmetic-app/modules/v1/processor/lib_processor"
	"cli-arithmetic-app/modules/v1/processor/naive_processor"
	"cli-arithmetic-app/modules/v1/processor/regex_processor"
	"cli-arithmetic-app/modules/v1/transformer"

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
	"zip": func() (transformer.Transformer, error) { return archivator.NewZIPTransformer(), nil },

	"aes": func() (transformer.Transformer, error) { return encryptor.NewAESTransformer() },
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
