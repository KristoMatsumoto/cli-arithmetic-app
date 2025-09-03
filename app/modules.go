package app

import (
	"cli-arithmetic-app/modules/v1/parser"
	"cli-arithmetic-app/modules/v1/processor"
	"cli-arithmetic-app/modules/v1/processor/naive_processor"
	"cli-arithmetic-app/modules/v1/processor/regex_processor"

	"fmt"
)

func CreateProcessor(mode string) (processor.Processor, error) {
	switch mode {
	case "1":
		return naive_processor.NewNaiveProcessor(), nil
	case "2":
		return regex_processor.NewRegexProcessor(), nil
	default:
		return nil, fmt.Errorf("unsupported processor mode: %s", mode)
	}
}

func CreateParser(format string) (parser.Parser, error) {
	switch format {
	case "text":
		return parser.NewTextParser(), nil
	case "txt":
		return parser.NewTextParser(), nil
	case "json":
		return parser.NewJSONParser(), nil
	case "xml":
		return parser.NewXMLParser(), nil
	case "yaml":
		return parser.NewYAMLParser(), nil
	case "html":
		return parser.NewHTMLParser(), nil
	default:
		return nil, fmt.Errorf("unsupported format")
	}
}
