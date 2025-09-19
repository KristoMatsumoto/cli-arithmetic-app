package roundtrip

import (
	"cli-arithmetic-app/modules/v1/parser"
	"os"
	"path/filepath"

	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func Roundtrip(t provider.T, p parser.Parser, inputPath string) {
	tempFile := filepath.Join(os.TempDir(), "roundtrip_output.tmp")

	var original []string
	t.WithNewStep("Read original file", func(sCtx provider.StepCtx) {
		var err error
		original, err = p.ReadFile(inputPath)
		t.Assert().NoError(err, "reading original file")
	})

	t.WithNewStep("Write content to temp file", func(sCtx provider.StepCtx) {
		err := p.WriteFile(tempFile, original)
		t.Assert().NoError(err, "writing roundtrip file")
	})

	var reconstructed []string
	t.WithNewStep("Read back from temp file", func(sCtx provider.StepCtx) {
		var err error
		reconstructed, err = p.ReadFile(tempFile)
		t.Assert().NoError(err, "reading roundtrip file")
	})

	t.WithNewStep("Compare original and reconstructed content", func(sCtx provider.StepCtx) {
		t.Assert().Equal(original, reconstructed, "roundtrip mismatch")
	})

	_ = os.Remove(tempFile)
}

func RoundtripBytes(t provider.T, p parser.Parser, original []string) {
	var serialized []byte
	t.WithNewStep("Serialize to bytes", func(sCtx provider.StepCtx) {
		var err error
		serialized, err = p.SerializeBytes(original)
		t.Assert().NoError(err, "serialization failed")
	})

	var parsed []string
	t.WithNewStep("Parse bytes back", func(sCtx provider.StepCtx) {
		var err error
		parsed, err = p.ParseBytes(serialized)
		t.Assert().NoError(err, "parse failed")
	})

	t.WithNewStep("Compare original and parsed data", func(sCtx provider.StepCtx) {
		t.Assert().Equal(original, parsed, "roundtrip mismatch")
	})
}
