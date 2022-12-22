package posts

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var err error

func PostRegister(router *gin.RouterGroup) {

	router.GET("/posts", indexAll)
	router.GET("/posts/:id", index)
	router.DELETE("/posts/:id", delete)
	router.PUT("/posts/:id", update)

}

func indexAll(c *gin.Context) {
	// var posts []PostModel
	if err != nil {
		fmt.Println("post get all err block")
	} else {
		c.JSON(200, gin.H{
			"message": "get all",
		})
	}
}

func index(c *gin.Context) {
	if err != nil {
		fmt.Println("get single error block from posts")
	} else {
		c.JSON(200, gin.H{
			"message": "Post get Single",
		})
	}
}

func create(c *gin.Context) {
	if err != nil {
		fmt.Println("post create error block")
	} else {
		c.JSON(200, gin.H{
			"message": "Post create",
		})
	}
}

func delete(c *gin.Context) {
	if err != nil {
		fmt.Println("post delete error block")
	} else {
		c.JSON(200, gin.H{
			"message": "Post delete",
		})
	}
}

func update(c *gin.Context) {
	if err != nil {
		fmt.Println("post update error block")
	} else {
		c.JSON(200, gin.H{
			"message": "Post update",
		})
	}
}
