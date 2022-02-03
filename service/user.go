package service

import (
	"conference/model"
	"conference/storage"
)

//go:generate mockgen -source user.go -destination ./mock/user_service.go -package mock IUserService
type IUserService interface {
	Login(req model.UserLoginReq) (*model.UserAuthResponse, error)
	RegisterUser(req model.UserSignUpReq) (*model.UserAuthResponse, error)
	LogOut(refresh string) error
}

type UserService struct {
	userRepo storage.IUserRepository
}

func NewUserService(s *storage.Storage) IUserService {
	return &UserService{
		userRepo: storage.NewUserRepository(s),
	}
}

func (c *UserService) Login(req model.UserLoginReq) (*model.UserAuthResponse, error) {
	user, _ := c.userRepo.GetUserByEmail(req.Email)

	return &model.UserAuthResponse{
		User: *user,
	}, nil
}

func (c *UserService) RegisterUser(req model.UserSignUpReq) (*model.UserAuthResponse, error) {
	return &model.UserAuthResponse{}, nil
}

func (c *UserService) LogOut(refresh string) error {
	return nil
}
