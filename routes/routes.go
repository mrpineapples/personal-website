package routes

import (
	"github.com/gin-gonic/gin"
)

// InitializeRoutes declares all valid routes
func InitializeRoutes(r *gin.Engine) {
	r.GET("/", home)
	r.GET("/about", about)
	r.GET("/contact", contact)

	// article routes
	r.GET("/articles", getArticles)
	r.GET("/articles/new", newArticle)
	r.GET("/articles/:slug", getArticle)
	r.GET("/articles/:slug/edit", editArticle)
	r.POST("/articles", createArticle)
	r.PATCH("/articles/:id", updateArticle)
	r.DELETE("/articles/:id", deleteArticle)
}
