package services

import (
	"errors"
	"goexperiments/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Printf("failed to connect database: %v", err.Error())
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	return db, err
}

func AddNewUser(user *models.User) (*models.User, error) {
	db, err := initDB()
	if err == nil {
		db.Create(user)
	}
	if user.ID == 0 {
		return user, errors.New("Some problem occured in database, no rows inserted")
	}

	return user, db.Error
}

func GetAllUsers() (*[]models.User, error) {
	db, err := initDB()
	users := make([]models.User, 1)
	if err == nil {
		db.Find(&users)
	}
	return &users, db.Error
}
