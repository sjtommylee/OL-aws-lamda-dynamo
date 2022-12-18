package main

import (
	"fmt"
	"go-service/common"
	"go-service/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
}

func main() {
	fmt.Println("Main running")
	db := common.Init()
	Migrate(db)

	// defer db.Close()

	r := gin.Default()
	r.GET("/api")
	r.Run(":1337")
}
