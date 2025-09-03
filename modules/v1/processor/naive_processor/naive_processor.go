package naive_processor

import (
	"cli-arithmetic-app/modules/v1/processor"
	"strings"
)

type NaiveProcessor struct{}

func NewNaiveProcessor() *NaiveProcessor {
	return &NaiveProcessor{}
}

func (p *NaiveProcessor) Process(lines []string) ([]string, error) {
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
