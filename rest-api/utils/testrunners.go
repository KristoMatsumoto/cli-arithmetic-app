package utils

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func RunHealthStatusCheck(t provider.T, router *gin.Engine, route string) {
	t.WithParameters(allure.NewParameter("status", http.StatusOK))

	req, _ := http.NewRequest(http.MethodGet, route, nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	t.Require().Equal(http.StatusOK, resp.Code)
	t.WithNewStep("Check health status", func(sCtx provider.StepCtx) {
		sCtx.Assert().Contains(resp.Body.String(), `"status":"ok"`)
	})
}

func RunTestRouteOK(t provider.T, router *gin.Engine, route string, payload []byte, f func(*httptest.ResponseRecorder) string, contains interface{}) {
	t.WithParameters(allure.NewParameter("status", http.StatusOK))

	req, _ := http.NewRequest(http.MethodPost, route, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	t.WithNewStep("Check response code", func(sCtx provider.StepCtx) {
		sCtx.Require().Equal(http.StatusOK, resp.Code)
	})
	t.WithNewStep("Check JSON structure", func(sCtx provider.StepCtx) {
		sCtx.Assert().Contains(f(resp), contains)
	})
}

func RunErrorCheckingTest(t provider.T, router *gin.Engine, route string, payload []byte, errorHTTP int) {
	t.WithParameters(allure.NewParameter("status", errorHTTP))

	req, _ := http.NewRequest(http.MethodPost, route, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	t.WithNewStep("Checking an invalid request", func(sCtx provider.StepCtx) {
		sCtx.Require().Equal(errorHTTP, resp.Code)
		sCtx.Assert().Contains(resp.Body.String(), `"error":`)
	})
}
