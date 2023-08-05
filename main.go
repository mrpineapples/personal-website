package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mrpineapples/personal-website/database"
	"github.com/mrpineapples/personal-website/middleware"
	"github.com/mrpineapples/personal-website/routes"
	"github.com/mrpineapples/personal-website/utils"
)

func main() {
	utils.LoadEnv()

	r := gin.Default()
	r.Use(middleware.MethodOverride(r))
	r.Static("/public", "./public")
	r.HTMLRender = utils.LoadTemplates()
	routes.InitializeRoutes(r)

	database.InitDB()
	defer database.Pool.Close()
	database.RunMigrations()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
