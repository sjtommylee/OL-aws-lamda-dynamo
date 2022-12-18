package users

import (
	"fmt"
	"go-service/common"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID             uint
	Fullname       string
	Email          string
	HashedPassword string
}

func AutoMigrate() {
	fmt.Println("User Migration")
	fmt.Println("User Model", UserModel{})
	db := common.GetDB()
	db.AutoMigrate(&UserModel{})
}
