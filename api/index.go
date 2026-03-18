package handler

import (
	"log"
	"net/http"
	"sync"

	"github.com/Albaihaqi354/Tickitz-BE/internal/config"
	"github.com/Albaihaqi354/Tickitz-BE/internal/middleware"
	"github.com/Albaihaqi354/Tickitz-BE/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

var (
	app  *gin.Engine
	db   *pgxpool.Pool
	rdb  *redis.Client
	once sync.Once
)

func initApp() {
	once.Do(func() {
		// Initialize DB
		var err error
		db, err = config.InitDb()
		if err != nil {
			log.Println("Vercel: Failed to Connect Database:", err)
		}

		// Initialize Redis
		rdb = config.InitRedis()

		// Initialize Gin
		gin.SetMode(gin.ReleaseMode)
		app = gin.New()
		app.Use(gin.Recovery())
		
		// Add CORS middleware
		app.Use(middleware.CORSMiddleware)

		// Initialize Routes
		router.Init(app, db, rdb)
	})
}

// Handler is the entry point for Vercel Serverless Functions
func Handler(w http.ResponseWriter, r *http.Request) {
	initApp()
	app.ServeHTTP(w, r)
}
