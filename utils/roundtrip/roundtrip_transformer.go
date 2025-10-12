package roundtrip

import (
	"cli-arithmetic-app/modules/transformer"

	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func RunRoundtripTransformer(t provider.T, tr transformer.Transformer, data []byte) {
	var encoded []byte
	t.WithNewStep("Encode data", func(sCtx provider.StepCtx) {
		var err error
		encoded, err = tr.Encode(data)
		t.Assert().NoError(err, "%s encode failed", tr.Name())
	})

	var decoded []byte
	t.WithNewStep("Decode data", func(sCtx provider.StepCtx) {
		var err error
		decoded, err = tr.Decode(encoded)
		t.Assert().NoError(err, "%s decode failed", tr.Name())
	})

	t.WithNewStep("Compare original and decoded data", func(sCtx provider.StepCtx) {
		t.Assert().Equal(data, decoded, "%s roundtrip mismatch", tr.Name())
	})
}
