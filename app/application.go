package app

import (
	"github.com/gin-gonic/gin"
	"usersApi/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapURLS()
	logger.Info("About to start the application")
	router.Run(":8081")
}
