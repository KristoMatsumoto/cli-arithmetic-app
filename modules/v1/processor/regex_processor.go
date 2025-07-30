package processor

import (
	"regexp"
	"strconv"
)

type RegexProcessor struct{}

func NewRegexProcessor() Processor {
	return &RegexProcessor{}
}

func (p *RegexProcessor) Process(lines []string) ([]string, error) {
	var results []string

	// Находит выражения вида "число оператор число", например: 3+4 или  12 * -3
	re := regexp.MustCompile(`(-?\d+)\s*([\+\-\*/])\s*(-?\d+)`)

	for _, line := range lines {
		processed := re.ReplaceAllStringFunc(line, func(expr string) string {
			match := re.FindStringSubmatch(expr)
			if len(match) != 4 {
				return expr
			}
			left, _ := strconv.Atoi(match[1])
			op := match[2]
			right, _ := strconv.Atoi(match[3])

			var result int
			switch op {
			case "+":
				result = left + right
			case "-":
				result = left - right
			case "*":
				result = left * right
			case "/":
				if right == 0 {
					return "NaN"
				}
				result = left / right
			}
			return strconv.Itoa(result)
		})
		results = append(results, processed)
	}

	return results, nil
}
