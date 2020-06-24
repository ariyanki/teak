package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"teak/modules/test"
)

// Must contain "Integration" in method name
func TestLoginUserUnitTest(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.UserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")
	h := LoginUser(c)

	t.Log("run unit")
	
	// Assertions
	if assert.NoError(t, h) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestLoginUserIntegrationTest(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.UserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")
	h := LoginUser(c)
	
	t.Log("run integration")

	// Assertions
	if assert.NoError(t, h) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}