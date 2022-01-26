package users

import (
	"log"
	"myapp/domain/users"
	"myapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func CreateUser(c *gin.Context) {
	var user users.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		log.Println(saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
