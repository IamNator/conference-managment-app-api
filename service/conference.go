package service

import (
	"conference/model"
	"conference/storage"
	"github.com/rs/zerolog"
	"os"
)

//go:generate mockgen -source conference.go -destination ./mock/conference_service.go -package mock IConferenceService
type IConferenceService interface {
	CreateConference(username string, conference model.CreateConferenceReq) (*model.Conference, error)
	UpdateConference(username string, conference model.UpdateConferenceReq) (*model.Conference, error)
	CreateTalk(username string, req model.CreateTalkReq) (*model.Talk, error)
	UpdateTalk(username string, req model.UpdateTalkReq) (*model.Talk, error)

	AddSpeaker(username string, req model.AddSpeakerReq) (*model.Speaker, error)
	GetSpeakers(talkId uint, page int, pageSize int) ([]model.Speaker, error)
	RemoveSpeaker(username string, conferenceId, speakerId uint, talkId uint) error

	AddParticipant(username string, req model.AddParticipantReq) (*model.Participant, error)
	GetParticipants(talkId uint, page int, pageSize int) ([]model.Participant, error)
	RemoveParticipant(username string, conferenceId, participantId uint, talkId uint) error

	GetTalks(conferenceId uint, page int, pageSize int) ([]model.Talk, error)
	GetConferences(page int, pageSize int) ([]model.Conference, error)
	GetEditHistory(conferenceId uint, page int, pageSize int) ([]model.EditHistory, error)
}

type ConferenceService struct {
	ConferenceRepo storage.IConferenceRepository
	logger         *zerolog.Logger
}

func NewConferenceService(s *storage.Storage) IConferenceService {
	_logger := zerolog.New(os.Stdout).With().Str("app", "conf_mgmt_sys").
		Str("module", "conference service").Logger()
	return &ConferenceService{
		ConferenceRepo: storage.NewConferenceRepository(s),
		logger:         &_logger,
	}
}

//To create/edit a conference

func (c *ConferenceService) CreateConference(username string, conference model.CreateConferenceReq) (*model.Conference, error) {
	resp, er := c.ConferenceRepo.CreateConference(model.Conference{
		Title:       conference.Title,
		Description: conference.Description,
		StartDate:   conference.StartDate,
		EndDate:     conference.EndDate,
	})
	if er != nil {
		return nil, er
	}

	er = c.ConferenceRepo.SaveEditHistory(model.EditHistory{
		ConferenceID:     resp.ID,
		PropertyAffected: "conference",
		Action:           "create",
		By:               username,
	})
	if er != nil {
		c.logger.Error().Err(er).Msg("unable to save edit history")
	}

	return resp, nil
}

func (c *ConferenceService) UpdateConference(username string, conference model.UpdateConferenceReq) (*model.Conference, error) {
	resp, er := c.ConferenceRepo.UpdateConference(model.Conference{
		General: model.General{
			ID: conference.ConferenceID,
		},
		Title:       conference.Title,
		Description: conference.Description,
		StartDate:   conference.StartDate,
		EndDate:     conference.EndDate,
	})
	if er != nil {
		return nil, er
	}

	er = c.ConferenceRepo.SaveEditHistory(model.EditHistory{
		ConferenceID:     resp.ID,
		PropertyAffected: "conference",
		Action:           "updated conference",
		By:               username,
	})
	if er != nil {
		c.logger.Error().Err(er).Msg("unable to save edit history")
	}

	return resp, nil
}

//○ To add/edit a talk

func (c *ConferenceService) CreateTalk(username string, req model.CreateTalkReq) (*model.Talk, error) {
	return c.ConferenceRepo.CreateTalk(model.Talk{
		ConferenceID: req.ConferenceID,
		Title:        req.Title,
		Description:  req.Description,
		Duration:     req.Duration,
		DateTime:     req.DateTime,
	})
}

func (c *ConferenceService) UpdateTalk(username string, req model.UpdateTalkReq) (*model.Talk, error) {
	resp, er := c.ConferenceRepo.UpdateTalk(model.Talk{
		General: model.General{
			ID: req.TalkID,
		},
		ConferenceID: req.ConferenceID,
		Title:        req.Title,
		Description:  req.Description,
		Duration:     req.Duration,
		DateTime:     req.DateTime,
	})
	if er != nil {
		return nil, er
	}

	er = c.ConferenceRepo.SaveEditHistory(model.EditHistory{
		ConferenceID:     resp.ID,
		PropertyAffected: "talk",
		Action:           "updated a talk",
		By:               username,
	})
	if er != nil {
		c.logger.Error().Err(er).Msg("unable to save edit history")
	}

	return resp, nil
}

//○ To add/remove speaker/participant from a talk.

func (c *ConferenceService) AddSpeaker(username string, req model.AddSpeakerReq) (*model.Speaker, error) {
	resp, er := c.ConferenceRepo.CreateSpeaker(model.Speaker{
		TalkID:   req.TalkID,
		Username: req.Username,
		Email:    req.Email,
	})
	if er != nil {
		return nil, er
	}

	er = c.ConferenceRepo.SaveEditHistory(model.EditHistory{
		ConferenceID:     req.ConferenceId,
		PropertyAffected: "speaker",
		Action:           "added a speaker",
		By:               username,
	})
	if er != nil {
		c.logger.Error().Err(er).Msg("unable to save edit history")
	}

	return resp, nil
}

func (c *ConferenceService) GetSpeakers(talkId uint, page int, pageSize int) ([]model.Speaker, error) {
	return c.ConferenceRepo.GetSpeakers(talkId, page, pageSize)
}

func (c *ConferenceService) RemoveSpeaker(username string, conferenceId, speakerId uint, talkId uint) error {
	er := c.ConferenceRepo.DeleteSpeaker(speakerId, talkId)
	if er != nil {
		return er
	}

	er = c.ConferenceRepo.SaveEditHistory(model.EditHistory{
		ConferenceID:     conferenceId,
		PropertyAffected: "speaker",
		Action:           "removed a speaker",
		By:               username,
	})
	if er != nil {
		c.logger.Error().Err(er).Msg("unable to save edit history")
	}

	return nil
}

func (c *ConferenceService) AddParticipant(username string, req model.AddParticipantReq) (*model.Participant, error) {
	resp, er := c.ConferenceRepo.CreateParticipant(model.Participant{
		TalkID:   req.TalkID,
		Username: req.Username,
		Email:    req.Email,
	})

	if er != nil {
		return nil, er
	}

	er = c.ConferenceRepo.SaveEditHistory(model.EditHistory{
		ConferenceID:     req.ConferenceId,
		PropertyAffected: "participant",
		Action:           "added a participant",
		By:               username,
	})
	if er != nil {
		c.logger.Error().Err(er).Msg("unable to save edit history")
	}

	return resp, nil
}

func (c *ConferenceService) GetParticipants(talkId uint, page int, pageSize int) ([]model.Participant, error) {
	return c.ConferenceRepo.GetParticipants(talkId, page, pageSize)
}

func (c *ConferenceService) RemoveParticipant(username string, conferenceId, participantId uint, talkId uint) error {

	er := c.ConferenceRepo.DeleteParticipant(participantId, talkId)
	if er != nil {
		return er
	}

	er = c.ConferenceRepo.SaveEditHistory(model.EditHistory{
		ConferenceID:     conferenceId,
		PropertyAffected: "participant",
		Action:           "removed a participant",
		By:               username,
	})
	if er != nil {
		c.logger.Error().Err(er).Msg("unable to save edit history")
	}

	return nil
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
