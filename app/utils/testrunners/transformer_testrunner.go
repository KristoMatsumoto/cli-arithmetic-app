package testrunners

import (
	"bytes"
	"fmt"

	"cli-arithmetic-app/app/modules/transformer"

	"github.com/ozontech/allure-go/pkg/framework/provider"
)

// Runs a general set of tests for a specific transformer
func RunCommonTransformerTests(t provider.T, tr transformer.Transformer, data []byte) {
	t.WithNewStep(fmt.Sprintf("Encode/Decode for %s", tr.Name()), func(sCtx provider.StepCtx) {
		encoded, err := tr.Encode(data)
		sCtx.Assert().NoError(err, "Encode should complete without errors")
		sCtx.Assert().True(len(encoded) > 0, "Encoded data should not be empty")

		decoded, err := tr.Decode(encoded)
		sCtx.Assert().NoError(err, "Decode should complete without errors")
		sCtx.Assert().Equal(data, decoded, "The result after decoding must match the original data")
	})
}

// Checks the correct processing of invalid data
func RunInvalidDataTest(t provider.T, tr transformer.Transformer) {
	t.WithNewStep(fmt.Sprintf("Decode invalid data for %s", tr.Name()), func(sCtx provider.StepCtx) {
		invalid := []byte("not a valid compressed stream")
		_, err := tr.Decode(invalid)
		sCtx.Assert().Error(err, "Decode should return an error for invalid data")
	})
}

// Checks Encode/Decode of empty input
func RunEmptyInputTest(t provider.T, tr transformer.Transformer) {
	t.WithNewStep(fmt.Sprintf("Encode/Decode of empty input for %s", tr.Name()), func(sCtx provider.StepCtx) {
		input := []byte{}
		encoded, err := tr.Encode(input)
		sCtx.Assert().NoError(err)
		decoded, err := tr.Decode(encoded)
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(input, decoded, "Empty input must be encoded and decoded correctly")
	})
}

// Tests resistance to big data
func RunLargeDataTest(t provider.T, tr transformer.Transformer) {
	t.WithNewStep(fmt.Sprintf("Encode/Decode big data for %s", tr.Name()), func(sCtx provider.StepCtx) {
		input := bytes.Repeat([]byte("test-data"), 100000) // ~800KB
		encoded, err := tr.Encode(input)
		sCtx.Assert().NoError(err)
		sCtx.Assert().True(len(encoded) > 0)
		decoded, err := tr.Decode(encoded)
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(input, decoded)
	})
}
