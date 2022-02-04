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

//To create/edit a conference
//○ To add/edit a talk
//○ To add/remove speaker/participant from a talk.
//○ To list talks in a conference
//○ To list conferences
//○ To view the history of edits

func (c *ConferenceService) CreateConference(conference model.CreateConferenceReq) (*model.Conference, error) {
	return c.conferenceRepo.CreateConference(model.Conference{
		Title:       conference.Title,
		Description: conference.Description,
		StartDate:   conference.StartDate,
		EndDate:     conference.EndDate,
	})
}

func (c *ConferenceService) UpdateConference(conference model.CreateConferenceReq) (*model.Conference, error) {
	return c.conferenceRepo.CreateConference(model.Conference{
		Title:       conference.Title,
		Description: conference.Description,
		StartDate:   conference.StartDate,
		EndDate:     conference.EndDate,
	})
}
