package routers

import (
	"github.com/gin-gonic/gin"
	"spincity.com/spincity/backend/handlers"
)

func WriteToDB () {
	router := gin.Default()
	router.POST("/writeDb", handlers.WriteToDb)
	router.Run("localhost:4000")
}