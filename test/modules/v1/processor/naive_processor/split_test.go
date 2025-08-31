package naive_processor_test

import (
	"cli-arithmetic-app/modules/v1/processor"
	"encoding/json"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

type SplitCase struct {
	Name     string                `json:"name"`
	Input    string                `json:"input"`
	Expected []processor.TokenPart `json:"expected"`
}

func TestSplitIntoTextAndExpr(t *testing.T) {
	data := loadCases(t, "testdata/split_cases.json")
	var cases []SplitCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	for _, c := range cases {
		runner.Run(t, c.Name, func(t provider.T) {
			result := processor.SplitIntoTokens(c.Input)
			t.Assert().Equal(c.Expected, result)
		})
	}
}
