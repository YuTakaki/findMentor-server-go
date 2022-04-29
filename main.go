package main

import (
	"github.com/YuTakaki/findMentor-server-go/api/auth"
	"github.com/YuTakaki/findMentor-server-go/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	app := gin.Default()

	router := app.Group("/api")

	auth.AuthApi(router)

	app.Run("localhost:8080")
}