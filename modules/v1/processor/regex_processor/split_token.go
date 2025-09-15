package regex_processor

import (
	"regexp"
	"strings"
)

type PartType string

const TextPart PartType = "text"
const ExprPart PartType = "expr"

type TokenPart struct {
	Type  PartType
	Value string
}

// var exprCandidateRegex = regexp.MustCompile(`(\d+(\.\d+)?|\([^\(\)]+\))(\s*[+\-*/^%]\s*(\d+(\.\d+)?|\([^\(\)]+\)))*`)
// (\d+(\.\d+)?)
// [+\-*/^%]
// ([+\-]\s*)*(\d+(\.\d+)?)
// \(\)
var exprCandidateRegex = regexp.MustCompile(`(([+\-(]|(\d+(\.\d+)?))(\s*([+\-*/^%()]|(\d+(\.\d+)?)))+)`)

func SplitIntoTokens(input string) []TokenPart {
	var tokens []TokenPart
	last := 0

	for _, loc := range exprCandidateRegex.FindAllStringIndex(input, -1) {
		start, end := loc[0], loc[1]

		// Text before candidate for expression
		// if start > last {
		// 	tokens = append(tokens, TokenPart{Type: TextPart, Value: input[last:start]})
		// }

		candidate := input[start:end]
		trimmedCandidate := strings.TrimSpace(candidate)
		if isValidExpression(trimmedCandidate) {
			if start > last {
				tokens = append(tokens, TokenPart{Type: TextPart, Value: input[last:start]})
			}
			tokens = append(tokens, TokenPart{Type: ExprPart, Value: trimmedCandidate})
			last = end
		}
	}

	if last < len(input) {
		tokens = append(tokens, TokenPart{Type: TextPart, Value: input[last:]})
	}

	return tokens
}

var numberRegex = regexp.MustCompile(`\d+(\.\d+)?`)
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
