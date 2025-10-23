package archivator_test

import (
	"fmt"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"cli-arithmetic-app/app/modules/archivator"
	"cli-arithmetic-app/app/modules/transformer"
	"cli-arithmetic-app/app/utils/cases"
	"cli-arithmetic-app/app/utils/testrunners"
)

type ArchivatorSuite struct {
	suite.Suite
	data []byte
}

func (s *ArchivatorSuite) BeforeAll(t provider.T) {
	t.WithNewStep("Load encryption test data", func(sCtx provider.StepCtx) {
		var err error
		s.data, err = cases.LoadCases("./sample.json")
		sCtx.Assert().NoError(err, "failed to load data")
	})
}

func (s *ArchivatorSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Archivators")
	t.Tags("app", "archive", "arhivator")
	// t.Owner("github.com/KristoMatsumoto")
}

func (s *ArchivatorSuite) testArchivator(t provider.T, a transformer.Transformer) {
	t.Title("Archivator")
	t.Description(fmt.Sprintf("Full encode/decode validation for %s transformer", a.Name()))
	t.WithParameters(allure.NewParameter("transformer", a.Name()))

	testrunners.RunCommonTransformerTests(t, a, s.data)
	testrunners.RunEmptyInputTest(t, a)
	testrunners.RunInvalidDataTest(t, a)
	testrunners.RunLargeDataTest(t, a)
}

func (s *ArchivatorSuite) TestBrotli(t provider.T) {
	s.testArchivator(t, archivator.NewBrotliTransformer())
}

func (s *ArchivatorSuite) TestGZIP(t provider.T) {
	s.testArchivator(t, archivator.NewGZIPTransformer())
}

func (s *ArchivatorSuite) TestLZ4(t provider.T) {
	s.testArchivator(t, archivator.NewLZ4Transformer())
}

func (s *ArchivatorSuite) TestTAR(t provider.T) {
	s.testArchivator(t, archivator.NewTARTransformer())
}

func (s *ArchivatorSuite) TestZIP(t provider.T) {
	s.testArchivator(t, archivator.NewZIPTransformer())
}

func (s *ArchivatorSuite) TestZIPX(t provider.T) {
	s.testArchivator(t, archivator.NewZIPXTransformer())
}

func (s *ArchivatorSuite) TestZSTD(t provider.T) {
	s.testArchivator(t, archivator.NewZSTDTransformer())
}

func TestArchivatorSuite(t *testing.T) {
	suite.RunSuite(t, new(ArchivatorSuite))
}
