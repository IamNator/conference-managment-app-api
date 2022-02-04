package model

import (
	"conference/testdata"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestCreateConferenceReq_Validate(t *testing.T) {

	type tData struct {
		Req   CreateConferenceReq `json:"req"`
		Error string              `json:"error"`
	}

	Reqs := make([]tData, 0)
	if er := testdata.Load("../testdata/model/create_conference_req.json", &Reqs); er != nil {
		t.Error(er.Error())
	}

	for i, v := range Reqs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			//check for validation errors
			if er := v.Req.Validate(); er != nil {
				assert.Equal(t, v.Error, er.Error())
			}
		})
	}

}

func TestUpdateConferenceReq_Validate(t *testing.T) {

	type tData struct {
		Req   UpdateConferenceReq `json:"req"`
		Error string              `json:"error"`
	}

	Reqs := make([]tData, 0)
	if er := testdata.Load("../testdata/model/update_conference_req.json", &Reqs); er != nil {
		t.Error(er.Error())
	}

	for i, v := range Reqs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			//check for validation errors
			if er := v.Req.Validate(); er != nil {
				assert.Equal(t, v.Error, er.Error())
			}
		})
	}

}

func TestCreateTalkReq_Validate(t *testing.T) {
	type tData struct {
		Req   CreateTalkReq `json:"req"`
		Error string        `json:"error"`
	}

	Reqs := make([]tData, 0)
	if er := testdata.Load("../testdata/model/create_talk_req.json", &Reqs); er != nil {
		t.Error(er.Error())
	}

	for i, v := range Reqs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			//check for validation errors
			if er := v.Req.Validate(); er != nil {
				assert.Equal(t, v.Error, er.Error())
			}
		})
	}
}
func TestUpdateTalkReq_Validate(t *testing.T) {
	type tData struct {
		Req   UpdateTalkReq `json:"req"`
		Error string        `json:"error"`
	}

	Reqs := make([]tData, 0)
	if er := testdata.Load("../testdata/model/update_talk_req.json", &Reqs); er != nil {
		t.Error(er.Error())
	}

	for i, v := range Reqs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			//check for validation errors
			if er := v.Req.Validate(); er != nil {
				assert.Equal(t, v.Error, er.Error())
			}
		})
	}
}

func TestAddSpeakerReq_Validate(t *testing.T) {
	type tData struct {
		Req   AddSpeakerReq `json:"req"`
		Error string        `json:"error"`
	}

	Reqs := make([]tData, 0)
	if er := testdata.Load("../testdata/model/add_speaker_req.json", &Reqs); er != nil {
		t.Error(er.Error())
	}

	for i, v := range Reqs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			//check for validation errors
			if er := v.Req.Validate(); er != nil {
				assert.Equal(t, v.Error, er.Error())
			}
		})
	}
}

func TestAddParticipantReq_Validate(t *testing.T) {
	type tData struct {
		Req   AddParticipantReq `json:"req"`
		Error string            `json:"error"`
	}

	Reqs := make([]tData, 0)
	if er := testdata.Load("../testdata/model/add_participant_req.json", &Reqs); er != nil {
		t.Error(er.Error())
	}

	for i, v := range Reqs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			//check for validation errors
			if er := v.Req.Validate(); er != nil {
				assert.Equal(t, v.Error, er.Error())
			}
		})
	}
}
