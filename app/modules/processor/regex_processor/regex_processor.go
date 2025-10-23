package regex_processor

import (
	"cli-arithmetic-app/app/modules/processor"
	"strings"
)

type RegexProcessor struct{}

func NewRegexProcessor() *RegexProcessor {
	return &RegexProcessor{}
}

func (p *RegexProcessor) Process(lines []string) ([]string, error) {
	var results []string
	for _, line := range lines {
		parts := SplitIntoTokens(line)

		var rebuilt strings.Builder
		for _, part := range parts {
			if part.Type == ExprPart {
				val, err := EvalExpression(part.Value)
				if err != nil {
					rebuilt.WriteString("NaN")
				} else {
					rebuilt.WriteString(processor.FormatFloat(val))
				}
			} else {
				rebuilt.WriteString(part.Value)
			}
		}
		results = append(results, rebuilt.String())
	}
	return results, nil
}
