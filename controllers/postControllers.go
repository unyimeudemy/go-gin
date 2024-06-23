package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/unyimmeudemy/go-gin/initializers"
	"github.com/unyimmeudemy/go-gin/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}


func CreatePost(c *gin.Context) {
	// get data from req body
	var reqBody struct {
		Title string
		Body string
	}

	c.Bind(&reqBody)

	// create post on db
	post := models.Post{Title: reqBody.Title, Body: reqBody.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil{
		c.Status(400)
		return
	}

	//return created post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context){

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}


func GetPost(c *gin.Context){

	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context){
	//get the id from the url
	id := c.Param("id")

	//get datat from req body
	var reqBody struct {
		Title string
		Body string
	}

	c.Bind(&reqBody)

	//find post to be updated
	var post models.Post
	initializers.DB.First(&post, id)

	//update post and save 
	initializers.DB.Model(&post).Updates(models.Post{
		Title: reqBody.Title,
		Body: reqBody.Body,
	})


	// respond with updated data
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context){

	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)
	c.Status(200)

}
