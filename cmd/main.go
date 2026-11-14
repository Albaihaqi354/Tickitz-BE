package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Albaihaqi354/Tickitz-BE/core/config"
	"github.com/Albaihaqi354/Tickitz-BE/core/middleware"
	"github.com/Albaihaqi354/Tickitz-BE/core/router"
	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

// @title           Tickitz Back End
// @version         1.0
// @description     Back End using go with gin.
// @host      		localhost:5000
// @BasePath 		/

// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description                Type "Bearer" followed by a space and then your token.

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed To Load env")
	}

	db, err := config.InitDb()
	if err != nil {
		log.Println("Failed to Connect Database:", err)
		return
	}
	defer db.Close()

	rdb := config.InitRedis()
	defer rdb.Close()

	app := gin.Default()
	app.MaxMultipartMemory = 8 << 20 // 8 MB limit untuk file upload

	app.Use(middleware.CORSMiddleware)
	router.Init(app, db, rdb)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      app,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	go func() {
		log.Printf("Server running on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	log.Println("Server stopped")
}
