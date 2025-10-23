package parser_test

import (
	"cli-arithmetic-app/app/modules/parser"
	"cli-arithmetic-app/app/utils/testrunners"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type ParserSuite struct {
	suite.Suite
}

func (s *ParserSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Parsers")
	t.Tags("app", "parser")
	// t.Owner("github.com/KristoMatsumoto")
}

func (*ParserSuite) testReadWrite(t provider.T, p parser.Parser, input []string) {
	t.Title("Read & Write")
	t.Description("Cheking how parser read and write data")
	t.WithParameters(allure.NewParameter("parser", p.Format()))

	tempFile := filepath.Join(os.TempDir(), fmt.Sprintf("%s_test.%s", p.Format(), p.Format()))

	t.WithNewStep(fmt.Sprintf("Write data to temp %s file", p.Format()), func(sCtx provider.StepCtx) {
		err := p.WriteFile(tempFile, input)
		sCtx.Assert().NoError(err)
	})

	var output []string
	t.WithNewStep(fmt.Sprintf("Read data back from temp %s file", p.Format()), func(sCtx provider.StepCtx) {
		var err error
		output, err = p.ReadFile(tempFile)
		sCtx.Assert().NoError(err)
	})

	t.WithNewStep("Compare written and read data", func(sCtx provider.StepCtx) {
		sCtx.Assert().Equal(input, output)
	})

	_ = os.Remove(tempFile)
}

func (s *ParserSuite) TestReadWriteHTML(t provider.T) {
	s.testReadWrite(t, parser.NewHTMLParser(), []string{"3+3", "<b>bold</b>", "simple"})
}
func (s *ParserSuite) TestReadWriteJSON(t provider.T) {
	s.testReadWrite(t, parser.NewJSONParser(), []string{"4+4", "5-2", "Test"})
}
func (s *ParserSuite) TestReadWriteText(t provider.T) {
	s.testReadWrite(t, parser.NewTextParser(), []string{"1+1", "2*3", "Hello"})
}
func (s *ParserSuite) TestReadWriteXML(t provider.T) {
	s.testReadWrite(t, parser.NewXMLParser(), []string{"1+2", "8/4", "XMLLine"})
}
func (s *ParserSuite) TestReadWriteYAML(t provider.T) {
	s.testReadWrite(t, parser.NewYAMLParser(), []string{"10/2", "6*7", "World"})
}

func (s *ParserSuite) testRoundtrip(t provider.T, p parser.Parser) {
	t.Title("Roundtrip")
	t.Description("Roundtrip file parsing")
	t.WithParameters(allure.NewParameter("parser", p.Format()))

	testrunners.RunRoundtripParser(t, p, fmt.Sprintf("./sample.%s", p.Format()))
}

func (s *ParserSuite) TestRoundtripHTML(t provider.T) {
	s.testRoundtrip(t, parser.NewHTMLParser())
}
func (s *ParserSuite) TestRoundtripJSON(t provider.T) {
	s.testRoundtrip(t, parser.NewJSONParser())
}
func (s *ParserSuite) TestRoundtripText(t provider.T) {
	s.testRoundtrip(t, parser.NewTextParser())
}
func (s *ParserSuite) TestRoundtripXML(t provider.T) {
	s.testRoundtrip(t, parser.NewXMLParser())
}
func (s *ParserSuite) TestRoundtripYAML(t provider.T) {
	s.testRoundtrip(t, parser.NewHTMLParser())
}

func (s *ParserSuite) testRoundtripBytes(t provider.T, p parser.Parser) {
	t.Title("Roundtrip bytes")
	t.Description("Roundtrip bytes parsing")
	t.WithParameters(allure.NewParameter("parser", p.Format()))

	testrunners.RunRoundtripParserBytes(t, p, []string{fmt.Sprintf("hello %s parser", p.Format())})
}

func (s *ParserSuite) TestRoundtripBytesHTML(t provider.T) {
	s.testRoundtripBytes(t, parser.NewHTMLParser())
}
func (s *ParserSuite) TestRoundtripBytesJSON(t provider.T) {
	s.testRoundtripBytes(t, parser.NewJSONParser())
}
func (s *ParserSuite) TestRoundtripBytesText(t provider.T) {
	s.testRoundtripBytes(t, parser.NewTextParser())
}
func (s *ParserSuite) TestRoundtripBytesXML(t provider.T) {
	s.testRoundtripBytes(t, parser.NewXMLParser())
}
func (s *ParserSuite) TestRoundtripBytesYAML(t provider.T) {
	s.testRoundtripBytes(t, parser.NewHTMLParser())
}

func TestParserSuite(t *testing.T) {
	suite.RunSuite(t, new(ParserSuite))
}
