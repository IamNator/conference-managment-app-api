package storage

import "gorm.io/gorm"

type Storage struct {
	db *gorm.DB
}

func New() *Storage {
	return &Storage{}
}
