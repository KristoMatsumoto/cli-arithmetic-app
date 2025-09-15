package lib_processor

import (
	"math"

	"github.com/Knetic/govaluate"
)

var functions = map[string]govaluate.ExpressionFunction{
	"pow": func(args ...interface{}) (interface{}, error) {
		base := args[0].(float64)
		exp := args[1].(float64)
		return math.Pow(base, exp), nil
	},
}

func EvalExpression(expr string) (float64, error) {
	expr = Normalize(expr)

	e, err := govaluate.NewEvaluableExpressionWithFunctions(expr, functions)
	if err != nil {
		return math.NaN(), err
	}

	result, err := e.Evaluate(nil)
	if err != nil {
		return math.NaN(), err
	}

	val := result.(float64)

	if math.IsInf(val, 0) || math.IsNaN(val) {
		return math.NaN(), nil
	}

	return val, nil
}
