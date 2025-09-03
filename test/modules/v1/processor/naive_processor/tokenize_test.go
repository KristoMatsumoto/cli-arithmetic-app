package naive_processor_test

import (
	"cli-arithmetic-app/modules/v1/processor/naive_processor"
	"cli-arithmetic-app/utils/cases"
	"encoding/json"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

type TokenizeCase struct {
	Name     string                  `json:"name"`
	Input    string                  `json:"input"`
	Expected []naive_processor.Token `json:"expected"`
}

func TestNaiveProcessor_Tokenize(t *testing.T) {
	data := cases.LoadCases(t, "../tokenize_cases.json")
	var cases []TokenizeCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	for _, c := range cases {
		runner.Run(t, c.Name, func(t provider.T) {
			result, err := naive_processor.Tokenize(c.Input)
			t.Require().NoError(err)
			t.Assert().Equal(c.Expected, result)
		})
	}
}
