package encryptor_test

import (
	"cli-arithmetic-app/modules/v1/encryptor"
	"cli-arithmetic-app/utils/roundtrip"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestAesTransformer_Roundtrip(t *testing.T) {
	runner.Run(t, "AES transformer (encryptor) Roundtrip", func(t provider.T) {
		txf, err := encryptor.NewAESTransformer()
		t.Require().NoError(err, "failed to create AES transformer")

		roundtrip.RoundtripTransformer(t, txf, []byte("secret message for aes transformer"))
	})
}
