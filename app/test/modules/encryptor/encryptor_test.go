package encryptor_test

import (
	"cli-arithmetic-app/modules/encryptor"
	"cli-arithmetic-app/modules/transformer"
	"cli-arithmetic-app/utils/cases"
	"cli-arithmetic-app/utils/testrunners"
	"fmt"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type EncryptorSuite struct {
	suite.Suite
	data []byte
}

// var encryptors = []func() (transformer.Transformer, error){
// 	func() (transformer.Transformer, error) { return encryptor.NewTripleDESTransformer() },
// 	func() (transformer.Transformer, error) { return encryptor.NewAESTransformer() },
// 	func() (transformer.Transformer, error) { return encryptor.NewAESCBCTransformer() },
// 	func() (transformer.Transformer, error) { return encryptor.NewAESGCMTransformer() },
// 	func() (transformer.Transformer, error) { return encryptor.NewBlowfishTransformer() },
// 	func() (transformer.Transformer, error) { return encryptor.NewChaCha20Transformer() },
// 	func() (transformer.Transformer, error) { return encryptor.NewChaCha20Poly1305Transformer() },
// 	func() (transformer.Transformer, error) { return encryptor.NewGOST28147Transformer() },
// 	func() (transformer.Transformer, error) { return encryptor.NewRC4Transformer() },
// 	func() (transformer.Transformer, error) { return encryptor.NewXORTransformer() },
// }

func (s *EncryptorSuite) BeforeAll(t provider.T) {
	t.WithNewStep("Load encryption test data", func(sCtx provider.StepCtx) {
		var err error
		s.data, err = cases.LoadCases("./sample.json")
		sCtx.Assert().NoError(err, "failed to load data")
	})
}

func (s *EncryptorSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Encryptors")
	t.Tags("app", "encrypt", "decrypt", "crypt", "crypto")
	// t.Owner("github.com/KristoMatsumoto")
	// t.Link(allure.LinkLink("pkg crypto", "https://pkg.go.dev/crypto"))
}

func (s *EncryptorSuite) testEncryptor(t provider.T, factory func() (transformer.Transformer, error)) {
	e, err := factory()

	t.Title("Encryptor")
	t.Description(fmt.Sprintf("Full encode/decode validation for %s transformer", e.Name()))
	t.WithParameters(allure.NewParameter("transformer", e.Name()))

	t.WithNewStep("Initialization", func(sCtx provider.StepCtx) {
		sCtx.Assert().NoError(err, "Failed to initialize encryptor")
	})

	testrunners.RunCommonTransformerTests(t, e, s.data)
	testrunners.RunEmptyInputTest(t, e)
	testrunners.RunInvalidDataTest(t, e)
	testrunners.RunLargeDataTest(t, e)
}

func (s *EncryptorSuite) TestTripleDES(t provider.T) {
	s.testEncryptor(t, func() (transformer.Transformer, error) { return encryptor.NewTripleDESTransformer() })
}

func (s *EncryptorSuite) TestAES(t provider.T) {
	s.testEncryptor(t, func() (transformer.Transformer, error) { return encryptor.NewAESTransformer() })
}

func (s *EncryptorSuite) TestAESCBC(t provider.T) {
	s.testEncryptor(t, func() (transformer.Transformer, error) { return encryptor.NewAESCBCTransformer() })
}

func (s *EncryptorSuite) TestAESGCM(t provider.T) {
	s.testEncryptor(t, func() (transformer.Transformer, error) { return encryptor.NewAESGCMTransformer() })
}

func (s *EncryptorSuite) TestBlowfish(t provider.T) {
	s.testEncryptor(t, func() (transformer.Transformer, error) { return encryptor.NewBlowfishTransformer() })
}

func (s *EncryptorSuite) TestChaCha20(t provider.T) {
	s.testEncryptor(t, func() (transformer.Transformer, error) { return encryptor.NewChaCha20Transformer() })
}

func (s *EncryptorSuite) TestChaCha20Poly1305(t provider.T) {
	s.testEncryptor(t, func() (transformer.Transformer, error) { return encryptor.NewChaCha20Poly1305Transformer() })
}

func (s *EncryptorSuite) TestGOST28147(t provider.T) {
	s.testEncryptor(t, func() (transformer.Transformer, error) { return encryptor.NewGOST28147Transformer() })
}

func (s *EncryptorSuite) TestRC4(t provider.T) {
	s.testEncryptor(t, func() (transformer.Transformer, error) { return encryptor.NewRC4Transformer() })
}

func (s *EncryptorSuite) TestXOR(t provider.T) {
	s.testEncryptor(t, func() (transformer.Transformer, error) { return encryptor.NewXORTransformer() })
}

func TestEncryptorSuite(t *testing.T) {
	suite.RunSuite(t, new(EncryptorSuite))
}
