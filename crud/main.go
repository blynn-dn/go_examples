package main

import (
	"github.com/blynn-dn/go_examples/crud/controllers"
	"github.com/blynn-dn/go_examples/crud/initializer"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.Run() // listen and serve on 0.0.0.0:8080 by default if env variables are not set
}
