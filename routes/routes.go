package routes

import "github.com/gin-gonic/gin"

// InitializeRoutes declares all valid routes
func InitializeRoutes(r *gin.Engine) {
	r.GET("/", home)
	r.GET("/about", about)
	r.GET("/contact", contact)
	r.GET("/articles", getArticles)
	r.GET("/articles/:slug", getArticle)
	r.GET("/articles/new", newArticle)
	r.POST("/articles", createArticle)
}
