package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"usersApi/domain/users"
	"usersApi/services"
	"usersApi/utils/errors"
)

func CreateUser(c *gin.Context) {
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	//	TODO
	//	return
	//}
	//error := json.Unmarshal(bytes, &user)
	//if error != nil {
	//	//	TODO
	//	return
	//}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid JSON")
		c.JSON(restErr.Status, restErr)
		fmt.Println(err)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		//	TODO user creation error
		return

	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("userId"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError(fmt.Sprintf("User Id not a number"))
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement")
}
