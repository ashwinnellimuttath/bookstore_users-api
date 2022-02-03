package users

import (
	"common_go/oauth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"usersApi/domain/users"
	"usersApi/services"
	"usersApi/utils/errors"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError(fmt.Sprintf("User Id not a number"))
	}
	return userId, nil
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid JSON")
		c.JSON(restErr.Status, restErr)
		fmt.Println(err)
		return
	}
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		//	TODO user creation error
		return

	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func GetUser(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		//c.JSON(err.Status, err)
		return
	}
	userId, userErr := strconv.ParseInt(c.Param("userId"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError(fmt.Sprintf("User Id not a number"))
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement")
}

func UpdateUser(c *gin.Context) {
	userId, idErr := getUserId(c.Param("userId"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid JSON")
		c.JSON(restErr.Status, restErr)
		fmt.Println(err)
		return
	}
	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func DeleteUser(c *gin.Context) {
	userId, idErr := getUserId(c.Param("userId"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	err := services.UsersService.DeleteUser(userId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status:": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON")
		c.JSON(restErr.Status, restErr)
	}
	user, err := services.UsersService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}
