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
	Description string `gorm:"size:2048"`
	Body        string `gorm:"size:2048"`
	Author      string
	DateCreated time.Time `json:"date_created" gorm:"column:date_created" sql:"DEFAULT:current_timestamp"`
}

func FindOne() {

}

func FindMany() ([]PostModel, error) {
	db := common.GetDB()
	var models []PostModel

	if err := db.Find(&models).Error; err != nil {
		return nil, err
	}

	return models, nil
}

func AutoMigrate() {
	fmt.Println("Post Migration")
	fmt.Println("Post", PostModel{})
	db := common.GetDB()
	db.AutoMigrate(&PostModel{})
}
