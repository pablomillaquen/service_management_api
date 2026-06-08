package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRBACTest(role string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("user_role", role)
	c.Request, _ = http.NewRequest("GET", "/test", nil)

	RequiredRole("administrator")(c)
	return w
}

func TestRequiredRoleAllows(t *testing.T) {
	w := setupRBACTest("administrator")
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRequiredRoleDenies(t *testing.T) {
	w := setupRBACTest("technician")
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestRequiredRoleMissing(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/test", nil)

	RequiredRole("administrator")(c)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestRequiredRoleMultiple(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("user_role", "technician")
	c.Request, _ = http.NewRequest("GET", "/test", nil)

	RequiredRole("administrator", "technician")(c)
	assert.Equal(t, http.StatusOK, w.Code)
}
