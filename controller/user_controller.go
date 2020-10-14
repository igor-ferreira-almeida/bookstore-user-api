package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/igor-ferreira-almeida/bookstore-user-api/domain/user"
	"github.com/igor-ferreira-almeida/bookstore-user-api/service"
	exception "github.com/igor-ferreira-almeida/bookstore-user-api/util/exception"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ex := exception.NewBadRequestException("Invalid json body")
		c.JSON(ex.Status, ex)
		return
	}

	result, ex := service.CreateUser(user)

	if ex != nil {
		c.JSON(ex.Status, ex)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		ex := exception.NewBadRequestException("user id param should be a number")
		c.JSON(ex.Status, ex)
		return
	}

	result, ex := service.GetUser(id)

	if ex != nil {
		c.JSON(ex.Status, ex)
		return
	}

	c.JSON(http.StatusOK, result)
}
