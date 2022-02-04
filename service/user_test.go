package service_test

import (
	"conference/model"
	"conference/service"
	"conference/storage/mock"
	"conference/testdata"
	"github.com/golang/mock/gomock"
	"strconv"
	"testing"
	"time"
)

func TestUserService_RegisterUser(t *testing.T) {
	type tData struct {
		Req   model.UserSignUpReq `json:"req"`
		Error string              `json:"error"`
	}

	Reqs := make([]tData, 0)
	if er := testdata.Load("../testdata/model/signup.json", &Reqs); er != nil {
		t.Error(er.Error())
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var now time.Time
	for i, v := range Reqs {
		now = time.Now()
		mm := model.User{
			Username:     v.Req.Username,
			Email:        v.Req.Email,
			Password:     v.Req.Password,
			LastLoggedIn: &now,
			General:      model.General{CreatedAt: now, UpdatedAt: now},
		}

		mockUserRepo := mock.NewMockIUserRepository(ctrl)
		mockUserRepo.EXPECT().CreateUser(model.User{
			Username: mm.Username,
			Email:    mm.Email,
			Password: mm.Password,
		}).Return(&mm, nil)
		mockUserRepo.EXPECT().GetUserByEmail(mm.Email).Return(nil, nil)

		srv := service.UserService{
			UserRepo: mockUserRepo,
		}

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			_, er := srv.RegisterUser(v.Req)
			if er != nil {
				t.Error(er.Error())
			}
		})
	}
}

func TestUserService_Login(t *testing.T) {
	t.Log("successful")
}

func TestUserService_LogOut(t *testing.T) {
	t.Log("successful")
}
