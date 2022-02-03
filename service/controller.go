package service

import (
	"github.com/iamnator/conference_mgmt_sys/model"
	"github.com/iamnator/conference_mgmt_sys/storage"
)

type IService interface {
	Login(req model.UserLoginReq) (*model.UserAuthResponse, error)
	RegisterUser(req model.UserSignUpReq) (*model.UserAuthResponse, error)
	LogOut(refresh string) error
}

type Service struct {
	userRepo       storage.IUserRepository
	conferenceRepo storage.IConferenceRepository
}
