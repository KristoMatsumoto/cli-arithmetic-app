package testrunners

import (
	"cli-arithmetic-app/app/modules/parser"
	"cli-arithmetic-app/app/modules/transformer"
	"os"

	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func RunRoundtripParser(t provider.T, p parser.Parser, inputPath string) {
	var tempFile *os.File
	t.WithNewStep("Create temp file", func(sCtx provider.StepCtx) {
		var err error
		tempFile, err = os.CreateTemp("", "roundtrip_*.tmp")
		t.Assert().NoError(err)
		defer os.Remove(tempFile.Name())
		defer tempFile.Close()
	})

	var original []string
	t.WithNewStep("Read original file", func(sCtx provider.StepCtx) {
		var err error
		original, err = p.ReadFile(inputPath)
		sCtx.Assert().NoError(err, "reading original file")
	})

	t.WithNewStep("Write content to temp file", func(sCtx provider.StepCtx) {
		err := p.WriteFile(tempFile.Name(), original)
		sCtx.Assert().NoError(err, "writing roundtrip file")
	})

	var reconstructed []string
	t.WithNewStep("Read back from temp file", func(sCtx provider.StepCtx) {
		var err error
		reconstructed, err = p.ReadFile(tempFile.Name())
		sCtx.Assert().NoError(err, "reading roundtrip file")
	})

	t.WithNewStep("Compare original and reconstructed content", func(sCtx provider.StepCtx) {
		sCtx.Assert().Equal(original, reconstructed, "roundtrip mismatch")
	})
}

func RunRoundtripParserBytes(t provider.T, p parser.Parser, original []string) {
	var serialized []byte
	t.WithNewStep("Serialize to bytes", func(sCtx provider.StepCtx) {
		var err error
		serialized, err = p.SerializeBytes(original)
		sCtx.Assert().NoError(err, "serialization failed")
	})

	var parsed []string
	t.WithNewStep("Parse bytes back", func(sCtx provider.StepCtx) {
		var err error
		parsed, err = p.ParseBytes(serialized)
		sCtx.Assert().NoError(err, "parse failed")
	})

	t.WithNewStep("Compare original and parsed data", func(sCtx provider.StepCtx) {
		sCtx.Assert().Equal(original, parsed, "roundtrip mismatch")
	})
}

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
