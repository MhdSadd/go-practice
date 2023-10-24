package db

import (
	"log"

	"spincity.com/spincity/backend/types"
)

func WriteToDB () {
	db:=InitializeDb()
	vinylRecord := types.RecordData{
			Title: "Sample Record",
			Artist: "Sample Artist",
			Image: "cloudinary/image/4",
	}
	
	

	res:=db.Create(&vinylRecord)

	if res.Error!=nil{
		log.Fatalf("Failed to insert record: %v", res.Error)
	}
	log.Println("Record inserted successfully!")
}