package app

import (
	"cli-arithmetic-app/modules/v1/parser"
	"cli-arithmetic-app/modules/v1/processor"
	"cli-arithmetic-app/modules/v1/processor/lib_processor"
	"cli-arithmetic-app/modules/v1/processor/naive_processor"
	"cli-arithmetic-app/modules/v1/processor/regex_processor"

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
