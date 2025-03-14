package main

import "github.com/gin-gonic/gin"

func main() {
	Init()
	router := gin.Default()

	RouterPersonSetup(router)

	router.Run(":8080")
}

func migrate() {
	db.AutoMigrate(&Person{})
}
