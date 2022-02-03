package service

import (
	"conference/model"
	"conference/storage"
)

//go:generate mockgen -source conference.go -destination ./mock/conference_service.go -package mock IConferenceService
type IConferenceService interface {
	CreateConference(conference model.CreateConferenceReq) (*model.Conference, error)
}

type ConferenceService struct {
	conferenceRepo storage.IConferenceRepository
}

func NewConferenceService(s *storage.Storage) IConferenceService {
	return &ConferenceService{
		conferenceRepo: storage.NewConferenceRepository(s),
	}
}

func (c *ConferenceService) CreateConference(conference model.CreateConferenceReq) (*model.Conference, error) {
	return c.conferenceRepo.SaveConference(model.Conference{
		Title:       conference.Title,
		Description: conference.Description,
		StartDate:   conference.StartDate,
		EndDate:     conference.EndDate,
	})
}
