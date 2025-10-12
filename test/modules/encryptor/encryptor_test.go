package encryptor_test

import (
	"fmt"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"

	"cli-arithmetic-app/modules/encryptor"
	"cli-arithmetic-app/modules/transformer"
	"cli-arithmetic-app/utils/cases"
	"cli-arithmetic-app/utils/testrunners"
)

var encryptors = []func() (transformer.Transformer, error){
	func() (transformer.Transformer, error) { return encryptor.NewTripleDESTransformer() },
	func() (transformer.Transformer, error) { return encryptor.NewAESTransformer() },
	func() (transformer.Transformer, error) { return encryptor.NewAESCBCTransformer() },
	func() (transformer.Transformer, error) { return encryptor.NewAESGCMTransformer() },
	func() (transformer.Transformer, error) { return encryptor.NewBlowfishTransformer() },
	func() (transformer.Transformer, error) { return encryptor.NewChaCha20Transformer() },
	func() (transformer.Transformer, error) { return encryptor.NewChaCha20Poly1305Transformer() },
	func() (transformer.Transformer, error) { return encryptor.NewGOST28147Transformer() },
	func() (transformer.Transformer, error) { return encryptor.NewRC4Transformer() },
	func() (transformer.Transformer, error) { return encryptor.NewXORTransformer() },
}

func TestTransformer_EncryptorsAll(t *testing.T) {
	var data []byte

	runner.Run(t, "Load encryption test data", func(tt provider.T) {
		data = cases.LoadCases(t, "./sample.json")
		tt.Assert().True(len(data) > 0)
	})

	for _, en := range encryptors {
		e, err := en()

		runner.Run(t, fmt.Sprintf("Testing %s transformer (encryptor)", e.Name()), func(t provider.T) {
			t.Assert().NoError(err, "Cannot create transformer")
			testrunners.RunCommonTransformerTests(t, e, data)
			testrunners.RunEmptyInputTest(t, e)
			testrunners.RunInvalidDataTest(t, e)
			testrunners.RunLargeDataTest(t, e)
		})
	}
}
