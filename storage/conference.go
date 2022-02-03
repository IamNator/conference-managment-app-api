package storage

import (
	"github.com/iamnator/conference_mgmt_sys/model"
	"gorm.io/gorm"
)

//go:generate mockgen -source conference.go -destination ./mock/conference.go -package mock IConferenceRepository
type IConferenceRepository interface {
	SaveConference(conference model.Conference) (*model.Conference, error)
	UpdateConference(conference model.Conference) (*model.Conference, error)
	GetConference(conferenceId uint, page, pageSize int) ([]model.Conference, error)
	GetAllConference(page, pageSize int) ([]model.Conference, error)

	SaveTalk(talk model.Talk) (*model.Talk, error)
	UpdateTalk(talk model.Talk) (*model.Talk, error)
	GetTalks(conferenceId uint, page, pageSize int) ([]model.Talk, error)

	SaveSpeaker(speaker model.Speaker) (*model.Speaker, error)
	GetSpeakers(TalkId uint, page, pageSize int) ([]model.Talk, error)
	DeleteSpeaker(speakerId, talkId uint) error

	SaveParticipant(participant model.Participant) (*model.Participant, error)
	GetParticipants(TalkId uint, page, pageSize int) ([]model.Talk, error)
	DeleteParticipant(participantId, talkId uint) error

	SaveEditHistory(history model.EditHistory) error
	GetEditHistory(conferenceID uint, page, pageSize int) ([]model.EditHistory, error)
}

type ConferenceRepository struct {
	storage *Storage
}

func (c *ConferenceRepository) WithTx(db *gorm.DB) IConferenceRepository {
	return &ConferenceRepository{
		storage: &Storage{
			db: db,
		},
	}
}

//CONFERENCES

func (c *ConferenceRepository) SaveConference(conference model.Conference) (*model.Conference, error) {
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

func (c *ConferenceRepository) SaveTalk(talk model.Talk) (*model.Talk, error) {
	return &model.Talk{}, nil
}

func (c *ConferenceRepository) GetTalks(conferenceId uint, page, pageSize int) ([]model.Talk, error) {
	return []model.Talk{}, nil
}

func (c *ConferenceRepository) UpdateTalk(talk model.Talk) (*model.Talk, error) {
	return &model.Talk{}, nil
}

//SPEAKERS

func (c *ConferenceRepository) SaveSpeaker(speaker model.Speaker) (*model.Speaker, error) {
	return &model.Speaker{}, nil
}

func (c *ConferenceRepository) GetSpeakers(TalkId uint, page, pageSize int) ([]model.Talk, error) {
	return []model.Talk{}, nil
}

func (c *ConferenceRepository) DeleteSpeaker(speakerId, talkId uint) error {
	return nil
}

// PARTICIPANTS

func (c *ConferenceRepository) SaveParticipant(participant model.Participant) (*model.Participant, error) {
	return &model.Participant{}, nil
}

func (c *ConferenceRepository) GetParticipants(TalkId uint, page, pageSize int) ([]model.Talk, error) {
	return []model.Talk{}, nil
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
