package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Init(app *gin.Engine, db *pgxpool.Pool) {
	RegisterAuthRouter(app, db)
	RegisterMovieRouter(app, db)
	RegisterAdminRouter(app, db)
	RegisterUserRouter(app, db)
}
