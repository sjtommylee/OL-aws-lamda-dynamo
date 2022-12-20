package users

import "github.com/gin-gonic/gin"

var err error

func UserRegister(router *gin.RouterGroup) {
	router.GET("/users/", GetAll)
	router.GET("/users/:id", GetSingle)
}

func GetAll(c *gin.Context) {
	if err != nil {
		println("error")
	} else {
		c.JSON(200, gin.H{
			"message": "get all",
		})
	}
}

func GetSingle(c *gin.Context) {
	if err != nil {
		println("error")
	} else {
		c.JSON(200, gin.H{
			"message": "get all",
		})
	}
}
