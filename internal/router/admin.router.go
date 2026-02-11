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

func RegisterAdminRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	adminRepository := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepository)
	adminController := controller.NewAdminController(adminService)

	g := app.Group("/admin")
	g.Use(middleware.VerifyToken(rdb))
	g.Use(middleware.CheckRole("admin"))
	{
		g.GET("/", adminController.GetAllMovieAdmin)
		g.POST("/movies", adminController.CreateMovieAdmin)
		g.DELETE("/movies/:id", adminController.DeleteMovieAdmin)
		g.PATCH("/movies/:id", adminController.UpdateMovieAdmin)
	}
}
