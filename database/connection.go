package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDbInstance() *gorm.DB {
	dsn := "host=localhost user=posgres password= dbname=jersey_dev2 port=5432 sslmode=disable"

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("gagal terhubung ke database")
	}

	return db
}
