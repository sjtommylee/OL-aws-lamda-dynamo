package users

import (
	"fmt"
	"go-service/common"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID             uint   `gorm:"primary_key"`
	Fullname       string `gorm:"column:fullname"`
	Email          string `gorm:"column:email"`
	HashedPassword string `gorm:"column:hashed_password"`
}

func AutoMigrate() {
	fmt.Println("User Migration")
	fmt.Println("User Model", UserModel{})
	db := common.GetDB()
	db.AutoMigrate(&UserModel{})
}
