package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrpineapples/personal-website/controllers"
	"github.com/mrpineapples/personal-website/middleware"
)

// InitializeRoutes declares all valid routes
func InitializeRoutes(r *gin.Engine) {
	r.GET("/", home)
	r.GET("/about", about)
	r.GET("/contact", contact)

	// article routes
	r.GET("/articles", controllers.GetArticles)
	r.GET("/articles/new", controllers.NewArticle)
	r.GET("/articles/:slug", controllers.GetArticle)
	r.GET("/articles/:slug/edit", controllers.EditArticle)
	r.POST("/articles", controllers.CreateArticle)
	r.PATCH("/articles/:id", controllers.UpdateArticle)
	r.DELETE("/articles/:id", controllers.DeleteArticle)

	// admin routes
	admin := r.Group("/admin", middleware.BasicAuth())
	admin.GET("/articles/new", controllers.NewArticle)

}
