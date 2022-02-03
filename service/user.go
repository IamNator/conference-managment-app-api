package service

import "github.com/iamnator/conference_mgmt_sys/model"

func (c *Service) Login(req model.UserLoginReq) (*model.UserAuthResponse, error) {
	return &model.UserAuthResponse{}, nil
}

func (c *Service) RegisterUser(req model.UserSignUpReq) (*model.UserAuthResponse, error) {
	return &model.UserAuthResponse{}, nil
}

func (c *Service) LogOut(refresh string) error {
	return nil
}
