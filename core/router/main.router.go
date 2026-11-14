package router

import (
	_ "github.com/Albaihaqi354/Tickitz-BE/docs"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.Static("/profile", "./public/profile")
	app.Static("/movie", "./public/movie")

	// Health check di root (untuk Vercel ping)
	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Health check lengkap di /api/health
	app.GET("/api/health", func(c *gin.Context) {
		dbErr := db.Ping(context.Background())
		dbStatus := "OK"
		if dbErr != nil {
			dbStatus = "Error: " + dbErr.Error()
		}
		rdbStatus := "OK"
		if err := rdb.Ping(context.Background()).Err(); err != nil {
			rdbStatus = "Error: " + err.Error()
		}
		c.JSON(http.StatusOK, gin.H{
			"database": dbStatus,
			"redis":    rdbStatus,
			"status":   "ok",
		})
	})

	// Semua route pakai prefix /api agar match path yang diterima dari Vercel
	api := app.Group("/api")
	RegisterAuthRouter(api, db, rdb)
	RegisterMovieRouter(api, db, rdb)
	RegisterAdminRouter(api, db, rdb)
	RegisterUserRouter(api, db, rdb)
	RegisterOrderRouter(api, db, rdb)
}
