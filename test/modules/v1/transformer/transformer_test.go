package transformer_test

import (
	"cli-arithmetic-app/modules/v1/archivator"
	"cli-arithmetic-app/modules/v1/encryptor"
	"cli-arithmetic-app/modules/v1/transformer"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestTransformer_ChainRoundtrip(t *testing.T) {
	runner.Run(t, "Transform chain roundtrip", func(t provider.T) {
		zipT := archivator.NewZIPTransformer()
		aesT, err := encryptor.NewAESTransformer()
		t.Require().NoError(err, "failed to create AES transformer")

		chain := []transformer.Transformer{zipT, aesT}

		original := []byte("secret data for composed transformers")

		var encoded []byte
		t.WithNewStep("Encode through transform chain", func(sCtx provider.StepCtx) {
			encoded = original
			for _, tr := range chain {
				var err error
				encoded, err = tr.Encode(encoded)
				t.Assert().NoError(err, "%s encode failed", tr.Name())
			}
		})

		var decoded []byte
		t.WithNewStep("Decode through transform chain in reverse", func(sCtx provider.StepCtx) {
			decoded = encoded
			for i := len(chain) - 1; i >= 0; i-- {
				var err error
				decoded, err = chain[i].Decode(decoded)
				t.Assert().NoError(err, "%s decode failed", chain[i].Name())
			}
		})

		t.WithNewStep("Compare original and decoded data", func(sCtx provider.StepCtx) {
			t.Assert().Equal(original, decoded, "composed transform roundtrip mismatch")
		})
	})
}
