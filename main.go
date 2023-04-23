package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mrpineapples/personal-website/middleware"
	"github.com/mrpineapples/personal-website/models"
	"github.com/mrpineapples/personal-website/routes"
	"github.com/mrpineapples/personal-website/utils"
)

func main() {
	utils.LoadEnv()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.Use(middleware.MethodOverride(r))
	r.Static("/public", "./public")
	r.HTMLRender = utils.LoadTemplates()
	routes.InitializeRoutes(r)

	models.InitDB()
	defer models.Pool.Close()

	// start the server on port 8080
	if err := r.Run(":" + port); err != nil {
		// handle the error if the server fails to start
		panic(err)
	}
}
