package app

import (
	"usersApi/controllers"
	"usersApi/controllers/users"
)

func mapURLS() {
	router.GET("/ping", controllers.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:userId", users.GetUser)
	router.GET("/users/search", users.SearchUser)
}
