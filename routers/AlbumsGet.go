package routers

import (
	"github.com/gin-gonic/gin"
	"spincity.com/spincity/backend/handlers"
	
)

func AlbumGet () {
	router:= gin.Default()
	router.GET("/getAlbums", handlers.GetAlbums)
	router.Run("localhost:4000")
}