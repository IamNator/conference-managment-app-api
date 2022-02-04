package storage

import (
	"conference/model"
	"gorm.io/gorm"
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
}

func NewConferenceRepository(s *Storage) IConferenceRepository {
	return &ConferenceRepository{
		storage: s,
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
	return &model.Conference{}, nil
}

func (c *ConferenceRepository) UpdateConference(conference model.Conference) (*model.Conference, error) {
	return &model.Conference{}, nil
}

func (c *ConferenceRepository) GetConference(conferenceId uint, page, pageSize int) ([]model.Conference, error) {
	return []model.Conference{}, nil
}

func (c *ConferenceRepository) GetAllConference(page, pageSize int) ([]model.Conference, error) {
	return []model.Conference{}, nil
}

//TALKS

func (c *ConferenceRepository) CreateTalk(talk model.Talk) (*model.Talk, error) {
	return &model.Talk{}, nil
}

func (c *ConferenceRepository) GetTalks(conferenceId uint, page, pageSize int) ([]model.Talk, error) {
	return []model.Talk{}, nil
}

func (c *ConferenceRepository) UpdateTalk(talk model.Talk) (*model.Talk, error) {
	return &model.Talk{}, nil
}

//SPEAKERS

func (c *ConferenceRepository) CreateSpeaker(speaker model.Speaker) (*model.Speaker, error) {
	return &model.Speaker{}, nil
}

func (c *ConferenceRepository) GetSpeakers(TalkId uint, page, pageSize int) ([]model.Speaker, error) {
	return []model.Speaker{}, nil
}

func (c *ConferenceRepository) DeleteSpeaker(speakerId, talkId uint) error {
	return nil
}

// PARTICIPANTS

func (c *ConferenceRepository) CreateParticipant(participant model.Participant) (*model.Participant, error) {
	return &model.Participant{}, nil
}

func (c *ConferenceRepository) GetParticipants(TalkId uint, page, pageSize int) ([]model.Participant, error) {
	return []model.Participant{}, nil
}

func (c *ConferenceRepository) DeleteParticipant(participantId, talkId uint) error {
	return nil
}

func (c *ConferenceRepository) SaveEditHistory(history model.EditHistory) error {
	return nil
}

func (c *ConferenceRepository) GetEditHistory(conferenceID uint, page, pageSize int) ([]model.EditHistory, error) {
	return []model.EditHistory{}, nil
}
