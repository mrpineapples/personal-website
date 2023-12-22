package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrpineapples/personal-website/components"
)

func Contact(c *gin.Context) {
	c.HTML(http.StatusOK, "", components.Contact())
}
