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

func RegisterUserRouter(app gin.IRouter, db *pgxpool.Pool, rdb *redis.Client) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	g := app.Group("/user")
	g.Use(middleware.VerifyToken(rdb))
	g.Use(middleware.CheckRole("user"))
	{
		g.GET("/", userController.GetProfile)
		g.GET("/history", userController.GetHistory)
		g.PATCH("/password", userController.UpdatePassword)
		g.PATCH("/profile", userController.UpdateProfile)
	}
}
