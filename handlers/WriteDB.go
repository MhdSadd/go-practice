package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "spincity.com/spincity/backend/db"
)

func WriteToDb (ctx *gin.Context) {
    // Call the function to write to the database
    err := db.WriteToDB

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to write to the database",
        })
        return
    }

    ctx.Status(http.StatusOK)
}
