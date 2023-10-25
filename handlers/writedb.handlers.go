package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"spincity.com/spincity/backend/db"
	"spincity.com/spincity/backend/types"
)

func WriteToDB(ctx *gin.Context) {
	db := db.InitializeDb()
	vinylRecord := types.RecordData{
		Title:  "Sample Record",
		Artist: "Sample Artist",
		Image:  "cloudinary/image/4",
	}

	res := db.Create(&vinylRecord)

	if res.Error != nil {
		// log.Fatalf("Failed to insert record: %v", res.Error)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to write to the database",
			"error":   res.Error,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "data inserted successfully"})
}
