package main

import (
	"context"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/mrpineapples/personal-website/database"
	"github.com/mrpineapples/personal-website/middleware"
	"github.com/mrpineapples/personal-website/routes"
	"github.com/mrpineapples/personal-website/utils"
)

type TemplRender struct {
	Code int
	Data templ.Component
}

func (t TemplRender) Render(w http.ResponseWriter) error {
	t.WriteContentType(w)
	w.WriteHeader(t.Code)
	if t.Data != nil {
		return t.Data.Render(context.Background(), w)
	}
	return nil
}

func (t TemplRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func (t *TemplRender) Instance(name string, data interface{}) render.Render {
	if templData, ok := data.(templ.Component); ok {
		return &TemplRender{
			Code: http.StatusOK,
			Data: templData,
		}
	}
	return nil
}

func main() {
	utils.LoadEnv()

	r := gin.Default()
	r.Use(middleware.MethodOverride(r))
	r.Static("/public", "./public")
	r.HTMLRender = &TemplRender{}
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
