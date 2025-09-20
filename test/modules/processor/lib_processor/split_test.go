package lib_processor_test

import (
	"cli-arithmetic-app/modules/processor/lib_processor"
	"cli-arithmetic-app/utils/cases"
	"encoding/json"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

type ParseCase struct {
	Name     string              `json:"name"`
	Input    string              `json:"input"`
	Expected []map[string]string `json:"expected"`
}

func TestLibProcessor_SplitIntoTokens(t *testing.T) {
	data := cases.LoadCases(t, "../split_cases.json")
	var cases []ParseCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	for _, c := range cases {
		runner.Run(t, c.Name, func(t provider.T) {
			result, err := lib_processor.Parse("", []byte(c.Input))
			t.Require().NoError(err)

			var rebuilt []map[string]string

			parts := result.([]interface{})
			for _, part := range parts {
				rebuilt = append(rebuilt, part.(map[string]string))
			}
			t.Assert().Equal(c.Expected, rebuilt)
		})
	}
}
