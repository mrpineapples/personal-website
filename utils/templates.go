package utils

import (
	"html/template"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-contrib/multitemplate"
)

func LoadTemplates() multitemplate.Renderer {
	templatesDir := "./templates/"
	r := multitemplate.NewRenderer()

	funcMap := template.FuncMap{
		"formatAsDate": func(t time.Time) string {
			return t.Format("January 2, 2006")
		},
	}

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

	// Generate our templates map from our layouts, partials, and views directories
	for _, view := range views {
		assets := []string{}
		assets = append(assets, layouts...)
		assets = append(assets, partials...)
		files := append(assets, view)
		fileName := filepath.Base(view)
		templateName := strings.TrimSuffix(fileName, ".html")
		r.AddFromFilesFuncs(templateName, funcMap, files...)
	}
	return r
}
