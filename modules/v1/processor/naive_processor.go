package processor

import (
	"fmt"
	"strconv"
	"strings"
)

type NaiveProcessor struct{}

func NewNaiveProcessor() *NaiveProcessor {
	return &NaiveProcessor{}
}

func (p *NaiveProcessor) Process(lines []string) ([]string, error) {
	var results []string

	for _, line := range lines {
		modified := line
		tokens := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' ' || r == ',' || r == '.' || r == ':' || r == ';'
		})

		for _, token := range tokens {
			val, err := evaluate(token)
			if err == nil {
				modified = strings.Replace(modified, token, val, 1)
			}
		}
		results = append(results, modified)
	}

	return results, nil
}

func evaluate(expr string) (string, error) {
	expr = strings.ReplaceAll(expr, " ", "")

	opIndex := -1
	for i, r := range expr {
		if r == '+' || r == '-' || r == '*' || r == '/' {
			opIndex = i
			break
		}
	}
	if opIndex == -1 {
		return "0", fmt.Errorf("no operation found")
	}

	aStr := expr[:opIndex]
	bStr := expr[opIndex+1:]
	operator := expr[opIndex]

	a, err1 := strconv.ParseFloat(aStr, 64)
	b, err2 := strconv.ParseFloat(bStr, 64)
	if err1 != nil || err2 != nil {
		return "0", fmt.Errorf("invalid numbers")
	}

	switch operator {
	case '+':
		return fmt.Sprintf("%v", a+b), nil
	case '-':
		return fmt.Sprintf("%v", a-b), nil
	case '*':
		return fmt.Sprintf("%v", a*b), nil
	case '/':
		if b == 0 {
			return "NaN", nil
		}
		return fmt.Sprintf("%v", a/b), nil
	default:
		return "0", fmt.Errorf("unsupported operator")
	}
}
