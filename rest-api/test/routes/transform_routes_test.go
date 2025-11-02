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

type TransformRoutesSuite struct {
	suite.Suite
	router *gin.Engine
}

func (s *TransformRoutesSuite) BeforeEach(t provider.T) {
	t.Epic("REST API")
	t.Feature("Routes")
	t.Story("Transform routes")
	t.Tags("rest-api", "routes", "transform")

	gin.SetMode(gin.TestMode)
	s.router = gin.Default()
	routes.RegisterTransformRoutes(s.router)
}

func (s *TransformRoutesSuite) TestTransformRoutes_Health(t provider.T) {
	t.Title("/transform/health")
	t.Description("Checks transform routes health")
	utils.RunHealthStatusCheck(t, s.router, "/api/v1/transform/health")
}

func (s *TransformRoutesSuite) TestTransformRoutes_TransformEncode_OK(t provider.T) {
	t.Title("/transform/encode")
	t.Description("Successful request processing test for /transform/encode")

	body := map[string]interface{}{
		"format": "aes",
		"data":   []byte("Hello world"),
	}
	payload, _ := json.Marshal(body)

	utils.RunTestRouteOK(t, s.router, "/api/v1/transform/encode", payload,
		func(resp *httptest.ResponseRecorder) string { return resp.Body.String() }, `"bytes"`)
}
func (s *TransformRoutesSuite) TestTransformRoutes_TransformEncode_BadRequest(t provider.T) {
	t.Title("/transform/encode")
	t.Description("Invalid request test for /transform/encode — missing required field")

	body := map[string]interface{}{
		"data": []byte("Hello world"),
	}
	payload, _ := json.Marshal(body)

	utils.RunErrorCheckingTest(t, s.router, "/api/v1/transform/encode", payload, http.StatusBadRequest)
}
func (s *TransformRoutesSuite) TestTransformRoutes_TransformEncode_InternalServerError(t provider.T) {
	t.Title("/transform/encode")
	t.Description("Invalid request test for /transform/encode — field with incorrect data")

	body := map[string]interface{}{
		"format": "abcd",
		"data":   []byte("Hello world"),
	}
	payload, _ := json.Marshal(body)

	utils.RunErrorCheckingTest(t, s.router, "/api/v1/transform/encode", payload, http.StatusInternalServerError)
}

func (s *TransformRoutesSuite) TestTransformRoutes_TransformEncodeChain_OK(t provider.T) {
	t.Title("/transform/encode/chain")
	t.Description("Successful request processing test for /transform/encode/chain")

	body := map[string]interface{}{
		"formats": []string{"aes", "gzip"},
		"data":    []byte("Hello world"),
	}
	payload, _ := json.Marshal(body)

	utils.RunTestRouteOK(t, s.router, "/api/v1/transform/encode/chain", payload,
		func(resp *httptest.ResponseRecorder) string { return resp.Body.String() }, `"bytes"`)
}
func (s *TransformRoutesSuite) TestTransformRoutes_TransformEncodeChain_BadRequest(t provider.T) {
	t.Title("/transform/encode/chain")
	t.Description("Invalid request test for /transform/encode/chain — missing required field")

	body := map[string]interface{}{
		"data": []byte("Hello world"),
	}
	payload, _ := json.Marshal(body)

	utils.RunErrorCheckingTest(t, s.router, "/api/v1/transform/encode/chain", payload, http.StatusBadRequest)
}
func (s *TransformRoutesSuite) TestTransformRoutes_TransformEncodeChain_InternalServerError(t provider.T) {
	t.Title("/transform/encode/chain")
	t.Description("Invalid request test for /transform/encode/chain — field with incorrect data")

	body := map[string]interface{}{
		"formats": []string{"a", "b"},
		"data":    []byte("Hello world"),
	}
	payload, _ := json.Marshal(body)

	utils.RunErrorCheckingTest(t, s.router, "/api/v1/transform/encode/chain", payload, http.StatusInternalServerError)
}

func (s *TransformRoutesSuite) TestTransformRoutes_TransformDecode_OK(t provider.T) {
	t.Title("/transform/decode")
	t.Description("Successful request processing test for /transform/decode")

	body := map[string]interface{}{
		"format": "aes",
		"data":   []byte("YWJjZGVmZ2hpamtsbW5vcAri6/i3imxlTt9Q8YJUDR9w"),
	}
	payload, _ := json.Marshal(body)

	utils.RunTestRouteOK(t, s.router, "/api/v1/transform/decode", payload,
		func(resp *httptest.ResponseRecorder) string { return resp.Body.String() }, `"bytes"`)
}
func (s *TransformRoutesSuite) TestTransformRoutes_TransformDecode_BadRequest(t provider.T) {
	t.Title("/transform/decode")
	t.Description("Invalid request test for /transform/decode — missing required field")

	body := map[string]interface{}{
		"data": []byte("YWJjZGVmZ2hpamtsbW5vcAri6/i3imxlTt9Q8YJUDR9w"),
	}
	payload, _ := json.Marshal(body)

	utils.RunErrorCheckingTest(t, s.router, "/api/v1/transform/decode", payload, http.StatusBadRequest)
}
func (s *TransformRoutesSuite) TestTransformRoutes_TransformDecode_InternalServerError(t provider.T) {
	t.Title("/transform/decode")
	t.Description("Invalid request test for /transform/decode — field with incorrect data")

	body := map[string]interface{}{
		"format": "abcd",
		"data":   []byte("YWJjZGVmZ2hpamtsbW5vcAri6/i3imxlTt9Q8YJUDR9w"),
	}
	payload, _ := json.Marshal(body)

	utils.RunErrorCheckingTest(t, s.router, "/api/v1/transform/decode", payload, http.StatusInternalServerError)
}

func (s *TransformRoutesSuite) TestTransformRoutes_TransformDecodeChain_OK(t provider.T) {
	t.Title("/transform/decode/chain")
	t.Description("Successful request processing test for /transform/decode/chain")

	body := map[string]interface{}{
		"formats": []string{"aes", "gzip"},
		"data":    []byte("YWJjZGVmZ2hpamtsbW5vcAri6/i3imxlTt9Q8YJUDR9w"),
	}
	payload, _ := json.Marshal(body)

	utils.RunTestRouteOK(t, s.router, "/api/v1/transform/decode/chain", payload,
		func(resp *httptest.ResponseRecorder) string { return resp.Body.String() }, `"bytes"`)
}
func (s *TransformRoutesSuite) TestTransformRoutes_TransformDecodeChain_BadRequest(t provider.T) {
	t.Title("/transform/decode/chain")
	t.Description("Invalid request test for /transform/decode/chain — missing required field")

	body := map[string]interface{}{
		"data": []byte("YWJjZGVmZ2hpamtsbW5vcAri6/i3imxlTt9Q8YJUDR9w"),
	}
	payload, _ := json.Marshal(body)

	utils.RunErrorCheckingTest(t, s.router, "/api/v1/transform/decode/chain", payload, http.StatusBadRequest)
}
func (s *TransformRoutesSuite) TestTransformRoutes_TransformDecodeChain_InternalServerError(t provider.T) {
	t.Title("/transform/decode/chain")
	t.Description("Invalid request test for /transform/decode/chain — field with incorrect data")

	body := map[string]interface{}{
		"formats": []string{"a", "b"},
		"data":    []byte("YWJjZGVmZ2hpamtsbW5vcAri6/i3imxlTt9Q8YJUDR9w"),
	}
	payload, _ := json.Marshal(body)

	utils.RunErrorCheckingTest(t, s.router, "/api/v1/transform/decode/chain", payload, http.StatusInternalServerError)
}

func TestTransformSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(TransformRoutesSuite))
}
