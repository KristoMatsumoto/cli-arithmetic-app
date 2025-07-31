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
	for i < len(expr) {
		ch := expr[i]

		switch {
		case unicode.IsSpace(rune(ch)):
			i++

		case ch == '(':
			tokens = append(tokens, Token{LParen, "("})
			i++

		case ch == ')':
			tokens = append(tokens, Token{RParen, ")"})
			i++

		case strings.ContainsRune("+-*/%^", rune(ch)):
			tokens = append(tokens, Token{Operator, string(ch)})
			i++

		case unicode.IsDigit(rune(ch)) || ch == '.':
			start := i
			for i < len(expr) && (unicode.IsDigit(rune(expr[i])) || expr[i] == '.') {
				i++
			}
			tokens = append(tokens, Token{Number, expr[start:i]})

		default:
			return nil, fmt.Errorf("invalid character: %c", ch)
		}
	}
	// for _, token := range tokens {
	// 	fmt.Print(token.Type, " ", token.Value, "\n")
	// }
	return tokens, nil
}
