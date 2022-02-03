package service

import "conference/model"

func (c *Service) Login(req model.UserLoginReq) (*model.UserAuthResponse, error) {
	user, _ := c.userRepo.GetUserByEmail(req.Email)

	return &model.UserAuthResponse{
		User: *user,
	}, nil
}

func (c *Service) RegisterUser(req model.UserSignUpReq) (*model.UserAuthResponse, error) {
	return &model.UserAuthResponse{}, nil
}

func (c *Service) LogOut(refresh string) error {
	return nil
}
