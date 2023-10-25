package main

import (
	"github.com/gin-gonic/gin"
	"spincity.com/spincity/backend/db"

	"spincity.com/spincity/backend/routers"
)

func main() {
	db.InitializeDb()
	r := gin.Default()

	routers.WriteToDB(r)
	routers.AlbumGet(r)

	r.Run("localhost:4000")
}
