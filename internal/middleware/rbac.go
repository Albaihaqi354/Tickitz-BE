package middleware

import (
	"net/http"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/pkg"
	"github.com/gin-gonic/gin"
)

func CheckRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, isExist := c.Get("token")
		if !isExist {
			c.AbortWithStatusJSON(http.StatusForbidden, dto.Response{
				Msg:     "Forbidden Access",
				Success: false,
				Data:    []any{},
				Error:   "Access Denied",
			})
			return
		}

		accessToken, ok := token.(pkg.JWTClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
				Msg:     "Internal Server Error",
				Success: false,
				Data:    []any{},
				Error:   "internal server error",
			})
			return
		}

		if accessToken.Role != role {
			c.AbortWithStatusJSON(http.StatusForbidden, dto.Response{
				Msg:     "Forbidden Access",
				Success: false,
				Data:    []any{},
				Error:   "Access Denied",
			})
			return
		}

		c.Next()
	}
}
