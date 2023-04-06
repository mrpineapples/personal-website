package middleware

import (
	"github.com/gin-gonic/gin"
)

func MethodOverride(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Only run this for POST requests
		if c.Request.Method != "POST" {
			return
		}

		method := c.PostForm("_method")
		if queryMethod := c.Query("_method"); queryMethod != "" {
			method = queryMethod
		}

		if method == "" {
			return
		}

		switch method {
		case "PUT", "PATCH", "DELETE":
			c.Request.Method = method

		// Unsupported method for overriding, request will be processed as POST
		default:
			return
		}

		// Cancel current request and pass modified context to Gin
		c.Abort()
		r.HandleContext(c)
	}
}
