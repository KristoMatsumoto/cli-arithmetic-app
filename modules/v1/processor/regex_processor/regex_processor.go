package regex_processor

import (
	"fmt"
	"math"
	"regexp"
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

// Не режет слова но еще не доработано
// var exprCandidateRegex = regexp.MustCompile(`-?\s*(?:-?\d+(?:\.\d+)?|-?\([^\(\)]*\))(?:\s*[+\-*/^%]\s*(?:-?\d+(?:\.\d+)?|-?\([^\(\)]*\)))*`)

// func SplitIntoTokens(input string) []TokenPart {
// 	var tokens []TokenPart
// 	last := 0

// 	for _, loc := range exprCandidateRegex.FindAllStringIndex(input, -1) {
// 		start, end := loc[0], loc[1]

// 		if start > last {
// 			tokens = append(tokens, TokenPart{Type: TextPart, Value: input[last:start]})
// 		}

// 		candidate := input[start:end]

// 		trimmedCandidate := strings.TrimSpace(candidate)
// 		if isValidExpression(trimmedCandidate) {
// 			leftSpaceLen := strings.Index(candidate, trimmedCandidate)
// 			if leftSpaceLen > 0 {
// 				tokens = append(tokens, TokenPart{Type: TextPart, Value: candidate[:leftSpaceLen]})
// 			}

// 			tokens = append(tokens, TokenPart{Type: ExprPart, Value: trimmedCandidate})

// 			rightSpace := candidate[leftSpaceLen+len(trimmedCandidate):]
// 			if rightSpace != "" {
// 				tokens = append(tokens, TokenPart{Type: TextPart, Value: rightSpace})
// 			}
// 		} else {
// 			tokens = append(tokens, TokenPart{Type: TextPart, Value: candidate})
// 		}

// 		last = end
// 	}

// 	if last < len(input) {
// 		tokens = append(tokens, TokenPart{Type: TextPart, Value: input[last:]})
// 	}

// 	return tokens
// }

var exprCandidateRegex = regexp.MustCompile(`(?:[+\-*/^%()\d.]|\s)+`)

func SplitIntoTokens(input string) []TokenPart {
	var tokens []TokenPart
	last := 0

	// Ищем все кандидаты на выражение через regex
	for _, loc := range exprCandidateRegex.FindAllStringIndex(input, -1) {
		start, end := loc[0], loc[1]

		// Текст до кандидата
		if start > last {
			tokens = append(tokens, TokenPart{Type: TextPart, Value: input[last:start]})
		}

		candidate := input[start:end]

		trimmedCandidate := strings.TrimSpace(candidate)

		if isValidExpression(trimmedCandidate) {
			leftSpaceLen := strings.Index(candidate, trimmedCandidate)
			if leftSpaceLen > 0 {
				tokens = append(tokens, TokenPart{Type: TextPart, Value: candidate[:leftSpaceLen]})
			}

			tokens = append(tokens, TokenPart{Type: ExprPart, Value: trimmedCandidate})

			rightSpace := candidate[leftSpaceLen+len(trimmedCandidate):]
			if rightSpace != "" {
				tokens = append(tokens, TokenPart{Type: TextPart, Value: rightSpace})
			}
		} else {
			tokens = append(tokens, TokenPart{Type: TextPart, Value: candidate})
		}

		last = end
	}

	if last < len(input) {
		tokens = append(tokens, TokenPart{Type: TextPart, Value: input[last:]})
	}

	return tokens
}

var numberRegex = regexp.MustCompile(`\d+(?:\.\d+)?`)
var operatorRegex = regexp.MustCompile(`[+\-*/^%]`)

func isValidExpression(s string) bool {
	numCount := len(numberRegex.FindAllString(s, -1))
	opCount := len(operatorRegex.FindAllString(s, -1))

	trimmed := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(s, "(", ""), ")", ""))
	if trimmed == "" {
		return false
	}

	return (numCount > 1 && opCount > 0) || (numCount > 0 && opCount > 1)
}

type tokenType int

const (
	number tokenType = iota
	operator
	lparen
	rparen
)

type Token struct {
	Type  tokenType
	Value string
}

func Tokenize(expr string) []Token {
	var tokens []Token
	runes := []rune(expr)
	i := 0
	for i < len(runes) {
		ch := runes[i]

		if unicode.IsSpace(ch) {
			i++
			continue
		}

		if unicode.IsDigit(ch) || ch == '.' {
			start := i
			for i < len(runes) && (unicode.IsDigit(runes[i]) || runes[i] == '.') {
				i++
			}
			tokens = append(tokens, Token{Type: number, Value: string(runes[start:i])})
			continue
		}

		if strings.ContainsRune("+-*/%^", ch) {
			// обработка унарного минуса
			if ch == '-' {
				if len(tokens) == 0 || tokens[len(tokens)-1].Type == operator || tokens[len(tokens)-1].Type == lparen {
					// унарный минус как отдельный оператор
					tokens = append(tokens, Token{Type: operator, Value: "UMINUS"})
					i++
					continue
				}
			}
			tokens = append(tokens, Token{Type: operator, Value: string(ch)})
			i++
			continue
		}

		if ch == '(' {
			tokens = append(tokens, Token{Type: lparen, Value: "("})
			i++
			continue
		}

		if ch == ')' {
			tokens = append(tokens, Token{Type: rparen, Value: ")"})
			i++
			continue
		}

		i++
	}
	return tokens
}

var precedence = map[string]int{
	"UMINUS": 3, "^": 3,
	"*": 2, "/": 2, "%": 2,
	"+": 1, "-": 1,
}

func applyOp(op string, stack *[]float64) error {
	if len(*stack) == 0 {
		return fmt.Errorf("empty stack for operator %s", op)
	}

	if op == "UMINUS" {
		val := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
		*stack = append(*stack, -val)
		return nil
	}

	if len(*stack) < 2 {
		return fmt.Errorf("not enough operands for %s", op)
	}

	b := (*stack)[len(*stack)-1]
	a := (*stack)[len(*stack)-2]
	*stack = (*stack)[:len(*stack)-2]

	var res float64
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			return fmt.Errorf("division by zero")
		}
		res = a / b
	case "%":
		if b == 0 {
			return fmt.Errorf("modulo by zero")
		}
		res = float64(int64(a) % int64(b))
	case "^":
		res = pow(a, b)
	default:
		return fmt.Errorf("unknown operator %s", op)
	}
	*stack = append(*stack, res)
	return nil
}

func pow(a, b float64) float64 {
	res := 1.0
	for i := 0; i < int(b); i++ {
		res *= a
	}
	return res
}

func EvalExpression(expr string) (float64, error) {
	tokens := Tokenize(expr)
	if len(tokens) == 0 {
		return math.NaN(), fmt.Errorf("empty expression")
	}

	var output []Token
	var ops []Token

	for _, t := range tokens {
		switch t.Type {
		case number:
			output = append(output, t)
		case operator:
			for len(ops) > 0 {
				top := ops[len(ops)-1]
				if top.Type == operator && precedence[top.Value] >= precedence[t.Value] {
					output = append(output, top)
					ops = ops[:len(ops)-1]
				} else {
					break
				}
			}
			ops = append(ops, t)
		case lparen:
			ops = append(ops, t)
		case rparen:
			for len(ops) > 0 && ops[len(ops)-1].Type != lparen {
				output = append(output, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			if len(ops) == 0 {
				return math.NaN(), fmt.Errorf("unbalanced parentheses")
			}
			ops = ops[:len(ops)-1] // убираем '('
		}
	}
	for len(ops) > 0 {
		if ops[len(ops)-1].Type == lparen {
			return math.NaN(), fmt.Errorf("unbalanced parentheses")
		}
		output = append(output, ops[len(ops)-1])
		ops = ops[:len(ops)-1]
	}

	var stack []float64
	for _, t := range output {
		switch t.Type {
		case number:
			val, err := strconv.ParseFloat(t.Value, 64)
			if err != nil {
				return math.NaN(), err
			}
			stack = append(stack, val)
		case operator:
			if err := applyOp(t.Value, &stack); err != nil {
				return math.NaN(), err
			}
		}
	}

	if len(stack) != 1 {
		return math.NaN(), fmt.Errorf("invalid expression")
	}
	return stack[0], nil
}

func FormatFloat(f float64) string {
	if math.IsNaN(f) {
		return "NaN"
	}
	if math.Mod(f, 1) == 0 {
		return fmt.Sprintf("%.0f", f)
	}
	return fmt.Sprintf("%.2f", f)
}

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
