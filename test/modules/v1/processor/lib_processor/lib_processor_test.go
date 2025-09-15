package lib_processor_test

import (
	"cli-arithmetic-app/modules/v1/processor/lib_processor"
	"cli-arithmetic-app/utils/cases"
	"encoding/json"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

type TestCase struct {
	Name     string   `json:"name"`
	Input    []string `json:"input"`
	Expected []string `json:"expected"`
}

func TestLibProcessor_Process(t *testing.T) {
	data := cases.LoadCases(t, "./lib_processor_cases.json")
	var cases []TestCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	p := lib_processor.NewLibProcessor()

	for _, c := range cases {
		runner.Run(t, c.Name, func(t provider.T) {
			result, err := p.Process(c.Input)
			t.Require().NoError(err)
			t.Assert().Equal(c.Expected, result)
		})
	}
}
