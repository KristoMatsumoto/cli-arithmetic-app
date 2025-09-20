package naive_processor_test

import (
	"cli-arithmetic-app/modules/processor"
	"cli-arithmetic-app/modules/processor/naive_processor"
	"cli-arithmetic-app/utils/cases"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

type EvalCase struct {
	Name     string `json:"name"`
	Input    string `json:"input"`
	Expected string `json:"expected"`
}

func TestNaiveProcessor_EvalExpression(t *testing.T) {
	data := cases.LoadCases(t, "../eval_expression_cases.json")
	var cases []EvalCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	for _, c := range cases {
		runner.Run(t, c.Name, func(t provider.T) {
			output, err := naive_processor.EvalExpression(c.Input)

			if c.Expected == "NaN" {
				t.Assert().Error(err, "expected an error for %q", c.Input)
				return
			}

			want, perr := strconv.ParseFloat(c.Expected, 64)
			if perr != nil {
				t.Fatalf("bad expected value %q in case %q: %v", c.Expected, c.Name, perr)
			}
			t.Require().NoError(err, "unexpected error for %q", c.Input)

			// сравниваем с допуском, чтобы не спотыкаться на double
			const eps = 1e-9
			if diff := output - want; diff < -eps || diff > eps {
				t.Errorf("EvalExpression(%q) = %v, want %v", c.Input, output, want)
			}

			t.WithNewStep("Compare input and output", func(sCtx provider.StepCtx) {
				t.Assert().Equal(c.Expected, processor.FormatFloat(output))
			})
		})
	}
}
