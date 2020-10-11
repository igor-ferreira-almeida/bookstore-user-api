package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/igor-ferreira-almeida/bookstore-user-api/domain/user"
	"github.com/igor-ferreira-almeida/bookstore-user-api/service"
	exception "github.com/igor-ferreira-almeida/bookstore-user-api/util/error"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ex := exception.NewException("Invalid json body", http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		c.JSON(http.StatusBadRequest, ex)
		return
	}

	result, err := service.CreateUser(user)

	if err != nil {
		//TODO: handle user creation error
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
