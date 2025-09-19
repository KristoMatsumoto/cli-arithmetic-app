package archivator_test

import (
	"cli-arithmetic-app/modules/v1/archivator"
	"cli-arithmetic-app/utils/roundtrip"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestZipTransformer_Roundtrip(t *testing.T) {
	runner.Run(t, "ZIP transformer (archivator) Roundtrip", func(t provider.T) {
		txf := archivator.NewZIPTransformer()
		roundtrip.RoundtripTransformer(t, txf, []byte("hello zip transformer"))
	})
}
