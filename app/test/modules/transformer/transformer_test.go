package transformer_test

import (
	"cli-arithmetic-app/app/modules/archivator"
	"cli-arithmetic-app/app/modules/encryptor"
	"cli-arithmetic-app/app/modules/transformer"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type TransformerSuite struct {
	suite.Suite
}

func (s *TransformerSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Transformers")
	t.Tags("app", "transformer", "encryptor", "archivator")
	// t.Owner("github.com/KristoMatsumoto")
}

func (s *TransformerSuite) TestTransformer_ChainRoundtrip(t provider.T) {
	t.Title("Chain roundtrip")
	t.Description("Roundtrip with chain zip->aes->zip")

	zipT := archivator.NewZIPTransformer()
	var aesT transformer.Transformer
	t.WithNewStep("Inicialization", func(sCtx provider.StepCtx) {
		var err error
		aesT, err = encryptor.NewAESTransformer()
		sCtx.Require().NoError(err, "failed to create AES transformer")
	})

	chain := []transformer.Transformer{zipT, aesT}

	original := []byte("secret data for composed transformers")

	var encoded []byte
	t.WithNewStep("Encode through transform chain", func(sCtx provider.StepCtx) {
		encoded = original
		for _, tr := range chain {
			var err error
			encoded, err = tr.Encode(encoded)
			sCtx.Assert().NoError(err, "%s encode failed", tr.Name())
		}
	})

	var decoded []byte
	t.WithNewStep("Decode through transform chain in reverse", func(sCtx provider.StepCtx) {
		decoded = encoded
		for i := len(chain) - 1; i >= 0; i-- {
			var err error
			decoded, err = chain[i].Decode(decoded)
			sCtx.Assert().NoError(err, "%s decode failed", chain[i].Name())
		}
	})

	t.WithNewStep("Compare original and decoded data", func(sCtx provider.StepCtx) {
		sCtx.Assert().Equal(original, decoded, "composed transform roundtrip mismatch")
	})
}

func TestTransformerSuite(t *testing.T) {
	suite.RunSuite(t, new(TransformerSuite))
}
