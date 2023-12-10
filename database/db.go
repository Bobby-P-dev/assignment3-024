package database

import (
	"log"

	"github.com/Bobby-P-dev/assignment3-024/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=flora.db.elephantsql.com user=ruydxddo password=OZVbhMAIHq8IlqS6CzGqujj8_06wXn_3 dbname=ruydxddo port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(&models.Microservice{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
}

func GetDB() *gorm.DB {
	return DB
}
