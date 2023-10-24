package handlers

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
	"spincity.com/spincity/backend/packages/seed"
)

func GetAlbums (ctx *gin.Context){
	ctx.JSON(http.StatusOK, seed.Records)
}