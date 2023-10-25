package routers

import (
	"github.com/gin-gonic/gin"
	"spincity.com/spincity/backend/handlers"
)

func WriteToDB(router *gin.Engine) {
	// router := gin.Default()
	router.POST("/writeDb", handlers.WriteToDB)
	// router.Run("localhost:4000")
}
