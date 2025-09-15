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

func Normalize(expr string) string {
	expr = reSpace.ReplaceAllString(expr, "")

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
