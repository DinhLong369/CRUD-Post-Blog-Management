package main

import (
	"post_blog/config"
	"post_blog/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()
	r1 := r.Group("/r1")
	routes.SetupRoute(r1)

	r.Run(":4000")
}
