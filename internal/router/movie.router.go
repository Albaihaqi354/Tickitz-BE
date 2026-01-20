package router

import (
	"github.com/Albaihaqi354/Tickitz-BE/internal/controller"
	"github.com/Albaihaqi354/Tickitz-BE/internal/repository"
	"github.com/Albaihaqi354/Tickitz-BE/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func RegisterMovieRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	movieRepository := repository.NewMoviesRepository(db)
	movieService := service.NewMovieService(movieRepository, rdb)
	movieController := controller.NewMovieController(movieService)

	g := app.Group("/movies")
	g.GET("/upcoming", movieController.GetUpcomingMovies)
	g.GET("/popular", movieController.GetPopularMovie)
	g.GET("/filter", movieController.GetMovieWithFilter)
	g.GET("/detail/:id", movieController.GetMovieDetail)
}
