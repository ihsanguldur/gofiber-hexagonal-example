package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"weatherApp/config"
	"weatherApp/domain"
)

type mysqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository() domain.Repository {
	var err error
	repo := new(mysqlRepository)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
		config.Config("DB_NAME"))

	repo.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error While Connect to database.")
	}
	fmt.Println("Connection Opened to Database.")

	if err = repo.db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatal("Error While Migrating.")
	}
	fmt.Println("Database Migrated.")

	return repo
}

func (r *mysqlRepository) Create(user *domain.User) error {
	var err error

	if err = r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
