package storage

import (
	"conference/model"
	"gorm.io/gorm"
)

//go:generate mockgen -source user.go -destination ./mock/user.go -package mock IUserRepository
type IUserRepository interface {
	CreateUser(user model.User) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUserPassword(email, password string) error
	WithTx(db *gorm.DB) IUserRepository
}

type UserRepository struct {
	storage *Storage
}

func NewUserRepository(s *Storage) IUserRepository {
	return &UserRepository{
		storage: s,
	}
}

func (u *UserRepository) WithTx(db *gorm.DB) IUserRepository {
	return &UserRepository{
		storage: &Storage{
			db: db,
		},
	}
}

func (u *UserRepository) CreateUser(user model.User) (*model.User, error) {
	return &user, u.storage.db.Create(&user).Error
}

func (u *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	return &user, u.storage.db.First(&user, "email = ? ", email).Error
}

func (u *UserRepository) UpdateUserPassword(email, password string) error {
	p := model.Password(password)
	p = p.Hash()
	return u.storage.db.Where("email = ?", email).Updates(model.User{Email: email, Password: p}).Error
}
