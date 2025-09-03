package processor_test

import (
	"cli-arithmetic-app/modules/v1/processor"
	"cli-arithmetic-app/utils/cases"
	"encoding/json"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/stretchr/testify/assert"
)

type FormatCase struct {
	Name     string  `json:"name"`
	Input    float64 `json:"input"`
	Expected string  `json:"expected"`
}

func TestNaiveProcessor_FormatFloat(t *testing.T) {
	data := cases.LoadCases(t, "formatfloat_cases.json")
	var cases []FormatCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	for _, c := range cases {
		runner.Run(t, c.Name, func(t provider.T) {
			result := processor.FormatFloat(c.Input)
			assert.Equal(t, c.Expected, result)
		})
	}
}
