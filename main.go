package main

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
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

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home", gin.H{
			"pageTitle": "Michael's Site",
		})
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about", gin.H{
			"pageTitle": "Michael's Site | About",
		})
	})

	r.GET("/contact", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"contact": gin.H{
				"github":   "https://github.com/mrpineapples",
				"linkedin": "https://www.linkedin.com/in/michaelmiranda18/",
				"twitter":  "https://twitter.com/mrpineapples24",
			},
		})
	})

	// start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		// handle the error if the server fails to start
		panic(err)
	}
}
