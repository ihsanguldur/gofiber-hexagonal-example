package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"weatherApp/api/config"
	"weatherApp/pkg/entities"
)

func Connect() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
		config.Config("DB_NAME"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error While Connect Database.")
	}
	fmt.Println("Connection Opened to Database.")

	err = DB.AutoMigrate(&entities.User{})
	if err != nil {
		log.Fatal("Error While Migrating.")
	}
	fmt.Println("Database Migrated.")
}
