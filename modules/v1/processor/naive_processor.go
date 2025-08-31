package processor

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type PartType int

const (
	TextPart PartType = iota
	ExprPart
)

type TokenPart struct {
	Type  PartType
	Value string
}

func SplitIntoTokens(input string) []TokenPart {
	var tokens []TokenPart
	buf := []rune{}
	subbuf := []rune{}
	spacebuf := []rune{}
	isExpr := false

	flush := func(asExpr bool) {
		if len(buf) == 0 {
			return
		}
		val := string(buf)
		if asExpr && isValidExpression(val) {
			tokens = append(tokens, TokenPart{Type: ExprPart, Value: val})
		} else {
			tokens = append(tokens, TokenPart{Type: TextPart, Value: val})
		}
		buf = []rune{}
		isExpr = false
	}

	for i := 0; i < len(input); i++ {
		ch := rune(input[i])

		switch {
		case unicode.IsDigit(ch) || ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '^':
			if !isExpr && len(buf) > 0 {
				flush(false)
			}
			if len(subbuf) > 0 {
				buf = append(buf, subbuf...)
				subbuf = nil
			}
			if len(spacebuf) > 0 {
				buf = append(buf, spacebuf...)
				spacebuf = nil
			}
			buf = append(buf, ch)
			isExpr = true

		case ch == '(':
			subbuf = append(subbuf, ch)

		case ch == ')':
			if len(subbuf) > 0 {
				buf = append(buf, subbuf...)
				subbuf = nil
			}
			if len(spacebuf) > 0 {
				buf = append(buf, spacebuf...)
				spacebuf = nil
			}
			buf = append(buf, ch)

		case ch == '.':
			if len(subbuf) == 0 && len(spacebuf) == 0 && isExpr {
				if len(buf) > 0 && unicode.IsDigit(buf[len(buf)-1]) {
					buf = append(buf, ch)
					continue
				} else {
					flush(isExpr)
				}
			}
			buf = append(buf, subbuf...)
			subbuf = nil
			buf = append(buf, spacebuf...)
			spacebuf = nil
			buf = append(buf, ch)
			isExpr = false

		case unicode.IsSpace(ch):
			if len(subbuf) > 0 {
				subbuf = append(subbuf, ch)
			} else {
				if isExpr {
					spacebuf = append(spacebuf, ch)
				} else {
					buf = append(buf, ch)
				}
			}

		default:
			if isExpr {
				flush(true)
			}
			if len(subbuf) > 0 {
				buf = append(buf, subbuf...)
				subbuf = nil
			}
			if len(spacebuf) > 0 {
				buf = append(buf, spacebuf...)
				spacebuf = nil
			}
			buf = append(buf, ch)
		}
	}

	flush(isExpr)
	return tokens
}

func isValidExpression(expr string) bool {
	tokens, err := Tokenize(expr)
	if err != nil {
		return false
	}

	numCount, opCount := 0, 0
	for _, t := range tokens {
		switch t.Type {
		case Number:
			numCount++
		case Operator:
			opCount++
		}
	}
	return (numCount > 1 && opCount > 0) || (numCount > 0 && opCount > 1)
}

func EvalSimple(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	case "%":
		if b == 0 {
			return 0, fmt.Errorf("modulo by zero")
		}
		return math.Mod(a, b), nil
	case "^":
		return math.Pow(a, b), nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", op)
	}
}

func precedence(op string) int {
	switch op {
	case "UMINUS":
		return 5
	case "^":
		return 4
	case "*", "/", "%":
		return 3
	case "+", "-":
		return 2
	default:
		return 0
	}
}

func isRightAssociative(op string) bool {
	return op == "^" || op == "UMINUS"
}

func EvalExpression(expr string) (float64, error) {
	tokens, err := Tokenize(expr)
	if err != nil {
		return 0, err
	}

	var values []float64
	var ops []string

	// applies the upper operator from the stack to the values
	apply := func() error {
		if len(ops) == 0 {
			return fmt.Errorf("no operators to apply")
		}
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]

		if op == "UMINUS" {
			if len(values) < 1 {
				return fmt.Errorf("not enough values for unary minus")
			}
			values[len(values)-1] = -values[len(values)-1]
			return nil
		}

		if len(values) < 2 {
			return fmt.Errorf("not enough values for binary operator %q", op)
		}
		b := values[len(values)-1]
		a := values[len(values)-2]
		values = values[:len(values)-2]

		res, err := EvalSimple(a, b, op)
		if err != nil {
			return err
		}
		values = append(values, res)
		return nil
	}

	for _, tok := range tokens {
		switch tok.Type {
		case Number:
			num, err := strconv.ParseFloat(tok.Value, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number %q: %w", tok.Value, err)
			}
			values = append(values, num)

		case Operator:
			opVal := tok.Value

			for len(ops) > 0 {
				top := ops[len(ops)-1]
				if top == "(" {
					break
				}
				if (isRightAssociative(opVal) && precedence(opVal) < precedence(top)) ||
					(!isRightAssociative(opVal) && precedence(opVal) <= precedence(top)) {
					if err := apply(); err != nil {
						return 0, err
					}
				} else {
					break
				}
			}
			ops = append(ops, opVal)

		case LParen:
			ops = append(ops, "(")

		case RParen:
			for len(ops) > 0 && ops[len(ops)-1] != "(" {
				if err := apply(); err != nil {
					return 0, err
				}
			}
			if len(ops) == 0 {
				return 0, fmt.Errorf("mismatched parenthesis")
			}
			ops = ops[:len(ops)-1]
		}
	}

	for len(ops) > 0 {
		if ops[len(ops)-1] == "(" {
			return 0, fmt.Errorf("mismatched parenthesis")
		}
		if err := apply(); err != nil {
			return 0, err
		}
	}

	if len(values) != 1 {
		return 0, fmt.Errorf("invalid result")
	}
	return values[0], nil
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
		parts := SplitIntoTokens(line)

		var rebuilt strings.Builder
		for _, part := range parts {
			if part.Type == ExprPart {
				val, err := EvalExpression(part.Value)
				if err != nil {
					rebuilt.WriteString("NaN")
				} else {
					rebuilt.WriteString(FormatFloat(val))
				}
			} else {
				rebuilt.WriteString(part.Value)
			}
		}
		results = append(results, rebuilt.String())
	}
	return results, nil
}
