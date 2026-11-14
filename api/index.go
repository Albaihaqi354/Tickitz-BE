package handler

import (
	"log"
	"net/http"
	"sync"

	"github.com/Albaihaqi354/Tickitz-BE/core/config"
	"github.com/Albaihaqi354/Tickitz-BE/core/middleware"
	"github.com/Albaihaqi354/Tickitz-BE/core/router"
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
		// Initialize DB (supports DATABASE_URL dari Neon)
		// pgxpool tidak langsung connect — koneksi dibuat saat query pertama.
		// Jadi InitDb() bisa succeed walaupun DB sementara down.
		var err error
		db, err = config.InitDb()
		if err != nil {
			log.Println("Vercel: Failed to init DB pool:", err)
		}

		// Initialize Redis
		rdb = config.InitRedis()

		// Initialize Gin
		gin.SetMode(gin.ReleaseMode)
		app = gin.New()
		app.Use(gin.Recovery())
		app.Use(middleware.CORSMiddleware)
		app.MaxMultipartMemory = 8 << 20 // 8 MB limit untuk file upload

		// Register routes — selalu daftarkan agar app tidak panic.
		// Kalau db nil, query akan return error dan controller kirim 500.
		router.Init(app, db, rdb)
	})
}

// Handler is the entry point for Vercel Serverless Functions
func Handler(w http.ResponseWriter, r *http.Request) {
	initApp()
	app.ServeHTTP(w, r)
}
