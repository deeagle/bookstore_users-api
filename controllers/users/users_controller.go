package users

import (
	"encoding/json"
	"io/ioutil"
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
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return
	}
	json_err := json.Unmarshal(bytes, &user)
	if json_err != nil {
		log.Println(json_err)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		log.Println(saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
