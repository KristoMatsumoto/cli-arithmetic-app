package lib_processor_test

import (
	"cli-arithmetic-app/modules/v1/processor/lib_processor"
	"cli-arithmetic-app/utils/cases"
	"encoding/json"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

type NormalizeCase struct {
	Name     string `json:"name"`
	Input    string `json:"input"`
	Expected string `json:"expected"`
}

func TestLibProcessor_Normalize(t *testing.T) {
	data := cases.LoadCases(t, "./normalize_cases.json")
	var cases []NormalizeCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	for _, c := range cases {
		runner.Run(t, c.Name, func(t provider.T) {
			output := lib_processor.Normalize(c.Input)

			t.WithNewStep("Compare input and output", func(sCtx provider.StepCtx) {
				t.Assert().Equal(c.Expected, output)
			})
		})
	}
}
