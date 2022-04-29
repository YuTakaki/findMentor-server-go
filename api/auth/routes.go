package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthApi(r *gin.RouterGroup) {
	router := r.Group("/auth")

	router.GET("/", sample)

}