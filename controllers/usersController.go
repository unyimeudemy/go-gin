package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/unyimmeudemy/go-gin/initializers"
	"github.com/unyimmeudemy/go-gin/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	//get email and password from request body
	var reqBody struct {
		Email    string
		Password string
	}

	if err := c.Bind(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})

		return
	}

	// hash password
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 10)
	if error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})

		return
	}

	// create user on db
	user := models.User{Email: reqBody.Email, Password: string(hashedPassword)}

	result := initializers.DB.Create(&user)

	if result.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})

		return
	}

	// send response
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context){
	// get email and password
	var reqBody struct {
		Email    string
		Password string
	}

	if err := c.Bind(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})

		return
	}

	//get requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", reqBody.Email)

	if user.ID == 0{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email or password not correct",
		})

		return
	}


	//compare provided password with one in db
	error := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password))
	if error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email or password not correct",
		})

		return
	}

	//generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})

		return
	}

	//send it back without cookie
	// c.JSON(http.StatusOK, gin.H{
	// 	"token": tokenString,
	// })

	//************************************************************
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})

}


func Validate(c *gin.Context){
	

	user, _ := c.Get("user")

	// to work with user, cast to User as shown below
	fmt.Println(user.(models.User).Email)

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}