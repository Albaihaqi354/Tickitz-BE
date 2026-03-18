package router

import (
	"github.com/Albaihaqi354/Tickitz-BE/core/controller"
	"github.com/Albaihaqi354/Tickitz-BE/core/repository"
	"github.com/Albaihaqi354/Tickitz-BE/core/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func RegisterMovieRouter(app gin.IRouter, db *pgxpool.Pool, rdb *redis.Client) {
	movieRepository := repository.NewMoviesRepository(db)
	movieService := service.NewMovieService(movieRepository, rdb)
	movieController := controller.NewMovieController(movieService)

	g := app.Group("/movies")
	g.GET("/", movieController.GetMovieWithFilter)
	g.GET("/upcoming", movieController.GetUpcomingMovies)
	g.GET("/popular", movieController.GetPopularMovie)
	g.GET("/detail/:id", movieController.GetMovieDetail)
}
