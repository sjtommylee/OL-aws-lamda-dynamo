package common

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=1234 port=1337 dbname=go-db sslmode=verify-full"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}

	DB = db
	fmt.Println(DB)
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
