package db

import (
	"log"

	"fmt"

	"spincity.com/spincity/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"github.com/joho/godotenv"
)


func getEnv (key string) string {
	err:= godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}
	return os.Getenv(key)
}

func InitializeDb() *gorm.DB {
	password:= getEnv("POSTGRES_PASSWORD")

	dbURL:= fmt.Sprintf("postgres://postgres:%s@localhost:5432/spincity", password)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil{
		log.Fatalln(err)
	} else{
		fmt.Println("Connected to db Successfully")
	}


	db.AutoMigrate(&models.Vinylrecord{})


return db


}