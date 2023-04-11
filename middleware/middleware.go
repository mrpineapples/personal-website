package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"

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

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()

		if ok {
			// Calculate SHA-256 hashes for the provided and expected
			// usernames and passwords.
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			// TODO: Grab these values from env
			expectedUsernameHash := sha256.Sum256([]byte("admin"))
			expectedPasswordHash := sha256.Sum256([]byte("password"))

			// Use the subtle.ConstantTimeCompare() function to check if
			// the provided username and password hashes equal the
			// expected username and password hashes. ConstantTimeCompare
			// will return 1 if the values are equal, or 0 otherwise.
			// Importantly, we should to do the work to evaluate both the
			// username and password before checking the return values to
			// avoid leaking information.
			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			// If the username and password are correct, then call
			// the next handler in the chain. Make sure to return
			// afterwards, so that none of the code below is run.
			if usernameMatch && passwordMatch {
				c.Next()
				return
			}
		}

		c.Header("WWW-Authenticate", `Basic realm="Authorization Required", charset="UTF-8"`)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
