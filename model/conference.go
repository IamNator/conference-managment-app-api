package model

import "time"

type (
	Conference struct {
		General
		Title       string `json:"title"`
		Description string `json:"description"`
		StartDate   string `json:"Start_date"`
		EndDate     string `json:"end_date"`
		Talks       []Talk `json:"talk"`
	}

	Talk struct {
		General
		ConferenceID uint          `json:"conference_id"`
		Title        string        `json:"title"`
		Description  string        `json:"description"`
		Duration     time.Duration `json:"duration"`
		DateTime     time.Time     `json:"time"`
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
		Title       string `json:"title"`
		Description string `json:"description"`
		StartDate   string `json:"Start_date"`
		EndDate     string `json:"end_date"`
	}
	CreateTalkReq struct {
		ConferenceID uint          `json:"conference_id"`
		Title        string        `json:"title"`
		Description  string        `json:"description"`
		Duration     time.Duration `json:"duration"`
		DateTime     time.Time     `json:"time"`
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
