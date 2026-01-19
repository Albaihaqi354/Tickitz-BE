package router

import (
	"github.com/Albaihaqi354/Tickitz-BE/internal/controller"
	"github.com/Albaihaqi354/Tickitz-BE/internal/middleware"
	"github.com/Albaihaqi354/Tickitz-BE/internal/repository"
	"github.com/Albaihaqi354/Tickitz-BE/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func RegisterUserRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, rdb)
	userController := controller.NewUserController(userService)

	g := app.Group("/user")
	g.Use(middleware.VerifyToken)
	g.Use(middleware.CheckRole("user"))
	{
		g.GET("/profile", userController.GetProfile)
		g.GET("/history", userController.GetHistory)
		g.PATCH("/password", userController.UpdatePassword)
		g.PATCH("/profile", userController.UpdateProfile)
	}
}
