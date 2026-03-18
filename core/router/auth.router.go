package router

import (
	"github.com/Albaihaqi354/Tickitz-BE/core/controller"
	"github.com/Albaihaqi354/Tickitz-BE/core/middleware"
	"github.com/Albaihaqi354/Tickitz-BE/core/repository"
	"github.com/Albaihaqi354/Tickitz-BE/core/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func RegisterAuthRouter(app gin.IRouter, db *pgxpool.Pool, rdb *redis.Client) {
	authRepository := repository.NewAuthRepository(db, rdb)
	authService := service.NewAuthService(authRepository, rdb)
	authController := controller.NewAuthController(authService)

	g := app.Group("/auth")
	g.POST("/register", authController.Register)
	g.POST("/login", authController.Login)
	g.DELETE("/logout", middleware.VerifyToken(rdb), authController.Logout)
}
