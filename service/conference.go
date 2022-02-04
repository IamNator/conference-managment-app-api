package service

import (
	"conference/model"
	"conference/storage"
)

//go:generate mockgen -source conference.go -destination ./mock/conference_service.go -package mock IConferenceService
type IConferenceService interface {
	CreateConference(conference model.CreateConferenceReq) (*model.Conference, error)
	UpdateConference(conference model.UpdateConferenceReq) (*model.Conference, error)
	CreateTalk(req model.CreateTalkReq) (*model.Talk, error)
	UpdateTalk(req model.UpdateTalkReq) (*model.Talk, error)

	AddSpeaker(userId uint, req model.AddSpeakerReq) (*model.Speaker, error)
	GetSpeakers(talkId uint, page int, pageSize int) ([]model.Speaker, error)
	RemoveSpeaker(speakerId uint, talkId uint) error

	AddParticipant(userId uint, req model.AddParticipantReq) (*model.Participant, error)
	GetParticipants(talkId uint, page int, pageSize int) ([]model.Participant, error)
	RemoveParticipant(participantId uint, talkId uint) error

	GetTalks(conferenceId uint, page int, pageSize int) ([]model.Talk, error)
	GetConferences(page int, pageSize int) ([]model.Conference, error)
	GetEditHistory(conferenceId uint, page int, pageSize int) ([]model.EditHistory, error)
}

type ConferenceService struct {
	ConferenceRepo storage.IConferenceRepository
}

func NewConferenceService(s *storage.Storage) IConferenceService {
	return &ConferenceService{
		ConferenceRepo: storage.NewConferenceRepository(s),
	}
}

//To create/edit a conference

func (c *ConferenceService) CreateConference(conference model.CreateConferenceReq) (*model.Conference, error) {
	return c.ConferenceRepo.CreateConference(model.Conference{
		Title:       conference.Title,
		Description: conference.Description,
		StartDate:   conference.StartDate,
		EndDate:     conference.EndDate,
	})
}

func (c *ConferenceService) UpdateConference(conference model.UpdateConferenceReq) (*model.Conference, error) {
	return c.ConferenceRepo.UpdateConference(model.Conference{
		General: model.General{
			ID: conference.ConferenceID,
		},
		Title:       conference.Title,
		Description: conference.Description,
		StartDate:   conference.StartDate,
		EndDate:     conference.EndDate,
	})
}

//○ To add/edit a talk

func (c *ConferenceService) CreateTalk(req model.CreateTalkReq) (*model.Talk, error) {
	return c.ConferenceRepo.CreateTalk(model.Talk{
		ConferenceID: req.ConferenceID,
		Title:        req.Title,
		Description:  req.Description,
		Duration:     req.Duration,
		DateTime:     req.DateTime,
	})
}

func (c *ConferenceService) UpdateTalk(req model.UpdateTalkReq) (*model.Talk, error) {
	return c.ConferenceRepo.UpdateTalk(model.Talk{
		General: model.General{
			ID: req.TalkID,
		},
		ConferenceID: req.ConferenceID,
		Title:        req.Title,
		Description:  req.Description,
		Duration:     req.Duration,
		DateTime:     req.DateTime,
	})
}

//○ To add/remove speaker/participant from a talk.

func (c *ConferenceService) AddSpeaker(userId uint, req model.AddSpeakerReq) (*model.Speaker, error) {
	return c.ConferenceRepo.CreateSpeaker(model.Speaker{
		TalkID:   req.TalkID,
		Username: req.Username,
		Email:    req.Email,
		UserID:   &userId,
	})
}

func (c *ConferenceService) GetSpeakers(talkId uint, page int, pageSize int) ([]model.Speaker, error) {
	return c.ConferenceRepo.GetSpeakers(talkId, page, pageSize)
}

func (c *ConferenceService) RemoveSpeaker(speakerId uint, talkId uint) error {
	return c.ConferenceRepo.DeleteSpeaker(speakerId, talkId)
}

func (c *ConferenceService) AddParticipant(userId uint, req model.AddParticipantReq) (*model.Participant, error) {
	return c.ConferenceRepo.CreateParticipant(model.Participant{
		TalkID:   req.TalkID,
		Username: req.Username,
		Email:    req.Email,
		UserID:   &userId,
	})
}

func (c *ConferenceService) GetParticipants(talkId uint, page int, pageSize int) ([]model.Participant, error) {
	return c.ConferenceRepo.GetParticipants(talkId, page, pageSize)
}

func (c *ConferenceService) RemoveParticipant(participantId uint, talkId uint) error {
	return c.ConferenceRepo.DeleteParticipant(participantId, talkId)
}

//○ To list talks in a conference

func (c *ConferenceService) GetTalks(conferenceId uint, page int, pageSize int) ([]model.Talk, error) {
	return c.ConferenceRepo.GetTalks(conferenceId, page, pageSize)
}

//○ To list conferences

func (c *ConferenceService) GetConferences(page int, pageSize int) ([]model.Conference, error) {
	return c.ConferenceRepo.GetAllConference(page, pageSize)
}

//○ To view the history of edits

func (c *ConferenceService) GetEditHistory(conferenceId uint, page int, pageSize int) ([]model.EditHistory, error) {
	return c.ConferenceRepo.GetEditHistory(conferenceId, page, pageSize)
}
