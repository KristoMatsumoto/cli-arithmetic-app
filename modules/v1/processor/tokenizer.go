package processor

import (
	"fmt"
	"strings"
	"unicode"
)

type tokenType int

const (
	Number tokenType = iota
	Operator
	LParen
	RParen
)

type Token struct {
	Type  tokenType
	Value string
}

func Tokenize(expr string) ([]Token, error) {
	var tokens []Token
	i := 0
	prevType := Operator

	for i < len(expr) {
		ch := expr[i]

		switch {
		case unicode.IsSpace(rune(ch)):
			i++

		case ch == '(':
			tokens = append(tokens, Token{LParen, "("})
			prevType = LParen
			i++

		case ch == ')':
			tokens = append(tokens, Token{RParen, ")"})
			prevType = RParen
			i++

		case strings.ContainsRune("+-*/%^", rune(ch)):
			if ch == '+' {
				if prevType == Operator || prevType == LParen {
					i++
					continue
				}
				tokens = append(tokens, Token{Operator, string(ch)})
			} else if ch == '-' {
				if prevType == Operator || prevType == LParen {
					tokens = append(tokens, Token{Operator, "UMINUS"})
				} else {
					tokens = append(tokens, Token{Operator, "-"})
				}
			} else {
				tokens = append(tokens, Token{Operator, string(ch)})
			}
			prevType = Operator
			i++

		case unicode.IsDigit(rune(ch)) || ch == '.':
			start := i
			for i < len(expr) && (unicode.IsDigit(rune(expr[i])) || expr[i] == '.') {
				i++
			}
			tokens = append(tokens, Token{Number, expr[start:i]})
			prevType = Number

		default:
			return nil, fmt.Errorf("invalid character: %c", ch)
		}
	}

	return tokens, nil
}
