package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Init() (*gorm.DB, error) {
    dsn := "host=localhost user=root password=secret dbname=todo_db sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
		log.Fatal(err)
        return nil, err
    }
    DB = db
    return db, nil
}