package router

import (
	"github.com/Albaihaqi354/Tickitz-BE/internal/controller"
	"github.com/Albaihaqi354/Tickitz-BE/internal/middleware"
	"github.com/Albaihaqi354/Tickitz-BE/internal/repository"
	"github.com/Albaihaqi354/Tickitz-BE/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Init(app *gin.Engine, db *pgxpool.Pool) {
	adminRepository := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepository)
	adminController := controller.NewAdminController(adminService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	movieRepository := repository.NewMoviesRepository(db)
	movieService := service.NewMovieService(movieRepository)
	movieController := controller.NewMovieController(movieService)

	app.POST("/auth/register", userController.AddUser)
	app.POST("/auth/login", userController.Login)

	app.GET("/movies/upcoming", movieController.GetUpcomingMovies)
	app.GET("/movies/popular", movieController.GetPopularMovie)
	app.GET("/movies/filter", movieController.GetMovieWithFilter)
	app.GET("/movies/detail", movieController.GetMovieDetail)

	app.GET("/admin/getMovies", middleware.VerifyToken, middleware.CheckRole("admin"), adminController.GetAllMovieAdmin)
	app.DELETE("/admin/movies/:id", middleware.VerifyToken, middleware.CheckRole("admin"), adminController.DeleteMovieAdmin)

	app.GET("/user/profile", middleware.VerifyToken, middleware.CheckRole("user"), userController.GetProfile)
}
