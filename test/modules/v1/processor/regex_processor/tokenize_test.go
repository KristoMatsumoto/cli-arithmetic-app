package regex_processor_test

import (
	"cli-arithmetic-app/modules/v1/processor/regex_processor"
	"encoding/json"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

type TokenizeCase struct {
	Name     string                  `json:"name"`
	Input    string                  `json:"input"`
	Expected []regex_processor.Token `json:"expected"`
}

func TestRegexProcessor_Tokenize(t *testing.T) {
	data := loadCases(t, "../tokenize_cases.json")
	var cases []TokenizeCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	for _, c := range cases {
		runner.Run(t, c.Name, func(t provider.T) {
			result := regex_processor.Tokenize(c.Input)
			t.Assert().Equal(c.Expected, result)
		})
	}
}
