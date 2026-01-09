package router

import (
	"github.com/Albaihaqi354/Tickitz-BE/internal/controller"
	"github.com/Albaihaqi354/Tickitz-BE/internal/repository"
	"github.com/Albaihaqi354/Tickitz-BE/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Init(app *gin.Engine, db *pgxpool.Pool) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	movieRepository := repository.NewMoviesRepository(db)
	movieService := service.NewMovieService(movieRepository)
	movieController := controller.NewMovieController(movieService)

	app.POST("/auth/register", userController.AddUser)
	app.POST("/auth/login", userController.Login)
	app.GET("/movies/upcoming", movieController.GetUpcomingMovies)
}
