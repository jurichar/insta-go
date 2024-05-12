package database

import (
	"backend/models"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to database!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migration...")
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{Db: db}

}
