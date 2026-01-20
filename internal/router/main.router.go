package router

import (
	_ "github.com/Albaihaqi354/Tickitz-BE/docs"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	RegisterAuthRouter(app, db)
	RegisterMovieRouter(app, db, rdb)
	RegisterAdminRouter(app, db)
	RegisterUserRouter(app, db)
	RegisterOrderRouter(app, db)
}
