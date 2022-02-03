package app

import (
	"usersApi/controllers"
	"usersApi/controllers/users"
)

func mapURLS() {
	router.GET("/ping", controllers.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:userId", users.GetUser)
	router.PUT("/users/:userId", users.UpdateUser)
	router.PATCH("/users/:userId", users.UpdateUser)
	router.DELETE("/users/:userId", users.DeleteUser)
	router.GET("/internal/users/search", users.Search)
	router.POST("/users/login", users.Login)

	router.GET("/users/search", users.SearchUser)
}
