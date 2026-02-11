package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware(c *gin.Context) {
	origin := c.GetHeader("Origin")
	whiteListOrigin := []string{
		"http://localhost:5000",
		"http://localhost:5050",
		"http://localhost:5173",
		"http://localhost:3000",
		"http://192.168.50.121:3000",
	}

	isAllowed := false
	if origin == "" {
		isAllowed = true
	} else {
		for _, o := range whiteListOrigin {
			if o == origin {
				isAllowed = true
				break
			}
		}
	}

	if isAllowed {
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
		}
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
