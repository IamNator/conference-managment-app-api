package service_test

import (
	"conference/model"
	"conference/service"
	"conference/storage/mock"
	"conference/testdata"
	"fmt"
	"github.com/golang/mock/gomock"
	"strconv"
	"testing"
)

func TestConferenceService_CreateConference(t *testing.T) {

	type tData struct {
		Req   model.CreateConferenceReq `json:"req"`
		Error string                    `json:"error"`
	}

	Reqs := make([]tData, 0)
	if er := testdata.Load("../testdata/model/create_conference_req.json", &Reqs); er != nil {
		t.Error(er.Error())
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for i, v := range Reqs {
		mm := model.Conference{
			Title:       v.Req.Title,
			Description: v.Req.Description,
			StartDate:   v.Req.StartDate,
			EndDate:     v.Req.EndDate,
		}

		mockConfRepo := mock.NewMockIConferenceRepository(ctrl)
		mockConfRepo.EXPECT().CreateConference(mm).Return(&model.Conference{
			General: model.General{
				ID: uint(i),
			},
			Title:       v.Req.Title,
			Description: v.Req.Description,
			StartDate:   v.Req.StartDate,
			EndDate:     v.Req.EndDate,
		}, nil)
		mockConfRepo.EXPECT().SaveEditHistory(model.EditHistory{
			ConferenceID:     uint(i),
			PropertyAffected: "conference",
			Action:           "created conference",
			By:               "username",
		}).Return(nil)

		srv := service.ConferenceService{
			ConferenceRepo: mockConfRepo,
		}

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			_, er := srv.CreateConference("username", v.Req)
			if er != nil {
				t.Error(er.Error())
			}
		})
	}

}

func TestConferenceService_UpdateConference(t *testing.T) {

	type tData struct {
		Req   model.UpdateConferenceReq `json:"req"`
		Error string                    `json:"error"`
	}

	Reqs := make([]tData, 0)
	if er := testdata.Load("../testdata/model/update_conference_req.json", &Reqs); er != nil {
		t.Error(er.Error())
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for i, v := range Reqs {

		mm := model.Conference{
			General: model.General{
				ID: v.Req.ConferenceID,
			},
			Title:       v.Req.Title,
			Description: v.Req.Description,
			StartDate:   v.Req.StartDate,
			EndDate:     v.Req.EndDate,
		}

		mockConfRepo := mock.NewMockIConferenceRepository(ctrl)
		mockConfRepo.EXPECT().UpdateConference(mm).Return(&mm, nil)
		mockConfRepo.EXPECT().SaveEditHistory(model.EditHistory{
			ConferenceID:     v.Req.ConferenceID,
			PropertyAffected: "conference",
			Action:           "updated conference",
			By:               "username",
		}).Return(nil)

		srv := service.ConferenceService{
			ConferenceRepo: mockConfRepo,
		}

		t.Run(fmt.Sprintf("test: %d", i), func(t *testing.T) {
			_, er := srv.UpdateConference("username", v.Req)
			if er != nil {
				t.Error(er.Error())
			}
		})
	}

}
