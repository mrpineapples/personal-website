package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrpineapples/personal-website/controllers"
	"github.com/mrpineapples/personal-website/middleware"
)

// InitializeRoutes declares all valid routes
func InitializeRoutes(r *gin.Engine) {
	r.GET("/", controllers.Home)
	r.GET("/about", controllers.About)
	r.GET("/contact", controllers.Contact)

	// article public routes
	r.GET("/articles", controllers.GetArticles)
	r.GET("/articles/:slug", controllers.GetArticle)
	r.POST("/articles", controllers.CreateArticle)
	r.PATCH("/articles/:id", controllers.UpdateArticle)
	r.DELETE("/articles/:id", controllers.DeleteArticle)

	// admin routes
	admin := r.Group("/admin", middleware.BasicAuth())
	admin.GET("", controllers.Admin)
	admin.GET("/articles", controllers.GetAdminArticles)
	admin.GET("/articles/new", controllers.NewArticle)
	admin.GET("/articles/:slug/edit", controllers.EditArticle)
}
