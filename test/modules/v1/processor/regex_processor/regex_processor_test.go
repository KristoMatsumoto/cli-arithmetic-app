package regex_processor_test

import (
	"cli-arithmetic-app/modules/v1/processor"
	"testing"

	// 	"encoding/json"
	// 	"os"
	// 	"path/filepath"
	"github.com/stretchr/testify/assert"
)

func TestRegexProcessor_Process(t *testing.T) {
	p := processor.NewRegexProcessor()

	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name: "basic expressions embedded",
			input: []string{
				"Line with 3+4 result",
				"Multiply: 5 * 2 and divide 10 / 5",
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
				"This is broken: 5 / 0",
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

// type TestCase struct {
// 	Name     string   `json:"name"`
// 	Input    []string `json:"input"`
// 	Expected []string `json:"expected"`
// }

// func loadTestCases(path string) ([]TestCase, error) {
// 	var cases []TestCase
// 	data, err := os.ReadFile(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = json.Unmarshal(data, &cases)
// 	return cases, err
// }

// func TestRegexProcessor_FromJSON(t *testing.T) {
// 	p := processor.NewRegexProcessor()

// 	wd, _ := os.Getwd()
// 	testPath := filepath.Join(wd, "..", "processor_cases.json")

// 	cases, err := loadTestCases(testPath)
// 	assert.NoError(t, err)

// 	for _, test := range cases {
// 		t.Run(test.Name, func(t *testing.T) {
// 			result, err := p.Process(test.Input)
// 			assert.NoError(t, err)
// 			assert.Equal(t, test.Expected, result)
// 		})
// 	}
// }
