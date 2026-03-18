package router

import (
	_ "github.com/Albaihaqi354/Tickitz-BE/docs"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"context"
)

func Init(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.Static("/profile", "./public/profile")
	app.Static("/movie", "./public/movie")

	// Wrap everything in /api for Vercel compatibility
	api := app.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			dbErr := db.Ping(context.Background())
			dbStatus := "OK"
			if dbErr != nil {
				dbStatus = "Error: " + dbErr.Error()
			}

			rdbStatus := "OK"
			rdbErr := rdb.Ping(context.Background()).Err()
			if rdbErr != nil {
				rdbStatus = "Error: " + rdbErr.Error()
			}

			c.JSON(http.StatusOK, gin.H{
				"database": dbStatus,
				"redis":    rdbStatus,
				"status":   "Backend is running",
			})
		})

		RegisterAuthRouter(api, db, rdb)
		RegisterMovieRouter(api, db, rdb)
		RegisterAdminRouter(api, db, rdb)
		RegisterUserRouter(api, db, rdb)
		RegisterOrderRouter(api, db, rdb)
	}
}
