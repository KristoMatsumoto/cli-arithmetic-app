package regex_processor

import (
	"regexp"
	"strings"
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
