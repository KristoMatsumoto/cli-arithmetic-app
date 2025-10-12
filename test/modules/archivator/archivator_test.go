package archivator_test

import (
	"fmt"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"

	"cli-arithmetic-app/modules/archivator"
	"cli-arithmetic-app/modules/transformer"
	"cli-arithmetic-app/utils/cases"
	"cli-arithmetic-app/utils/testrunners"
)

var transformers = []transformer.Transformer{
	archivator.NewBrotliTransformer(),
	archivator.NewGZIPTransformer(),
	archivator.NewLZ4Transformer(),
	archivator.NewZIPTransformer(),
	archivator.NewZIPXTransformer(),
	archivator.NewTARTransformer(),
	archivator.NewZSTDTransformer(),
}

func TestTransformer_ArchivatorsAll(t *testing.T) {
	var data []byte

	runner.Run(t, "Load data", func(tt provider.T) {
		data = cases.LoadCases(t, "./sample.json")
		tt.Assert().True(data != nil)
	})

	for _, tr := range transformers {
		tr := tr
		runner.Run(t, fmt.Sprintf("Testing %s transformer (archivator)", tr.Name()), func(t provider.T) {
			testrunners.RunCommonTransformerTests(t, tr, data)
			testrunners.RunEmptyInputTest(t, tr)
			testrunners.RunInvalidDataTest(t, tr)
			testrunners.RunLargeDataTest(t, tr)
		})
	}
}
