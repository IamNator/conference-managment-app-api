package service

import (
	"conference/model"
	"conference/pkg/middleware"
	"conference/storage"
	"errors"
	"log"
	"time"
)

//go:generate mockgen -source user.go -destination ./mock/user_service.go -package mock IUserService
type IUserService interface {
	Login(req model.UserLoginReq) (*model.UserAuthResponse, error)
	RegisterUser(req model.UserSignUpReq) (*model.UserAuthResponse, error)
	LogOut(refresh, access string) error
}

type UserService struct {
	UserRepo storage.IUserRepository
	MidWare  middleware.IMiddleware
}

func NewUserService(s *storage.Storage) IUserService {
	return &UserService{
		UserRepo: storage.NewUserRepository(s),
		MidWare:  middleware.NewMiddleWare(),
	}
}

func (c *UserService) Login(req model.UserLoginReq) (*model.UserAuthResponse, error) {

	user, er := c.UserRepo.GetUserByEmail(req.Email)
	if er != nil {
		return nil, er
	}

	if !user.Password.Compare(req.Password.String()) {
		return nil, errors.New("incorrect login details")
	}

	if er := c.UserRepo.UpdateLastLoggedIn(user.Email, time.Now()); er != nil {
		log.Println(er.Error()) //just log the error
	}

	auth, er := c.MidWare.GenerateToken(*user)
	if er != nil {
		return nil, er
	}

	return auth, nil
}

func (c *UserService) RegisterUser(req model.UserSignUpReq) (*model.UserAuthResponse, error) {

	user, err := c.UserRepo.GetUserByEmail(req.Email)
	if user != nil && err == nil {
		return nil, errors.New("email already exists")
	}

	user, er := c.UserRepo.CreateUser(model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	if er != nil {
		return nil, er
	}

	///add auth
	auth, er := c.MidWare.GenerateToken(*user)
	if er != nil {
		return nil, er
	}

	return auth, nil
}

func (c *UserService) LogOut(refreshToken, accessToken string) error {
	return nil
}
