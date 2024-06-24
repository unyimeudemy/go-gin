package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/unyimmeudemy/go-gin/initializers"
	"github.com/unyimmeudemy/go-gin/models"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie from request
	tokenString, err := c.Cookie("Authorization")
	// fmt.Println("===============",tokenString)
	if err != nil {
		fmt.Println("======================>", err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//decode cookie
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
	
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
	//check exp date
	if float64(time.Now().Unix()) > claims["exp"].(float64){
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//find the user in subject
	var user models.User
	initializers.DB.First(&user, claims["sub"])
	
	if user.ID == 0{
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//attach to req body
	c.Set("user", user)

	//continue

	c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}