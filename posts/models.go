package posts

import (
	"fmt"
	"go-service/common"
	"time"

	"gorm.io/gorm"
)

type PostModel struct {
	gorm.Model
	ID          string `gorm:"unique_index"`
	Title       string
	Description string ``
	Body        string ``
	Author      string
	DateCreated time.Time `json:"date_created" gorm:"column:date_created" sql:"DEFAULT:current_timestamp"`
}

func AutoMigrate() {
	fmt.Println("Post Migration")
	fmt.Println("Post", PostModel{})
	db := common.GetDB()
	db.AutoMigrate(&PostModel{})
}
