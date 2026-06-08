package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pablomillaquen/speckit_golang_api/pkg/response"
)

func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Header("Cache-Control", "no-store")
		c.Next()
	}
}

func CORS(allowedOrigins []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		for _, o := range allowedOrigins {
			if o == "*" || o == origin {
				c.Header("Access-Control-Allow-Origin", origin)
				break
			}
		}
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, X-Requested-With")
			c.Header("Access-Control-Max-Age", "86400")
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

type RateLimiter struct {
	mu       sync.Mutex
	requests map[string]int
	resetAt  time.Time
	limit    int
}

func NewRateLimiter(requestsPerMinute int) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string]int),
		resetAt:  time.Now().Add(time.Minute),
		limit:    requestsPerMinute,
	}
}

func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rl.mu.Lock()
		if time.Now().After(rl.resetAt) {
			rl.requests = make(map[string]int)
			rl.resetAt = time.Now().Add(time.Minute)
		}
		ip := c.ClientIP()
		rl.requests[ip]++
		count := rl.requests[ip]
		rl.mu.Unlock()
		if count > rl.limit {
			response.Error(c, 429, "Rate limit exceeded. Try again later.", nil)
			return
		}
		c.Next()
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.InternalError(c, "An unexpected error occurred")
			}
		}()
		c.Next()
	}
}
