package processor_test

import (
	"cli-arithmetic-app/app/modules/processor"
	"cli-arithmetic-app/app/modules/processor/lib_processor"
	"cli-arithmetic-app/app/utils/cases"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type LibProcessorSuite struct {
	suite.Suite
}

func (s *LibProcessorSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Processors")
	t.Story("Library processor")
	t.WithParameters(allure.NewParameter("processor", "lib"))
	t.Tags("app", "processor", "lib processor", "library", "lib", "processor")
	// t.Owner("github.com/KristoMatsumoto")
}

func (s *LibProcessorSuite) TestProcess(t provider.T) {
	t.Title("Process")
	t.Description("")

	p := lib_processor.NewLibProcessor()
	var testCases []ProcessCase
	t.WithNewStep("Load test data", func(sCtx provider.StepCtx) {
		data, err := cases.LoadCases("./lib_processor/lib_processor_cases.json")
		sCtx.Require().NoError(err, "failed to load data")
		errUnm := json.Unmarshal(data, &testCases)
		sCtx.Require().NoError(errUnm, "failed to unmarshal data")
	})

	for _, c := range testCases {
		t.WithNewStep(c.Name, func(sCtx provider.StepCtx) {
			result, err := p.Process(c.Input)
			sCtx.Require().NoError(err)
			sCtx.Assert().Equal(c.Expected, result)
		})
	}
}

func (s *LibProcessorSuite) TestEvalExpression(t provider.T) {
	t.Title("Eval expression")
	t.Description("")

	var testCases []EvalCase
	t.WithNewStep("Load test data", func(sCtx provider.StepCtx) {
		data, err := cases.LoadCases("./lib_processor/eval_expression_cases.json")
		sCtx.Require().NoError(err, "failed to load data")
		errUnm := json.Unmarshal(data, &testCases)
		sCtx.Require().NoError(errUnm, "failed to unmarshal data")
	})

	for _, c := range testCases {
		t.WithNewStep(c.Name, func(step provider.StepCtx) {
			output, err := lib_processor.EvalExpression(c.Input)

			if c.Expected == "ERROR" {
				step.WithNewStep("Check on error", func(sCtx provider.StepCtx) {
					sCtx.Assert().Error(err, "expected an error for %q", c.Input)
				})
			} else {
				step.WithNewStep("Check on no error", func(sCtx provider.StepCtx) {
					want, perr := strconv.ParseFloat(c.Expected, 64)
					if perr != nil {
						t.Fatalf("bad expected value %q in case %q: %v", c.Expected, c.Name, perr)
					}
					sCtx.Require().NoError(err, "unexpected error for %q", c.Input)

					// сравниваем с допуском, чтобы не спотыкаться на double
					const eps = 1e-9
					if diff := output - want; diff < -eps || diff > eps {
						t.Errorf("EvalExpression(%q) = %v, want %v", c.Input, output, want)
					}
				})

				step.WithNewStep("Compare input and output", func(sCtx provider.StepCtx) {
					sCtx.Assert().Equal(c.Expected, processor.FormatFloat(output))
				})
			}
		})
	}
}

func (s *LibProcessorSuite) TestNormalize(t provider.T) {
	t.Title("Normalize")
	t.Description("")

	var testCases []NormalizeCase
	t.WithNewStep("Load test data", func(sCtx provider.StepCtx) {
		data, err := cases.LoadCases("./lib_processor/normalize_cases.json")
		sCtx.Require().NoError(err, "failed to load data")
		errUnm := json.Unmarshal(data, &testCases)
		sCtx.Require().NoError(errUnm, "failed to unmarshal data")
	})

	for _, c := range testCases {
		t.WithNewStep(c.Name, func(sCtx provider.StepCtx) {
			output := lib_processor.Normalize(c.Input)
			sCtx.Assert().Equal(c.Expected, output)
		})
	}
}

func (s *LibProcessorSuite) TestSplitIntoTokens(t provider.T) {
	t.Title("Split into tokens (parsing)")
	t.Description("")

	var testCases []ParseCase
	t.WithNewStep("Load test data", func(sCtx provider.StepCtx) {
		data, err := cases.LoadCases("./split_cases.json")
		sCtx.Require().NoError(err, "failed to load data")
		errUnm := json.Unmarshal(data, &testCases)
		sCtx.Require().NoError(errUnm, "failed to unmarshal data")
	})

	for _, c := range testCases {
		t.WithNewStep(c.Name, func(sCtx provider.StepCtx) {
			result, err := lib_processor.Parse("", []byte(c.Input))
			sCtx.Require().NoError(err)

			var rebuilt []map[string]string

			parts := result.([]interface{})
			for _, part := range parts {
				rebuilt = append(rebuilt, part.(map[string]string))
			}
			sCtx.Assert().Equal(c.Expected, rebuilt)
		})
	}
}

func TestLibProcessorSuite(t *testing.T) {
	suite.RunSuite(t, new(LibProcessorSuite))
}
