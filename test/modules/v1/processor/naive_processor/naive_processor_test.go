package naive_processor_test

import (
	"cli-arithmetic-app/modules/v1/processor"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name     string   `json:"name"`
	Input    []string `json:"input"`
	Expected []string `json:"expected"`
}

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
				"1 + 2 then 3*4 then 10-5",
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

// Should be more
func TestExtractExpressions(t *testing.T) {
	line := "Some text 2+2 and (3*3) and plain text"
	exprs := processor.ExtractExpressions(line)
	assert.ElementsMatch(t, []string{"2+2", "(3*3)"}, exprs)
}

func TestReplaceFirst(t *testing.T) {
	input := "calculate 2+2 and 2+2 again"
	result := processor.ReplaceFirst(input, "2+2", "4")
	assert.Equal(t, "calculate 4 and 2+2 again", result)
}

func TestFormatFloat(t *testing.T) {
	assert.Equal(t, "5", processor.FormatFloat(5.0))
	assert.Equal(t, "5.50", processor.FormatFloat(5.5))
}
