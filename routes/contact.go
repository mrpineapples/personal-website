package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func contact(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"contact": gin.H{
			"github":   "https://github.com/mrpineapples",
			"linkedin": "https://www.linkedin.com/in/michaelmiranda18/",
			"twitter":  "https://twitter.com/mrpineapples24",
		},
	})
}
