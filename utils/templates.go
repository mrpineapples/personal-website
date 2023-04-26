package utils

import (
	"fmt"
	"html/template"
	"math"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/gin-contrib/multitemplate"
)

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
