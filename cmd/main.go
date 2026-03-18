package main

import (
	"log"

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
		log.Println("Failed to Connect Database")
		return
	}
	defer db.Close()

	rdb := config.InitRedis()
	defer rdb.Close()

	app := gin.Default()

	app.Use(middleware.CORSMiddleware)
	router.Init(app, db, rdb)
	app.Run(":5000")
}
