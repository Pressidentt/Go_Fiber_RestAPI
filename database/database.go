package database

import (
	"log"
	"os"
	"testApi/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed connect attempt \n", err.Error())
		os.Exit(2)
	}

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Migrations")
	db.AutoMigrate(&models.User{}, &models.StaticFiles{})

	Database = DbInstance{
		Db: db,
	}

}
