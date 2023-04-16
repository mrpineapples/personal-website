package routes

import (
	"net/http"

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
	r.PUT("/articles/:id", controllers.UpdateArticle)
	r.DELETE("/articles/:id", controllers.DeleteArticle)
	// redirect to admin route for ease of use!
	r.GET("/articles/:slug/edit", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/admin/articles/"+c.Param("slug")+"/edit")
	})

	// admin routes
	admin := r.Group("/admin", middleware.BasicAuth())
	admin.GET("", controllers.Admin)
	admin.GET("/articles", controllers.GetAdminArticles)
	admin.GET("/articles/new", controllers.NewArticle)
	admin.GET("/articles/:slug/edit", controllers.EditArticle)
}
