package server

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasrod100/posgoexpert/RateLimiter/internal/limiter"
)

type WebServer struct {
	Router        *gin.Engine
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        gin.Default(),
		WebServerPort: serverPort,
	}
}

func (ws *WebServer) Run(ctx context.Context, rl limiter.RateLimiterInterface) {
	ws.Router.Use(ws.Middleware(ctx, rl))

	ws.Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Desafio RateLimiter")
	})

	log.Println("Server running on port " + ws.WebServerPort)
	ws.Router.Run(":" + ws.WebServerPort)
}

func (ws *WebServer) Middleware(ctx context.Context, rl limiter.RateLimiterInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		token := c.GetHeader("API_KEY")
		var key string
		var isIp bool

		if token != "" {
			key = "token:" + token
			isIp = false
		} else {
			key = "ip:" + ip
			isIp = true
		}

		allow, err := rl.Allow(ctx, key, isIp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking rate limit"})
			c.Abort()
			return
		}
		if !allow {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "you have reached the maximum number of requests or actions allowed within a certain time frame"})
			c.Abort()
			return
		}

		c.Next()
	}
}
