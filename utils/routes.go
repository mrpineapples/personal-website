package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "not-found", gin.H{"PageTitle": "Page not found!"})
}
