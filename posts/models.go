package posts

import (
	"fmt"
	"go-service/common"

	"gorm.io/gorm"
)

type PostModel struct {
	gorm.Model
	Id          string
	Title       string
	Description string
	Body        string
	Author      string
	DateCreated string
}

func AutoMigrate() {
	fmt.Println("Post Migration")
	fmt.Println("Post", PostModel{})
	db := common.GetDB()
	db.AutoMigrate(&PostModel{})
}
