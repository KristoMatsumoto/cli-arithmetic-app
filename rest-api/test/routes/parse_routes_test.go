package routes_test

import (
	"cli-arithmetic-app/rest-api/routes"
	"cli-arithmetic-app/rest-api/utils"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type ParseRoutesSuite struct {
	suite.Suite
	router *gin.Engine
}

func (s *ParseRoutesSuite) BeforeEach(t provider.T) {
	t.Epic("REST API")
	t.Feature("Routes")
	t.Story("Parse routes")
	t.Tags("rest-api", "routes", "parse")

	gin.SetMode(gin.TestMode)
	s.router = gin.Default()
	routes.RegisterParseRoutes(s.router)
}

func (s *ParseRoutesSuite) TestParseRoutes_Health(t provider.T) {
	t.Title("/parse/health")
	t.Description("Checks parse routes health")

	utils.RunHealthStatusCheck(t, s.router, "/api/v1/parse/health")
}

func (s *ParseRoutesSuite) TestParseRoutes_Parse_OK(t provider.T) {
	t.Title("/parse")
	t.Description("Successful request processing test for /parse")

	body := map[string]interface{}{
		"format": "txt",
		"data":   []byte("1+2\n3*4"),
	}
	payload, _ := json.Marshal(body)

	utils.RunTestRouteOK(t, s.router, "/api/v1/parse", payload,
		func(resp *httptest.ResponseRecorder) string { return resp.Body.String() }, `"lines":["1+2","3*4"]`)
}

func (s *ParseRoutesSuite) TestParseRoutes_Compose_OK(t provider.T) {
	t.Title("/compose")
	t.Description("Successful request processing test for /compose")

	body := map[string]interface{}{
		"format": "txt",
		"lines":  []string{"1+2", "3*4"},
	}
	payload, _ := json.Marshal(body)

	utils.RunTestRouteOK(t, s.router, "/api/v1/compose", payload,
		func(resp *httptest.ResponseRecorder) string { return resp.Header().Get("Content-Type") }, "text/plain")
}

func (s *ParseRoutesSuite) TestParseRoutes_Parse_BadRequest(t provider.T) {
	t.Title("/parse")
	t.Description("Invalid request test for /parse — missing required field")

	body := map[string]interface{}{
		"data": []byte("1+2\n3*4"),
	}
	payload, _ := json.Marshal(body)

	utils.RunErrorCheckingTest(t, s.router, "/api/v1/parse", payload, http.StatusBadRequest)
}

func (s *ParseRoutesSuite) TestParseRoutes_Compose_BadRequest(t provider.T) {
	t.Title("/compose")
	t.Description("Invalid request test for /compose — missing required field")

	body := map[string]interface{}{
		"lines": []string{"1+2", "3*4"},
	}
	payload, _ := json.Marshal(body)

	utils.RunErrorCheckingTest(t, s.router, "/api/v1/compose", payload, http.StatusBadRequest)
}

func (s *ParseRoutesSuite) TestParseRoutes_Parse_InternalServerError(t provider.T) {
	t.Title("/parse")
	t.Description("Invalid request test for /parse — field with incorrect data")

	body := map[string]interface{}{
		"format": "t",
		"data":   []byte("1+2\n3*4"),
	}
	payload, _ := json.Marshal(body)

	utils.RunErrorCheckingTest(t, s.router, "/api/v1/parse", payload, http.StatusInternalServerError)
}

func (s *ParseRoutesSuite) TestParseRoutes_Compose_InternalServerError(t provider.T) {
	t.Title("/compose")
	t.Description("Invalid request test for /compose — field with incorrect data")

	body := map[string]interface{}{
		"format": "t",
		"lines":  []string{"1+2", "3*4"},
	}
	payload, _ := json.Marshal(body)

	utils.RunErrorCheckingTest(t, s.router, "/api/v1/compose", payload, http.StatusInternalServerError)
}

func TestParseSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(ParseRoutesSuite))
}
