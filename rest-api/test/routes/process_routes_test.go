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

type ProcessRoutesSuite struct {
	suite.Suite
	router *gin.Engine
}

func (s *ProcessRoutesSuite) BeforeEach(t provider.T) {
	t.Epic("REST API")
	t.Feature("Routes")
	t.Story("Process routes")
	t.Tags("rest-api", "routes", "process", "processor")

	gin.SetMode(gin.TestMode)
	s.router = gin.Default()
	routes.RegisterProcessRoutes(s.router)
}

func (s *ProcessRoutesSuite) TestProcessRoutes_Health(t provider.T) {
	t.Title("/process/health")
	t.Description("Checks process routes health")

	utils.RunHealthStatusCheck(t, s.router, "/api/v1/process/health")
}

func (s *ProcessRoutesSuite) TestProcessRoutes_Process_OK(t provider.T) {
	t.Title("/process")
	t.Description("Successful request processing test for /process")

	body := map[string]interface{}{
		"processor": "3",
		"data":      []string{"1+2", "3*4"},
	}
	payload, _ := json.Marshal(body)

	utils.RunTestRouteOK(t, s.router, "/api/v1/process", payload,
		func(resp *httptest.ResponseRecorder) string { return resp.Body.String() }, `result":["3","12"]`)
}

func (s *ProcessRoutesSuite) TestProcessRoutes_Process_BadRequest(t provider.T) {
	t.Title("/process")
	t.Description("Request with incorrect data processing test for /process — missing required field")

	body := map[string]interface{}{
		"data": []string{"1+2"},
	}
	payload, _ := json.Marshal(body)

	utils.RunErrorCheckingTest(t, s.router, "/api/v1/process", payload, http.StatusBadRequest)
}

func (s *ProcessRoutesSuite) TestProcessRoutes_Process_InternalServerError(t provider.T) {
	t.Title("/process")
	t.Description("Request with incorrect data processing test for /process — field with incorrect data")

	body := map[string]interface{}{
		"processor": "invalid",
		"data":      []string{"1+2"},
	}
	payload, _ := json.Marshal(body)

	utils.RunErrorCheckingTest(t, s.router, "/api/v1/process", payload, http.StatusInternalServerError)
}

func TestProcessSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(ProcessRoutesSuite))
}
