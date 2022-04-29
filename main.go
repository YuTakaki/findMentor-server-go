package main

import (
	"example/main/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	app := gin.Default()



	app.Run("localhost:8080")
}