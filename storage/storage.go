package storage

import (
	"conference/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Storage struct {
	db *gorm.DB
}

func New() (*Storage, error) {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}

	log.Println("connected to database")
	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) RunMigration() error {

	er := model.User{}.CreateTable(s.db)
	if er != nil {
		return er
	}

	er = model.Conference{}.CreateTable(s.db)
	if er != nil {
		return er
	}

	er = model.Talk{}.CreateTable(s.db)
	if er != nil {
		return er
	}

	er = model.Speaker{}.CreateTable(s.db)
	if er != nil {
		return er
	}
	er = model.Participant{}.CreateTable(s.db)
	if er != nil {
		return er
	}

	er = model.EditHistory{}.CreateTable(s.db)
	if er != nil {
		return er
	}

	return nil
}
