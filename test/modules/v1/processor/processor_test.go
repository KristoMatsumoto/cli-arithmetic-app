package processor_test

import (
	"cli-arithmetic-app/modules/v1/processor"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name     string   `json:"name"`
	Input    []string `json:"input"`
	Expected []string `json:"expected"`
}

func loadTestCases(t *testing.T) []TestCase {
	t.Helper()

	path := filepath.Join("processor_cases.json")

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read test cases: %v", err)
	}

	var cases []TestCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("failed to unmarshal test cases: %v", err)
	}

	return cases
}

func TestNaiveProcessor(t *testing.T) {
	p := processor.NewNaiveProcessor()

	cases := loadTestCases(t)

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			result, err := p.Process(test.Input)
			assert.NoError(t, err)
			assert.Equal(t, test.Expected, result)
		})
	}
}

func TestRegexProcessor(t *testing.T) {
	p := processor.NewRegexProcessor()

	cases := loadTestCases(t)

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			result, err := p.Process(test.Input)
			assert.NoError(t, err)
			assert.Equal(t, test.Expected, result)
		})
	}
}
