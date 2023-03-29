package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func about(c *gin.Context) {
	c.HTML(http.StatusOK, "about", gin.H{
		"pageTitle": "Michael's Site | About",
	})
}
