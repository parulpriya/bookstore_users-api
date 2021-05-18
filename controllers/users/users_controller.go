package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/parulpriya/bookstore_users-api/domain/users"
	"github.com/parulpriya/bookstore_users-api/services"
	"github.com/parulpriya/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := errors.BadRequestErr("invalid json")
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	fmt.Println(user)

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.StatusCode, saveErr)
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context)  {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.BadRequestErr("invalid user_id")
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
	}
	c.JSON(http.StatusOK, user)
}
