package utils

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"math"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/gin-contrib/multitemplate"
)

var TemplateFS embed.FS

// getReadingTime calculates the amount of time it should take to read the
// given text and returns it in a formatted string (example: "5 min read")
func getReadingTime(text string) string {
	wordsPerMin := 250
	// Match one or more word characters
	reg := regexp.MustCompile(`\w+`)
	// Find all matches in the text
	matches := reg.FindAllString(text, -1)
	count := len(matches)
	time := math.Ceil(float64((count / wordsPerMin)))

	s := "<1 min read"
	if time > 0 {
		s = fmt.Sprintf("%v min read", time)
	}

	return s
}

func LoadTemplates() multitemplate.Renderer {
	templatesDir := "templates/"
	r := multitemplate.NewRenderer()

	funcMap := template.FuncMap{
		"formatAsDate": func(t time.Time) string {
			return t.Format("Jan 2, 2006")
		},
		"getReadingTime": func(text string) string {
			return getReadingTime(text)
		},
	}

	layouts, err := fs.Glob(TemplateFS, templatesDir+"layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	partials, err := fs.Glob(TemplateFS, templatesDir+"partials/*.html")
	if err != nil {
		panic(err.Error())
	}

	views, err := fs.Glob(TemplateFS, templatesDir+"views/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts, partials, and views directories
	for _, view := range views {
		assets := []string{}
		assets = append(assets, layouts...)
		assets = append(assets, partials...)
		files := append(assets, view)

		// should be same name as the root file so that we don't get "incomplete" template error
		tname := filepath.Base(files[0])
		t := template.Must(template.New(tname).Funcs(funcMap).ParseFS(
			TemplateFS,
			files...,
		))

		fileName := filepath.Base(view)
		templateName := strings.TrimSuffix(fileName, ".html")
		r.Add(templateName, t)
	}

	return r
}
