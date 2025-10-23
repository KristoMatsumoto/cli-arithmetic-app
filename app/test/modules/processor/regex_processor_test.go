package processor_test

import (
	"cli-arithmetic-app/app/modules/processor"
	"cli-arithmetic-app/app/modules/processor/regex_processor"
	"cli-arithmetic-app/app/utils/cases"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type RegexSplitCase struct {
	Name     string                      `json:"name"`
	Input    string                      `json:"input"`
	Expected []regex_processor.TokenPart `json:"expected"`
}
type RegexTokenizeCase struct {
	Name     string                  `json:"name"`
	Input    string                  `json:"input"`
	Expected []regex_processor.Token `json:"expected"`
}

type RegexProcessorSuite struct {
	suite.Suite
}

func (s *RegexProcessorSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Processors")
	t.Story("Regex processor")
	t.WithParameters(allure.NewParameter("processor", "regex"))
	t.Tags("app", "processor", "regex processor", "regex", "processor")
	// t.Owner("github.com/KristoMatsumoto")
}

func (s *RegexProcessorSuite) TestProcess(t provider.T) {
	t.Title("Process")
	t.Description("")

	p := regex_processor.NewRegexProcessor()
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

func (s *RegexProcessorSuite) TestEvalExpression(t provider.T) {
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
			output, err := regex_processor.EvalExpression(c.Input)

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

func (s *RegexProcessorSuite) TestSplitIntoTokens(t provider.T) {
	t.Title("Split into tokens (parsing)")
	t.Description("")

	var testCases []RegexSplitCase
	t.WithNewStep("Load test data", func(sCtx provider.StepCtx) {
		data, err := cases.LoadCases("./split_cases.json")
		sCtx.Require().NoError(err, "failed to load data")
		errUnm := json.Unmarshal(data, &testCases)
		sCtx.Require().NoError(errUnm, "failed to unmarshal data")
	})

	for _, c := range testCases {
		t.WithNewStep(c.Name, func(sCtx provider.StepCtx) {
			result := regex_processor.SplitIntoTokens(c.Input)
			sCtx.Assert().Equal(c.Expected, result)
		})
	}
}

func (s *RegexProcessorSuite) TestTokenize(t provider.T) {
	t.Title("Tokenize")
	t.Description("")

	var testCases []RegexTokenizeCase
	t.WithNewStep("Load test data", func(sCtx provider.StepCtx) {
		data, err := cases.LoadCases("./tokenize_cases.json")
		sCtx.Require().NoError(err, "failed to load data")
		errUnm := json.Unmarshal(data, &testCases)
		sCtx.Require().NoError(errUnm, "failed to unmarshal data")
	})

	for _, c := range testCases {
		t.WithNewStep(c.Name, func(sCtx provider.StepCtx) {
			result := regex_processor.Tokenize(c.Input)
			sCtx.Assert().Equal(c.Expected, result)
		})
	}
}

func TestRegexProcessorSuite(t *testing.T) {
	suite.RunSuite(t, new(RegexProcessorSuite))
}
