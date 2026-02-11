package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func VerifyToken(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		result := strings.Split(bearerToken, " ")
		if len(result) < 2 || result[0] != "Bearer" {
			log.Println("token is not bearer token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Msg:     "Unauthorized Access",
				Success: false,
				Data:    []any{},
				Error:   "Invalid Token",
			})
			return
		}

		token := result[1]

		rkey := "bian:tickitz:whitelist:" + token
		exists, err := rdb.Exists(c.Request.Context(), rkey).Result()
		if err != nil || exists == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Msg:     "Unauthorized Access",
				Success: false,
				Data:    []any{},
				Error:   "Token not active",
			})
			return
		}

		var jc pkg.JWTClaims
		_, err = jc.VerifyToken(token)
		if err != nil {
			log.Println(err.Error())
			if errors.Is(err, jwt.ErrTokenExpired) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
					Msg:     "Unauthorized Access",
					Success: false,
					Data:    []any{},
					Error:   "Expired Token, Please Login Again",
				})
				return
			}
			if errors.Is(err, jwt.ErrTokenInvalidIssuer) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
					Msg:     "Unauthorized Access",
					Success: false,
					Data:    []any{},
					Error:   "Expired Token, Please Login Again",
				})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
				Msg:     "Internal Server Error",
				Success: false,
				Data:    []any{},
				Error:   "internal server error",
			})
			return
		}
		c.Set("token", jc)
		c.Set("user_id", jc.Id)
		c.Next()
	}
}
