package lib_processor

// To generate parser using pigeon run command:
//go:generate pigeon -o split_token.go expression.peg

import (
	"cli-arithmetic-app/app/modules/processor"
	"math"
	"strings"
)

type LibProcessor struct{}

func NewLibProcessor() *LibProcessor {
	return &LibProcessor{}
}

func (p *LibProcessor) Process(lines []string) ([]string, error) {
	var results []string

	for _, line := range lines {
		parsed, err := Parse("", []byte(line))
		if err != nil {
			// парсинг не удался — возвращаем строку как есть
			results = append(results, line)
			continue
		}

		parts := parsed.([]interface{})
		var rebuilt strings.Builder

		for _, part := range parts {
			m := part.(map[string]string)
			if m["Type"] == "expr" {
				val, _ := EvalExpression(m["Value"])
				if math.IsNaN(val) {
					rebuilt.WriteString("NaN")
				} else {
					rebuilt.WriteString(processor.FormatFloat(val))
				}
			} else {
				rebuilt.WriteString(m["Value"])
			}
		}

		results = append(results, rebuilt.String())
	}

	return results, nil
}
