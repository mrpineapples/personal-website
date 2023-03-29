package main

import (
	"path/filepath"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/mrpineapples/personal-website/routes"
)

func loadTemplates() multitemplate.Renderer {
	templatesDir := "./templates/"
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	partials, err := filepath.Glob(templatesDir + "partials/*.html")
	if err != nil {
		panic(err.Error())
	}

	views, err := filepath.Glob(templatesDir + "views/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, view := range views {
		assets := []string{}
		assets = append(assets, layouts...)
		assets = append(assets, partials...)
		files := append(assets, view)
		fileName := filepath.Base(view)
		templateName := strings.TrimSuffix(fileName, ".html")
		r.AddFromFiles(templateName, files...)
	}
	return r
}

func main() {
	r := gin.Default()
	// serve the files in the "static" directory
	r.Static("/public", "./public")
	r.HTMLRender = loadTemplates()
	routes.InitializeRoutes(r)

	// start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		// handle the error if the server fails to start
		panic(err)
	}
}
