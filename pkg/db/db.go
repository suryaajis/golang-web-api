package db

import (
	"golang-web-api/pkg/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "postgres://postgres:postgres@localhost:5432/profession-api"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Failed")
	}

	db.AutoMigrate(&model.Profession{})

	return db
}
