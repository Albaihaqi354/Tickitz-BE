package middleware

import (
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware(c *gin.Context) {
	defer c.Next()

	whiteListOrigin := []string{"http://localhost:5000", "http://localhost:5050"}
	origin := c.GetHeader("Origin")

	if slices.Contains(whiteListOrigin, origin) {
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Credentials", "true")

		allowedHeaders := []string{
			"Content-Type", "Authorization",
		}
		c.Header("Access-Control-Allow-Headers", strings.Join(allowedHeaders, ", "))

		allowedMethods := []string{
			http.MethodPost, http.MethodOptions, http.MethodGet, http.MethodPut,
			http.MethodDelete, http.MethodPatch,
		}
		c.Header("Access-Control-Allow-Methods", strings.Join(allowedMethods, ", "))

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
	} else {
		log.Printf("Origin is not in the whitelist: %s", origin)
	}
}
