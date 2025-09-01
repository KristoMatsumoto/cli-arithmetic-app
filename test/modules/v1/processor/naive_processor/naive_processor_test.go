package naive_processor_test

import (
	"cli-arithmetic-app/modules/v1/processor/naive_processor"
	"encoding/json"
	"os"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

type TestCase struct {
	Name     string   `json:"name"`
	Input    []string `json:"input"`
	Expected []string `json:"expected"`
}

func loadCases(t *testing.T, path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read %s: %v", path, err)
	}

	return data
}

func TestNaiveProcessor_Process(t *testing.T) {
	data := loadCases(t, "../processor_cases.json")
	var cases []TestCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	p := naive_processor.NewNaiveProcessor()

	for _, c := range cases {
		runner.Run(t, c.Name, func(t provider.T) {
			result, err := p.Process(c.Input)
			t.Require().NoError(err)
			t.Assert().Equal(c.Expected, result)
		})
	}
}
