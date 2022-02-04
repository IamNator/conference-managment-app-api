package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Storage struct {
	db *gorm.DB
}

func New() *Storage {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic(err.Error())
	}

	log.Println("connected to database")
	return &Storage{
		db: db,
	}
}
