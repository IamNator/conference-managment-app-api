package storage

import (
	"conference/model"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"os"
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
	logger  *zerolog.Logger
}

func NewUserRepository(s *Storage) IUserRepository {
	_logger := zerolog.New(os.Stdout).With().Str("app", "conf_mgmt_sys").
		Str("module", "user storage").Logger()
	return &UserRepository{
		storage: s,
		logger:  &_logger,
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
	user.Password = user.Password.Hash()
	er := u.storage.db.Create(&user).Error
	if er != nil {
		u.logger.Error().Err(er).Msg("unable to create new user")
		return nil, er
	}

	return &user, nil
}

func (u *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	er := u.storage.db.First(&user, "email = ? ", email).Error
	if er != nil {
		u.logger.Error().Err(er).Msg("unable to find user by email")
		return nil, er
	}

	return &user, nil
}

func (u *UserRepository) UpdateUserPassword(email, password string) error {
	p := model.Password(password)
	p = p.Hash()
	er := u.storage.db.Where("email = ?", email).Updates(model.User{Email: email, Password: p}).Error
	if er != nil {
		u.logger.Error().Err(er).Msg("unable to user password")
		return er
	}

	return nil
}
