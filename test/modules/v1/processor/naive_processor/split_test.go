package naive_processor_test

import (
	"cli-arithmetic-app/modules/v1/processor/naive_processor"
	"encoding/json"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

type SplitCase struct {
	Name     string                      `json:"name"`
	Input    string                      `json:"input"`
	Expected []naive_processor.TokenPart `json:"expected"`
}

func TestNaiveProcessor_SplitIntoTokens(t *testing.T) {
	data := loadCases(t, "../split_cases.json")
	var cases []SplitCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	for _, c := range cases {
		runner.Run(t, c.Name, func(t provider.T) {
			result := naive_processor.SplitIntoTokens(c.Input)
			t.Assert().Equal(c.Expected, result)
		})
	}
}
