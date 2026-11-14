package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware(c *gin.Context) {
	origin := c.GetHeader("Origin")

	// Default whitelist untuk development
	whiteListOrigin := []string{
		"http://localhost:5000",
		"http://localhost:5050",
		"http://localhost:5173",
		"http://localhost:3000",
	}

	// FRONTEND_URL — tambahkan satu URL tambahan
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL != "" {
		whiteListOrigin = append(whiteListOrigin, frontendURL)
	}

	// ALLOWED_ORIGINS — comma-separated list untuk multiple origins di production
	// Contoh: https://tickitz.vercel.app,https://tickitz-staging.vercel.app
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins != "" {
		for _, o := range strings.Split(allowedOrigins, ",") {
			o = strings.TrimSpace(o)
			if o != "" {
				whiteListOrigin = append(whiteListOrigin, o)
			}
		}
	}

	// Request tanpa Origin header (curl, server-to-server, health check)
	// Tetap diizinkan tapi tanpa CORS headers — browser tidak akan terpengaruh
	if origin == "" {
		c.Next()
		return
	}

	// Cek apakah origin ada di whitelist
	isAllowed := false
	for _, o := range whiteListOrigin {
		if o == origin {
			isAllowed = true
			break
		}
	}

	if isAllowed {
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	} else {
		log.Printf("CORS Blocked: Unknown Origin: %s", origin)
		c.AbortWithStatus(http.StatusForbidden)
	}
}
