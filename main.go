package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unyimmeudemy/go-gin/controllers"
	"github.com/unyimmeudemy/go-gin/initializers"
)

func init(){
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	r.POST("/post", controllers.CreatePost)
	r.GET("/post", controllers.GetPosts)
	r.GET("/post/:id", controllers.GetPost)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)
	r.Run()
}