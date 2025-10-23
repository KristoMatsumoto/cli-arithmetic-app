package processor_test

import (
	"cli-arithmetic-app/app/modules/processor"
	"cli-arithmetic-app/app/modules/processor/naive_processor"
	"cli-arithmetic-app/app/utils/cases"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type NaiveSplitCase struct {
	Name     string                      `json:"name"`
	Input    string                      `json:"input"`
	Expected []naive_processor.TokenPart `json:"expected"`
}
type NaiveTokenizeCase struct {
	Name     string                  `json:"name"`
	Input    string                  `json:"input"`
	Expected []naive_processor.Token `json:"expected"`
}

type NaiveProcessorSuite struct {
	suite.Suite
}

func (s *NaiveProcessorSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Processors")
	t.Story("Naive processor")
	t.WithParameters(allure.NewParameter("processor", "naive"))
	t.Tags("app", "processor", "naive processor", "naive", "processor")
	// t.Owner("github.com/KristoMatsumoto")
}

func (s *NaiveProcessorSuite) TestProcess(t provider.T) {
	t.Title("Process")
	t.Description("")

	p := naive_processor.NewNaiveProcessor()
	var testCases []ProcessCase
	t.WithNewStep("Load test data", func(sCtx provider.StepCtx) {
		data, err := cases.LoadCases("./processor_cases.json")
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

func (s *NaiveProcessorSuite) TestEvalExpression(t provider.T) {
	t.Title("Eval expression")
	t.Description("")

	var testCases []EvalCase
	t.WithNewStep("Load test data", func(sCtx provider.StepCtx) {
		data, err := cases.LoadCases("./eval_expression_cases.json")
		sCtx.Require().NoError(err, "failed to load data")
		errUnm := json.Unmarshal(data, &testCases)
		sCtx.Require().NoError(errUnm, "failed to unmarshal data")
	})

	for _, c := range testCases {
		t.WithNewStep(c.Name, func(step provider.StepCtx) {
			output, err := naive_processor.EvalExpression(c.Input)

			if c.Expected == "NaN" {
				step.WithNewStep("Check on NaN", func(sCtx provider.StepCtx) {
					sCtx.Assert().Error(err, "expected an error for %q", c.Input)
				})
			} else {
				step.WithNewStep("Check on no error", func(sCtx provider.StepCtx) {
					want, perr := strconv.ParseFloat(c.Expected, 64)
					if perr != nil {
						t.Fatalf("bad expected value %q in case %q: %v", c.Expected, c.Name, perr)
					}
					sCtx.Assert().NoError(err, "unexpected error for %q", c.Input)

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

func (s *NaiveProcessorSuite) TestSplitIntoTokens(t provider.T) {
	t.Title("Split into tokens (parsing)")
	t.Description("")

	var testCases []NaiveSplitCase
	t.WithNewStep("Load test data", func(sCtx provider.StepCtx) {
		data, err := cases.LoadCases("./naive_processor/split_cases.json")
		sCtx.Require().NoError(err, "failed to load data")
		errUnm := json.Unmarshal(data, &testCases)
		sCtx.Require().NoError(errUnm, "failed to unmarshal data")
	})

	for _, c := range testCases {
		t.WithNewStep(c.Name, func(sCtx provider.StepCtx) {
			result := naive_processor.SplitIntoTokens(c.Input)
			sCtx.Assert().Equal(c.Expected, result)
		})
	}
}

func (s *NaiveProcessorSuite) TestTokenize(t provider.T) {
	t.Title("Tokenize")
	t.Description("")

	var testCases []NaiveTokenizeCase
	t.WithNewStep("Load test data", func(sCtx provider.StepCtx) {
		data, err := cases.LoadCases("./tokenize_cases.json")
		sCtx.Require().NoError(err, "failed to load data")
		errUnm := json.Unmarshal(data, &testCases)
		sCtx.Require().NoError(errUnm, "failed to unmarshal data")
	})

	for _, c := range testCases {
		t.WithNewStep(c.Name, func(sCtx provider.StepCtx) {
			result, err := naive_processor.Tokenize(c.Input)
			sCtx.Require().NoError(err)
			sCtx.Assert().Equal(c.Expected, result)
		})
	}
}

func TestNaiveProcessorSuite(t *testing.T) {
	suite.RunSuite(t, new(NaiveProcessorSuite))
}
