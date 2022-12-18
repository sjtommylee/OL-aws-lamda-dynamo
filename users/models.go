package users

import (
	"go-service/common"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             uint
	Fullname       string
	Email          string
	HashedPassword string
}

func AutoMigrateDB() {
	db := common.GetDB()
	db.AutoMigrate(&User{})
}
