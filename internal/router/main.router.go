package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func Init(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	RegisterAuthRouter(app, db)
	RegisterMovieRouter(app, db)
	RegisterAdminRouter(app, db)
	RegisterUserRouter(app, db)
	RegisterOrderRouter(app, db)
}
