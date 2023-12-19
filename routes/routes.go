package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrpineapples/personal-website/components"
	"github.com/mrpineapples/personal-website/controllers"
	"github.com/mrpineapples/personal-website/middleware"
)

// InitializeRoutes declares all valid routes
func InitializeRoutes(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "", components.NotFound())
	})
	r.GET("/", controllers.Home)
	r.GET("/about", controllers.About)
	r.GET("/contact", controllers.Contact)

	// blog routes
	r.GET("/blog", controllers.GetArticles)
	r.GET("/blog/:slug", controllers.GetArticle)
	r.POST("/articles", controllers.CreateArticle)
	r.PUT("/articles/:id", controllers.UpdateArticle)
	r.DELETE("/articles/:id", controllers.DeleteArticle)
	// redirect to admin route for ease of use!
	r.GET("/blog/new", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/admin/blog/new")
	})
	r.GET("/blog/:slug/edit", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/admin/blog/"+c.Param("slug")+"/edit")
	})

	// admin routes
	admin := r.Group("/admin", middleware.BasicAuth())
	admin.GET("", controllers.Admin)
	admin.GET("/blog", controllers.GetAdminArticles)
	admin.GET("/blog/new", controllers.NewArticle)
	admin.GET("/blog/:slug/edit", controllers.EditArticle)
}
