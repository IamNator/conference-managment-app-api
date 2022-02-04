package service

import (
	"conference/model"
	"conference/storage"
	"errors"
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
	user, er := c.userRepo.GetUserByEmail(req.Email)
	if er != nil {
		return nil, er
	}

	if !user.Password.Compare(req.Password.String()) {
		return nil, errors.New("incorrect login details")
	}

	//TODO
	//add auth

	return &model.UserAuthResponse{
		User: *user,
	}, nil
}

func (c *UserService) RegisterUser(req model.UserSignUpReq) (*model.UserAuthResponse, error) {

	user, _ := c.userRepo.GetUserByEmail(req.Email)
	if user != nil {
		return nil, errors.New("email already exists")
	}

	user, er := c.userRepo.CreateUser(model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	if er != nil {
		return nil, er
	}

	//TODO
	//add auth

	return &model.UserAuthResponse{
		User: *user,
	}, nil
}

func (c *UserService) LogOut(refreshToken string) error {
	return nil
}
