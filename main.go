package main

import (
	
	"spincity.com/spincity/backend/db"
	
	"spincity.com/spincity/backend/routers"
)


func main() {
	db.InitializeDb()
	
	routers.AlbumGet()
	routers.WriteToDB()
}