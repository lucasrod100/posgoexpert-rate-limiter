package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Mock de RateLimiterInterface
type MockRateLimiter struct {
	allow bool
	err   error
}

func (m *MockRateLimiter) Allow(ctx context.Context, key string, isIp bool) (bool, error) {
	if m.err != nil {
		return false, m.err
	}
	return m.allow, nil
}

func TestMiddleware_AllowRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockLimiter := &MockRateLimiter{allow: true} // Simula que a requisição é permitida
	ws := NewWebServer("8080")

	ws.Router.Use(ws.Middleware(context.Background(), mockLimiter))
	ws.Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Desafio RateLimiter")
	})

	req, _ := http.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	ws.Router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Desafio RateLimiter", rec.Body.String())
}

func TestMiddleware_BlockRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockLimiter := &MockRateLimiter{allow: false}
	ws := NewWebServer("8080")

	ws.Router.Use(ws.Middleware(context.Background(), mockLimiter))
	ws.Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Desafio RateLimiter")
	})

	req, _ := http.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	ws.Router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusTooManyRequests, rec.Code)
}
