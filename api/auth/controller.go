package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sample(c *gin.Context){
	c.IndentedJSON(http.StatusOK, gin.H{"yes": "yes"})

}