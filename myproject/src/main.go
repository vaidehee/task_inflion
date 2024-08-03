package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	initDB()
	router := gin.Default()
	router.GET("/person/:person_id/info", getPersonInfo)
	router.Run(":8080")
	fmt.Println("hi")
}
