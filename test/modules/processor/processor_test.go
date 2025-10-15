package processor_test

import (
	"cli-arithmetic-app/modules/processor"
	"cli-arithmetic-app/utils/cases"
	"encoding/json"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type ProcessCase struct {
	Name     string   `json:"name"`
	Input    []string `json:"input"`
	Expected []string `json:"expected"`
}
type EvalCase struct {
	Name     string `json:"name"`
	Input    string `json:"input"`
	Expected string `json:"expected"`
}
type NormalizeCase struct {
	Name     string `json:"name"`
	Input    string `json:"input"`
	Expected string `json:"expected"`
}
type ParseCase struct {
	Name     string              `json:"name"`
	Input    string              `json:"input"`
	Expected []map[string]string `json:"expected"`
}
type FormatCase struct {
	Name     string  `json:"name"`
	Input    float64 `json:"input"`
	Expected string  `json:"expected"`
}

type ProcessorSuite struct {
	suite.Suite
}

func (s *ProcessorSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Processors")
	t.Tags("app", "processor")
	// t.Owner("github.com/KristoMatsumoto")
}

func (s *ProcessorSuite) TestFormatFloat(t provider.T) {
	t.Title("Format")
	t.Description("Test how method FormatFloat() format data")

	var testCases []FormatCase
	t.WithNewStep("Load test data", func(sCtx provider.StepCtx) {
		data, err := cases.LoadCases("./formatfloat_cases.json")
		sCtx.Require().NoError(err, "failed to load data")
		errUnm := json.Unmarshal(data, &testCases)
		sCtx.Require().NoError(errUnm, "failed to unmarshal data")
	})

	for _, c := range testCases {
		t.WithNewStep(c.Name, func(sCtx provider.StepCtx) {
			result := processor.FormatFloat(c.Input)
			sCtx.Assert().Equal(c.Expected, result)
		})
	}
}

func TestProcessorSuite(t *testing.T) {
	suite.RunSuite(t, new(ProcessorSuite))
}
