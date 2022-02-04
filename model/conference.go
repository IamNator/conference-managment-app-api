package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gorm.io/gorm"
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
		General
		TalkID   uint   `json:"talk_id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		UserID   *uint  `json:"user_id"`
	}

	Participant struct {
		General
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

func (c Conference) TableName(tx *gorm.DB) string {
	return "conference"
}

func (t Talk) TableName(tx *gorm.DB) string {
	return "talk"
}

func (s Speaker) TableName(tx *gorm.DB) string {
	return "speaker"
}

func (s Participant) TableName(tx *gorm.DB) string {
	return "participant"
}

func (e EditHistory) TableName(tx *gorm.DB) string {
	return "edit_history"
}

//

func (c Conference) CreateTable(tx *gorm.DB) error {
	return tx.AutoMigrate(c)
}

func (t Talk) CreateTable(tx *gorm.DB) error {
	return tx.AutoMigrate(t)
}

func (s Speaker) CreateTable(tx *gorm.DB) error {
	return tx.AutoMigrate(s)
}

func (s Participant) CreateTable(tx *gorm.DB) error {
	return tx.AutoMigrate(s)
}

func (e EditHistory) CreateTable(tx *gorm.DB) error {
	return tx.AutoMigrate(e)
}

type (
	CreateConferenceReq struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		StartDate   time.Time `json:"start_date"`
		EndDate     time.Time `json:"end_date"`
	}

	UpdateConferenceReq struct {
		ConferenceID uint      `json:"conference_id"`
		Title        string    `json:"title"`
		Description  string    `json:"description"`
		StartDate    time.Time `json:"start_date"`
		EndDate      time.Time `json:"end_date"`
	}

	GetSpeakerReq struct {
		TalkId   uint `json:"talk_id" form:"talk_id"`
		Page     int  `json:"page" form:"page"`
		PageSize int  `json:"page_size" form:"page_size"`
	}

	GetParticipantReq struct {
		TalkId   uint `json:"talk_id" form:"talk_id"`
		Page     int  `json:"page" form:"page"`
		PageSize int  `json:"page_size" form:"page_size"`
	}

	CreateTalkReq struct {
		ConferenceID uint          `json:"conference_id"`
		Title        string        `json:"title"`
		Description  string        `json:"description"`
		Duration     time.Duration `json:"duration"`
		DateTime     time.Time     `json:"date_time"`
	}

	UpdateTalkReq struct {
		TalkID       uint          `json:"talk_id"`
		ConferenceID uint          `json:"conference_id"`
		Title        string        `json:"title"`
		Description  string        `json:"description"`
		Duration     time.Duration `json:"duration"`
		DateTime     time.Time     `json:"date_time"`
	}

	GetTalkReq struct {
		ConferenceId uint `json:"conference_id" form:"conference_id"`
		Page         int  `json:"page" form:"page"`
		PageSize     int  `json:"page_size" form:"page_size"`
	}

	GetConferenceReq struct {
		Page     int `json:"page" form:"page"`
		PageSize int `json:"page_size" form:"page_size"`
	}

	AddSpeakerReq struct {
		TalkID   uint   `json:"talk_id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	RemoveSpeakerReq struct {
		TalkID    uint `json:"talk_id" form:"talk_id"`
		SpeakerID uint `json:"speaker_id" form:"speaker_id"`
	}

	AddParticipantReq struct {
		TalkID   uint   `json:"talk_id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	RemoveParticipantReq struct {
		TalkID        uint `json:"talk_id" form:"talk_id"`
		ParticipantID uint `json:"participant_id" form:"participant_id"`
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

	GetEditHistoryReq struct {
		ConferenceId uint `json:"conference_id" form:"conference_id"`
		Page         int  `json:"page" form:"page"`
		PageSize     int  `json:"page_size" form:"page_size"`
	}
)

func (g GetEditHistoryReq) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.ConferenceId, validation.Required),
		//validation.Field(&g.Page, validation.Required),
		//validation.Field(&g.PageSize, validation.Required),
	)
}

//
func (g GetTalkReq) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.ConferenceId, validation.Required),
		//validation.Field(&g.Page, validation.Required),
		//validation.Field(&g.PageSize, validation.Required),
	)
}

func (g GetConferenceReq) Validate() error {
	return validation.ValidateStruct(&g)
}

func (c CreateConferenceReq) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Required, validation.Length(1, 50)),
		validation.Field(&c.Description, validation.Required, validation.Length(5, 1000)),
		validation.Field(&c.StartDate, validation.Required),
		validation.Field(&c.EndDate, validation.Required),
	)
}

func (c UpdateConferenceReq) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ConferenceID, validation.Required),
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

func (c UpdateTalkReq) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.TalkID, validation.Required),
		validation.Field(&c.Title, validation.Required, validation.Length(1, 50)),
		validation.Field(&c.Description, validation.Required, validation.Length(5, 1000)),
		validation.Field(&c.DateTime, validation.Required),
		validation.Field(&c.ConferenceID, validation.Required),
		validation.Field(&c.Duration, validation.Required),
	)
}

func (g GetSpeakerReq) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.TalkId, validation.Required),
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

func (r RemoveSpeakerReq) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.SpeakerID, validation.Required),
		validation.Field(&r.TalkID, validation.Required),
	)
}

func (g GetParticipantReq) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.TalkId, validation.Required),
	)
}

func (r RemoveParticipantReq) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.ParticipantID, validation.Required),
		validation.Field(&r.TalkID, validation.Required),
	)
}
