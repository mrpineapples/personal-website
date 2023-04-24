package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mrpineapples/personal-website/middleware"
	"github.com/mrpineapples/personal-website/models"
	"github.com/mrpineapples/personal-website/routes"
	"github.com/mrpineapples/personal-website/utils"
)

var (
	//go:embed templates
	templates embed.FS
	//go:embed public
	staticFiles embed.FS
)

func main() {
	utils.TemplateFS = templates
	utils.LoadEnv()
	staticFS, err := fs.Sub(staticFiles, "public")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Use(middleware.MethodOverride(r))
	r.StaticFS("/public", http.FS(staticFS))
	r.HTMLRender = utils.LoadTemplates()
	routes.InitializeRoutes(r)

	models.InitDB()
	defer models.Pool.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		// handle the error if the server fails to start
		panic(err)
	}
}
