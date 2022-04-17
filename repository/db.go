package repository

import (
	"fmt"

	"github.com/onurdogustemel/final-project/model"

	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//DB connection to Database
func DB() *gorm.DB {

	database := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOSTNAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(database), &gorm.Config{})
	if err != nil {
		return nil
	}

	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{})
	return db

}
