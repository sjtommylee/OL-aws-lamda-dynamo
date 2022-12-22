package main

import (
	"fmt"
	"go-service/common"
	"go-service/posts"
	"go-service/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
	posts.AutoMigrate()
}

func main() {
	fmt.Println("Main running")
	db := common.Init()
	Migrate(db)

	// defer db.Close()

	r := gin.Default()
	api := r.Group("/api")
	posts.PostRegister(api)

	r.Run(":1337")
}
