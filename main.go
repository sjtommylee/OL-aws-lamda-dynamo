package main

import (
	"fmt"
	"go-service/common"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Main running")
	db, err := common.Init()
	if err != nil {
		log.Fatal(err)
		fmt.Println("db fail")
	}
	defer db.Close()

	r := gin.Default()
	r.GET("/api")

	r.Run("localhost:9090")
}
