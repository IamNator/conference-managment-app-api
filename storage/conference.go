package storage

import (
	"conference/model"
	"errors"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"os"
)

//go:generate mockgen -source conference.go -destination ./mock/conference.go -package mock IConferenceRepository
type IConferenceRepository interface {
	CreateConference(conference model.Conference) (*model.Conference, error)
	UpdateConference(conference model.Conference) (*model.Conference, error)
	GetConference(conferenceId uint, page, pageSize int) ([]model.Conference, error)
	GetAllConference(page, pageSize int) ([]model.Conference, error)

	CreateTalk(talk model.Talk) (*model.Talk, error)
	UpdateTalk(talk model.Talk) (*model.Talk, error)
	GetTalks(conferenceId uint, page, pageSize int) ([]model.Talk, error)

	CreateSpeaker(speaker model.Speaker) (*model.Speaker, error)
	GetSpeakers(TalkId uint, page, pageSize int) ([]model.Speaker, error)
	DeleteSpeaker(speakerId, talkId uint) error

	CreateParticipant(participant model.Participant) (*model.Participant, error)
	GetParticipants(TalkId uint, page, pageSize int) ([]model.Participant, error)
	DeleteParticipant(participantId, talkId uint) error

	SaveEditHistory(history model.EditHistory) error
	GetEditHistory(conferenceID uint, page, pageSize int) ([]model.EditHistory, error)

	WithTx(db *gorm.DB) IConferenceRepository
}

type ConferenceRepository struct {
	storage *Storage
	logger  *zerolog.Logger
}

func NewConferenceRepository(s *Storage) IConferenceRepository {
	_logger := zerolog.New(os.Stdout).With().Str("app", "conf_mgmt_sys").
		Str("module", "conference storage").Logger()
	return &ConferenceRepository{
		storage: s,
		logger:  &_logger,
	}
}

func (c *ConferenceRepository) WithTx(db *gorm.DB) IConferenceRepository {
	return &ConferenceRepository{
		storage: &Storage{
			db: db,
		},
	}
}

//CONFERENCES

func (c *ConferenceRepository) CreateConference(conference model.Conference) (*model.Conference, error) {
	return &conference, c.storage.db.Create(&conference).Error
}

func (c *ConferenceRepository) UpdateConference(conference model.Conference) (*model.Conference, error) {
	er := c.storage.db.Where("id = ?", conference.ID).Updates(&conference).First(&conference).Error
	if er != nil {
		return nil, er
	}
	return &conference, nil
}

func (c *ConferenceRepository) GetConference(conferenceId uint, page, pageSize int) ([]model.Conference, error) {
	var conferences []model.Conference
	return conferences, c.storage.db.Scopes(model.Paginate(page, pageSize)).Where("conference_id = ?", conferenceId).First(&conferences).Error
}

func (c *ConferenceRepository) GetAllConference(page, pageSize int) ([]model.Conference, error) {
	var conferences []model.Conference
	return conferences, c.storage.db.Scopes(model.Paginate(page, pageSize)).Find(&conferences).Error
}

//TALKS

func (c *ConferenceRepository) CreateTalk(talk model.Talk) (*model.Talk, error) {
	return &talk, c.storage.db.Create(&talk).Error
}

func (c *ConferenceRepository) GetTalks(conferenceId uint, page, pageSize int) ([]model.Talk, error) {
	var talks []model.Talk
	return talks, c.storage.db.Scopes(model.Paginate(page, pageSize)).Where("conference_id = ? ", conferenceId).Find(&talks).Error
}

func (c *ConferenceRepository) UpdateTalk(talk model.Talk) (*model.Talk, error) {
	er := c.storage.db.Where("id = ? AND conference_id = ? ", talk.General.ID, talk.ConferenceID).Updates(&talk).First(&talk).Error
	if er != nil {
		return nil, er
	}
	return &talk, nil
}

//SPEAKERS

func (c *ConferenceRepository) CreateSpeaker(speaker model.Speaker) (*model.Speaker, error) {
	var _speaker model.Speaker
	er := c.storage.db.Where("email = ? AND talk_id = ?", speaker.Email, speaker.TalkID).First(&_speaker).Error
	if er == nil || !errors.Is(er, gorm.ErrRecordNotFound) {
		return nil, errors.New("email already exist")
	}

	return &speaker, c.storage.db.Create(&speaker).Error
}

func (c *ConferenceRepository) GetSpeakers(TalkId uint, page, pageSize int) ([]model.Speaker, error) {
	var speakers []model.Speaker
	return speakers, c.storage.db.Where("talk_id = ?", TalkId).Scopes(model.Paginate(page, pageSize)).Find(&speakers).Error
}

func (c *ConferenceRepository) DeleteSpeaker(speakerId, talkId uint) error {
	return c.storage.db.Delete(&model.Speaker{}, "id = ? AND talk_id = ?", speakerId, talkId).Error
}

// PARTICIPANTS

func (c *ConferenceRepository) CreateParticipant(participant model.Participant) (*model.Participant, error) {
	var _participant model.Participant
	er := c.storage.db.Where("email = ? AND talk_id = ?", participant.Email, participant.TalkID).First(&_participant).Error
	if er == nil || !errors.Is(er, gorm.ErrRecordNotFound) {
		return nil, errors.New("email already exist")
	}

	return &participant, c.storage.db.Create(&participant).Error
}

func (c *ConferenceRepository) GetParticipants(TalkId uint, page, pageSize int) ([]model.Participant, error) {
	var participants []model.Participant
	return participants, c.storage.db.Scopes(model.Paginate(page, pageSize)).Where("talk_id = ?", TalkId).Find(&participants).Error
}

func (c *ConferenceRepository) DeleteParticipant(participantId, talkId uint) error {
	return c.storage.db.Delete(&model.Participant{}, "id = ? AND talk_id = ?", participantId, talkId).Error
}

func (c *ConferenceRepository) SaveEditHistory(history model.EditHistory) error {
	return c.storage.db.Create(&history).Error
}

func (c *ConferenceRepository) GetEditHistory(conferenceID uint, page, pageSize int) ([]model.EditHistory, error) {
	var history []model.EditHistory
	return history, c.storage.db.Scopes(model.Paginate(page, pageSize)).Find(&history, "conference_id = ?", conferenceID).Error
}
