package setup

import (
	"../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDatabase() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=fiberTest1 sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&models.User{})
}
