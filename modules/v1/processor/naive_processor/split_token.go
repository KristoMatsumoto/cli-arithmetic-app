package naive_processor

import (
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
