package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"time"
)

type (
	Conference struct {
		General
		Title       string    `json:"title"`
		Description string    `json:"description"`
		StartDate   time.Time `json:"Start_date"`
		EndDate     time.Time `json:"end_date"`
		Talks       []Talk    `json:"talk"`
	}

	Talk struct {
		General
		ConferenceID uint          `json:"conference_id"`
		Title        string        `json:"title"`
		Description  string        `json:"description"`
		Duration     time.Duration `json:"duration"`
		DateTime     time.Time     `json:"date_time"`
		Speakers     []Speaker     `json:"speakers"`
		Participants []Participant `json:"participants"`
	}

	Speaker struct {
		TalkID   uint   `json:"talk_id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		UserID   *uint  `json:"user_id"`
	}

	Participant struct {
		TalkID   uint   `json:"talk_id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		UserID   *uint  `json:"user_id"`
	}

	EditHistory struct {
		ConferenceID uint   `json:"conference_id"`
		PropertyID   uint   `json:"property_id"`
		Property     string `json:"property"` //
		Field        string `json:"field"`
		PreValue     string `json:"pre_value"`
		CurrentValue string `json:"current_value"`
		By           uint   `json:"by"`
		General
	}
)

type (
	CreateConferenceReq struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		StartDate   time.Time `json:"start_date"`
		EndDate     time.Time `json:"end_date"`
	}

	CreateTalkReq struct {
		ConferenceID uint          `json:"conference_id"`
		Title        string        `json:"title"`
		Description  string        `json:"description"`
		Duration     time.Duration `json:"duration"`
		DateTime     time.Time     `json:"date_time"`
	}

	AddSpeakerReq struct {
		TalkID   uint   `json:"talk_id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	AddParticipantReq struct {
		TalkID   uint   `json:"talk_id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	//edits to a conference
	EditHistoryResp struct {
		ConferenceID uint      `json:"conference_id"`
		Property     string    `json:"property"` //
		PropertyID   uint      `json:"property_id"`
		Field        string    `json:"field"`
		PreValue     string    `json:"pre_value"`
		CurrentValue string    `json:"current_value"`
		By           uint      `json:"by"`
		CreatedAt    time.Time `json:"created_at"`
	}
)

func (c CreateConferenceReq) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Required, validation.Length(1, 50)),
		validation.Field(&c.Description, validation.Required, validation.Length(5, 1000)),
		validation.Field(&c.StartDate, validation.Required),
		validation.Field(&c.EndDate, validation.Required),
	)
}

func (c CreateTalkReq) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Required, validation.Length(1, 50)),
		validation.Field(&c.Description, validation.Required, validation.Length(5, 1000)),
		validation.Field(&c.DateTime, validation.Required),
		validation.Field(&c.ConferenceID, validation.Required),
		validation.Field(&c.Duration, validation.Required),
	)
}

func (a AddSpeakerReq) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Username, validation.Required, validation.Length(1, 50)),
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.TalkID, validation.Required),
	)
}

func (a AddParticipantReq) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Username, validation.Required, validation.Length(1, 50)),
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.TalkID, validation.Required),
	)
}
