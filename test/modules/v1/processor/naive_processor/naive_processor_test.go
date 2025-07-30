package naive_processor_test

import (
	"cli-arithmetic-app/modules/v1/processor"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNaiveProcessor_Process(t *testing.T) {
	p := processor.NewNaiveProcessor()

	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name: "basic expressions embedded",
			input: []string{
				"Line with 3+4 result",
				"Multiply: 5*2 and divide 10/5",
				"No expressions here",
			},
			expected: []string{
				"Line with 7 result",
				"Multiply: 10 and divide 2",
				"No expressions here",
			},
		},
		{
			name: "division by zero",
			input: []string{
				"This is broken: 5/0",
			},
			expected: []string{
				"This is broken: NaN",
			},
		},
		{
			name: "multiple in one line",
			input: []string{
				"1+2 then 3*4 then 10-5",
			},
			expected: []string{
				"3 then 12 then 5",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := p.Process(test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		})
	}
}
