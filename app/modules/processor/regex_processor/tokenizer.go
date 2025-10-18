package regex_processor

import (
	"strings"
	"unicode"
)

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
			if ch == '-' {
				if len(tokens) == 0 || tokens[len(tokens)-1].Type == operator || tokens[len(tokens)-1].Type == lparen {
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
