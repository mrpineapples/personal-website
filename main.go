package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrpineapples/personal-website/routes"
	"github.com/mrpineapples/personal-website/utils"
)

func main() {
	r := gin.Default()
	r.Static("/public", "./public")
	r.HTMLRender = utils.LoadTemplates()
	routes.InitializeRoutes(r)

	// start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		// handle the error if the server fails to start
		panic(err)
	}
}
