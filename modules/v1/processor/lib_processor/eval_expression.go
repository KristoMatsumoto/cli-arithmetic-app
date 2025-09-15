package lib_processor

import (
	"math"

	"github.com/Knetic/govaluate"
)

func EvalExpression(expr string) (float64, error) {
	expr = Normalize(expr)

	e, err := govaluate.NewEvaluableExpression(expr)
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
