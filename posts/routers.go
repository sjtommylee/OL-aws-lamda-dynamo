package posts

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var err error

func PostRegister(router *gin.RouterGroup) {

	router.GET("/posts/", GetAll)
	router.GET("/posts/:id", GetSingle)
}

func GetAll(c *gin.Context) {
	// var posts []PostModel
	if err != nil {
		fmt.Println("post get all err block")
	} else {
		c.JSON(200, gin.H{
			"message": "get all",
		})
	}
}

func GetSingle(c *gin.Context) {
	if err != nil {
		fmt.Println("get single error block from posts")
	} else {
		c.JSON(200, gin.H{
			"message": "Post get Single",
		})
	}
}
