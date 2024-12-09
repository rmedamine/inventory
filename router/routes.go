package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPersons(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getPersons Called"})
}
