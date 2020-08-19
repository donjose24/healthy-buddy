package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres driver
)

func Initialize() *gorm.DB {
	dbURL := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=5432 sslmode=disable connect_timeout=5", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open("postgres", dbURL)

	if err != nil {
		panic(err)
	}

	return db
}
