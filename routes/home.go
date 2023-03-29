package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "home", gin.H{
		"pageTitle": "Michael's Site",
	})

}
