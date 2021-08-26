package common

import (
	"fmt"
	"github.com/yusrilmr/todolist/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

// GetDB gets database access
func GetDB() *gorm.DB {
	if db == nil {
		var err error

		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			AppConfig.DBHost, AppConfig.DBPort, AppConfig.DBUser,
			AppConfig.DBPwd, AppConfig.Database)

		db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
		if err != nil {
			log.Fatalf("[GetDB]: %s\n", err)
		}
	}
	return db
}

// createDBConnection creates database session
func createDBConnection() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		AppConfig.DBHost, AppConfig.DBPort, AppConfig.DBUser,
		AppConfig.DBPwd, AppConfig.Database)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		log.Fatalf("[createDBConnection]: %s\n", err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Task{}, &models.Label{})
	log.Println("DB Connection established...")
}