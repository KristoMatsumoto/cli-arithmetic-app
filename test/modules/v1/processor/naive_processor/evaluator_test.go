package naive_processor_test

import (
	"testing"

	"cli-arithmetic-app/modules/v1/processor"

	"github.com/stretchr/testify/assert"
)

func TestEvalExpression(t *testing.T) {
	tests := []struct {
		expr     string
		expected float64
	}{
		{"2+2", 4},
		{"3*4+1", 13},
		{"(5+5)/2", 5},
		{"2^3", 8},
		{"10%3", 1},
		{"2 + 3 * (4 ^ 2)", 50},
	}

	for _, tt := range tests {
		result, err := processor.EvalExpression(tt.expr)
		assert.NoError(t, err, "expression: %s", tt.expr)
		assert.InDelta(t, tt.expected, result, 1e-6)
	}
}

func TestEvalExpression_DivZero(t *testing.T) {
	_, err := processor.EvalExpression("5 / 0")
	assert.Error(t, err)
}

func TestEvalExpression_Malformed(t *testing.T) {
	_, err := processor.EvalExpression("2 + (3")
	assert.Error(t, err)
}
