package processor

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func ExtractExpressions(line string) []string {
	var tokens []string
	var exprBuilder strings.Builder
	parensCount := 0
	inExpr := false

	isExprRune := func(r rune) bool {
		return unicode.IsDigit(r) || unicode.IsSpace(r) || strings.ContainsRune("+-*/%^().", r)
	}

	for i, r := range line {
		if isExprRune(r) {
			switch r {
			case '(':
				parensCount++
			case ')':
				parensCount--
			}
			exprBuilder.WriteRune(r)
			inExpr = true
		} else {
			if inExpr && parensCount <= 0 {
				expr := strings.TrimSpace(exprBuilder.String())
				if len(expr) > 0 {
					tokens = append(tokens, expr)
				}
				exprBuilder.Reset()
				inExpr = false
				parensCount = 0
			}
		}

		// If end of line reached and still building expression
		if i == len(line)-1 && inExpr {
			expr := strings.TrimSpace(exprBuilder.String())
			if len(expr) > 0 {
				tokens = append(tokens, expr)
			}
		}
	}

	return tokens
}

func ReplaceFirst(line, old, new string) string {
	idx := strings.Index(line, old)
	if idx == -1 {
		return line
	}
	return line[:idx] + new + line[idx+len(old):]
}

func FormatFloat(f float64) string {
	if math.Mod(f, 1) == 0 {
		return fmt.Sprintf("%.0f", f)
	}
	return fmt.Sprintf("%.2f", f)
}

type NaiveProcessor struct{}

func NewNaiveProcessor() *NaiveProcessor {
	return &NaiveProcessor{}
}

func (p *NaiveProcessor) Process(lines []string) ([]string, error) {
	var results []string

	for _, line := range lines {
		original := line
		expressions := ExtractExpressions(line)

		for _, expr := range expressions {
			val, err := EvalExpression(expr)
			var replacement string

			if err != nil {
				replacement = "NaN"
			} else {
				replacement = FormatFloat(val)
			}

			original = ReplaceFirst(original, expr, replacement)
		}
		results = append(results, original)
	}

	return results, nil
}
