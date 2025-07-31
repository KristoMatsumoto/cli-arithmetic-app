package processor

import (
	"fmt"
	"math"
	"strconv"
)

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

func applyBinar(a, b float64, op string) (float64, error) {
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

func EvalExpression(expr string) (float64, error) {
	tokens, err := Tokenize(expr)
	if err != nil {
		return 0, err
	}

	var values []float64
	var ops []string

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
			val := values[len(values)-1]
			values[len(values)-1] = -val
			return nil
		}

		if len(values) < 2 {
			return fmt.Errorf("not enough values for binary operator")
		}

		right := values[len(values)-1]
		left := values[len(values)-2]
		values = values[:len(values)-2]

		res, err := applyBinar(left, right, op)
		if err != nil {
			return err
		}
		values = append(values, res)
		return nil
	}

	i := 0
	for i < len(tokens) {
		tok := tokens[i]

		switch tok.Type {
		case Number:
			num, _ := strconv.ParseFloat(tok.Value, 64)
			values = append(values, num)
			i++

		case Operator:
			// Unary minus
			if tok.Value == "-" && (i == 0 || tokens[i-1].Type == Operator || tokens[i-1].Type == LParen) {
				ops = append(ops, "UMINUS")
				i++
				continue
			}

			// Other operations
			for len(ops) > 0 {
				top := ops[len(ops)-1]
				if (isRightAssociative(tok.Value) && precedence(tok.Value) < precedence(top)) ||
					(!isRightAssociative(tok.Value) && precedence(tok.Value) <= precedence(top)) {
					if err := apply(); err != nil {
						return 0, err
					}
				} else {
					break
				}
			}
			ops = append(ops, tok.Value)
			i++

		case LParen:
			ops = append(ops, "(")
			i++

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
			i++
		}
	}

	for len(ops) > 0 {
		if err := apply(); err != nil {
			return 0, err
		}
	}

	if len(values) != 1 {
		return 0, fmt.Errorf("invalid result")
	}
	return values[0], nil
}
