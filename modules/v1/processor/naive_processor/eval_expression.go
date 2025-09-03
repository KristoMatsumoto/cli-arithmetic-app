package naive_processor

import (
	"fmt"
	"math"
	"strconv"
)

func EvalSimple(a, b float64, op string) (float64, error) {
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

func EvalExpression(expr string) (float64, error) {
	tokens, err := Tokenize(expr)
	if err != nil {
		return 0, err
	}

	var values []float64
	var ops []string

	// applies the upper operator from the stack to the values
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
			values[len(values)-1] = -values[len(values)-1]
			return nil
		}

		if len(values) < 2 {
			return fmt.Errorf("not enough values for binary operator %q", op)
		}
		b := values[len(values)-1]
		a := values[len(values)-2]
		values = values[:len(values)-2]

		res, err := EvalSimple(a, b, op)
		if err != nil {
			return err
		}
		values = append(values, res)
		return nil
	}

	for _, tok := range tokens {
		switch tok.Type {
		case Number:
			num, err := strconv.ParseFloat(tok.Value, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number %q: %w", tok.Value, err)
			}
			values = append(values, num)

		case Operator:
			opVal := tok.Value

			for len(ops) > 0 {
				top := ops[len(ops)-1]
				if top == "(" {
					break
				}
				if (isRightAssociative(opVal) && precedence(opVal) < precedence(top)) ||
					(!isRightAssociative(opVal) && precedence(opVal) <= precedence(top)) {
					if err := apply(); err != nil {
						return 0, err
					}
				} else {
					break
				}
			}
			ops = append(ops, opVal)

		case LParen:
			ops = append(ops, "(")

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
		}
	}

	for len(ops) > 0 {
		if ops[len(ops)-1] == "(" {
			return 0, fmt.Errorf("mismatched parenthesis")
		}
		if err := apply(); err != nil {
			return 0, err
		}
	}

	if len(values) != 1 {
		return 0, fmt.Errorf("invalid result")
	}
	return values[0], nil
}
