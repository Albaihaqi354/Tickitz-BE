package main

import (
	"log"

	"github.com/Albaihaqi354/Tickitz-BE/internal/config"
	"github.com/Albaihaqi354/Tickitz-BE/internal/middleware"
	"github.com/Albaihaqi354/Tickitz-BE/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed To Load env")
	}

	db, err := config.InitDb()
	if err != nil {
		log.Println("Failed to Connect Database")
		return
	}

	app := gin.Default()
	app.Use(middleware.CORSMiddleware)
	router.Init(app, db)
	app.Run("localhost:5000")
}
