package main

import (
	"github.com/unyimmeudemy/go-gin/initializers"
	"github.com/unyimmeudemy/go-gin/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Post{})
}