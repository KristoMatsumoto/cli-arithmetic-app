package lib_processor

import (
	"regexp"
	"strings"
)

// Deleting spaces
// Catch the sequence + and -
//		\d|\)			a digit or a closening bracket (using where chain with binary sign)
//		^|[\*\/\%\^]	beginning of the line or any sign except + and - (using where chain without binary sign)
//		[\+\-]{2,}		two or more signs + or -
//		\+|[\+\-]{2,}	unary + OR two or more signs + or - (only for unary chains)
//		\d|\(			a digit or a opening bracket (where chain of signs is ending)

var reSpace = regexp.MustCompile(`\s+`)
var reBinaryChain = regexp.MustCompile(`(\d|\))([\+\-]{2,})(\d|\()`)
var reUnaryChain = regexp.MustCompile(`(^|[\*\/\%\^])(\+|[\+\-]{2,})(\d|\()`)

func normalizeSigns(expr string) string {
	expr = reBinaryChain.ReplaceAllStringFunc(expr, func(s string) string {
		m := reBinaryChain.FindStringSubmatch(s)
		prefix := m[1]
		chain := m[2]
		operand := m[3]

		var bin string
		if strings.Count(chain, "-")%2 == 0 {
			bin = "+"
		} else {
			bin = "-"
		}

		return prefix + bin + operand
	})

	expr = reUnaryChain.ReplaceAllStringFunc(expr, func(s string) string {
		m := reUnaryChain.FindStringSubmatch(s)
		prefix := m[1]
		signs := m[2]
		num := m[3]

		neg := strings.Count(signs, "-")%2 == 1
		if neg {
			return prefix + "-" + num
		}
		return prefix + num
	})

	return expr
}

func extractBrackets(expr string, i, step int) int {
	balance := 0

	for j := i; j >= 0 && j < len(expr); j += step {
		switch expr[j] {
		case '(':
			balance++
		case ')':
			balance--
		}

		if balance == 0 {
			if step > 0 {
				return j + 1
			}
			return j
		}
	}
	return i
}

func extractOperandIndex(expr string, i int, step int) int {
	if step > 0 {
		if strings.HasPrefix(expr[i:], "pow(") {
			return extractBrackets(expr, i+3, step)
		}
		if expr[i] == '(' {
			return extractBrackets(expr, i, step)
		}

		m := reNumber.FindStringIndex(expr[i:])
		if m != nil {
			return i + m[1]
		}
	} else {
		if expr[i] == ')' {
			return extractBrackets(expr, i, step)
		}

		j := i
		for j >= 0 && (expr[j] == '.' || (expr[j] >= '0' && expr[j] <= '9')) {
			j--
		}
		return j + 1
	}

	return i
}

var reNumber = regexp.MustCompile(`-?\d+(\.\d+)?`)

func normalizePower(expr string) string {
	for {
		idx := strings.LastIndex(expr, "^")
		if idx == -1 {
			break
		}

		li := extractOperandIndex(expr, idx-1, -1)
		ri := extractOperandIndex(expr, idx+1, 1)

		expr = expr[:li] + "pow(" + expr[li:idx] + "," + expr[idx+1:ri] + ")" + expr[ri:]
	}
	return expr
}

func Normalize(expr string) string {
	expr = reSpace.ReplaceAllString(expr, "")
	expr = normalizeSigns(expr)
	expr = normalizePower(expr)
	return expr
}
