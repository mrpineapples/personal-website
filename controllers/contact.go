package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Contact(c *gin.Context) {
	c.HTML(http.StatusOK, "contact", nil)
}
