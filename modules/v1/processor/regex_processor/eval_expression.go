package regex_processor

import (
	"fmt"
	"math"
	"strconv"
)

var precedence = map[string]int{
	"UMINUS": 3, "^": 3,
	"*": 2, "/": 2, "%": 2,
	"+": 1, "-": 1,
}

func applyOp(op string, stack *[]float64) error {
	if len(*stack) == 0 {
		return fmt.Errorf("empty stack for operator %s", op)
	}

	if op == "UMINUS" {
		val := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
		*stack = append(*stack, -val)
		return nil
	}

	if len(*stack) < 2 {
		return fmt.Errorf("not enough operands for %s", op)
	}

	b := (*stack)[len(*stack)-1]
	a := (*stack)[len(*stack)-2]
	*stack = (*stack)[:len(*stack)-2]

	var res float64
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			return fmt.Errorf("division by zero")
		}
		res = a / b
	case "%":
		if b == 0 {
			return fmt.Errorf("modulo by zero")
		}
		res = float64(int64(a) % int64(b))
	case "^":
		res = pow(a, b)
	default:
		return fmt.Errorf("unknown operator %s", op)
	}
	*stack = append(*stack, res)
	return nil
}

func pow(a, b float64) float64 {
	res := 1.0
	for i := 0; i < int(b); i++ {
		res *= a
	}
	return res
}

func EvalExpression(expr string) (float64, error) {
	tokens := Tokenize(expr)
	if len(tokens) == 0 {
		return math.NaN(), fmt.Errorf("empty expression")
	}

	var output []Token
	var ops []Token

	for _, t := range tokens {
		switch t.Type {
		case number:
			output = append(output, t)
		case operator:
			for len(ops) > 0 {
				top := ops[len(ops)-1]
				if top.Type == operator && precedence[top.Value] >= precedence[t.Value] {
					output = append(output, top)
					ops = ops[:len(ops)-1]
				} else {
					break
				}
			}
			ops = append(ops, t)
		case lparen:
			ops = append(ops, t)
		case rparen:
			for len(ops) > 0 && ops[len(ops)-1].Type != lparen {
				output = append(output, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			if len(ops) == 0 {
				return math.NaN(), fmt.Errorf("unbalanced parentheses")
			}
			ops = ops[:len(ops)-1] // убираем '('
		}
	}
	for len(ops) > 0 {
		if ops[len(ops)-1].Type == lparen {
			return math.NaN(), fmt.Errorf("unbalanced parentheses")
		}
		output = append(output, ops[len(ops)-1])
		ops = ops[:len(ops)-1]
	}

	var stack []float64
	for _, t := range output {
		switch t.Type {
		case number:
			val, err := strconv.ParseFloat(t.Value, 64)
			if err != nil {
				return math.NaN(), err
			}
			stack = append(stack, val)
		case operator:
			if err := applyOp(t.Value, &stack); err != nil {
				return math.NaN(), err
			}
		}
	}

	if len(stack) != 1 {
		return math.NaN(), fmt.Errorf("invalid expression")
	}
	return stack[0], nil
}
