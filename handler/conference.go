package handler

import (
	"conference/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handlers) CreateConference(ctx *gin.Context) {

	var Req model.CreateConferenceReq
	if er := ctx.BindJSON(&Req); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := Req.Validate(); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	conf, er := h.confSrv.CreateConference(Req)
	if er != nil {
		ctx.String(http.StatusUnprocessableEntity, er.Error())
		return
	}
	ctx.JSONP(http.StatusCreated, conf)
}

func (h *Handlers) UpdateConference(ctx *gin.Context) {
	var Req model.UpdateConferenceReq
	if er := ctx.BindJSON(&Req); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := Req.Validate(); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	conf, er := h.confSrv.UpdateConference(Req)
	if er != nil {
		ctx.String(http.StatusUnprocessableEntity, er.Error())
		return
	}

	ctx.JSONP(http.StatusOK, conf)
}
func (h *Handlers) CreateTalk(ctx *gin.Context) {
	var Req model.CreateTalkReq
	if er := ctx.BindJSON(&Req); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := Req.Validate(); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	talk, er := h.confSrv.CreateTalk(Req)
	if er != nil {
		ctx.String(http.StatusUnprocessableEntity, er.Error())
		return
	}

	ctx.JSONP(http.StatusCreated, talk)
}

func (h *Handlers) UpdateTalk(ctx *gin.Context) {
	var Req model.UpdateTalkReq
	if er := ctx.BindJSON(&Req); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := Req.Validate(); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	talk, er := h.confSrv.UpdateTalk(Req)
	if er != nil {
		ctx.String(http.StatusUnprocessableEntity, er.Error())
		return
	}

	ctx.JSONP(http.StatusOK, talk)
}

func (h *Handlers) AddSpeaker(ctx *gin.Context) {

	authToken := strings.TrimSpace(strings.Replace(ctx.GetHeader("Authorization"), "Bearer", "", -1))
	userId, er := h.midWare.Verify(authToken)
	if er != nil {
		ctx.String(http.StatusUnauthorized, er.Error())
		return
	}

	var Req model.AddSpeakerReq
	if er := ctx.BindJSON(&Req); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := Req.Validate(); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	speaker, er := h.confSrv.AddSpeaker(*userId, Req)
	if er != nil {
		ctx.String(http.StatusUnprocessableEntity, er.Error())
		return
	}

	ctx.JSONP(http.StatusCreated, speaker)
}

func (h *Handlers) GetSpeaker(ctx *gin.Context) {

	var Req model.GetSpeakerReq
	if er := ctx.BindQuery(&Req); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := Req.Validate(); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	speakers, er := h.confSrv.GetSpeakers(Req.TalkId, Req.Page, Req.PageSize)
	if er != nil {
		ctx.String(http.StatusUnprocessableEntity, er.Error())
		return
	}

	ctx.JSONP(http.StatusOK, speakers)
}

func (h *Handlers) RemoveSpeaker(ctx *gin.Context) {
	var Req model.RemoveSpeakerReq
	if er := ctx.BindQuery(&Req); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := Req.Validate(); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	er := h.confSrv.RemoveSpeaker(Req.SpeakerID, Req.TalkID)
	if er != nil {
		ctx.String(http.StatusUnprocessableEntity, er.Error())
		return
	}

	ctx.JSONP(http.StatusOK, "successful")
}

func (h *Handlers) AddParticipant(ctx *gin.Context) {

	authToken := strings.TrimSpace(strings.Replace(ctx.GetHeader("Authorization"), "Bearer", "", -1))
	userId, er := h.midWare.Verify(authToken)
	if er != nil {
		ctx.String(http.StatusUnauthorized, er.Error())
		return
	}

	var Req model.AddParticipantReq
	if er := ctx.BindJSON(&Req); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := Req.Validate(); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}
	participant, er := h.confSrv.AddParticipant(*userId, Req)
	if er != nil {
		ctx.String(http.StatusUnprocessableEntity, er.Error())
		return
	}

	ctx.JSONP(http.StatusCreated, participant)

}

func (h *Handlers) RemoveParticipant(ctx *gin.Context) {
	var Req model.RemoveParticipantReq
	if er := ctx.BindQuery(&Req); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := Req.Validate(); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	er := h.confSrv.RemoveParticipant(Req.ParticipantID, Req.TalkID)
	if er != nil {
		ctx.String(http.StatusUnprocessableEntity, er.Error())
		return
	}

	ctx.JSONP(http.StatusOK, "successful")
}

func (h *Handlers) GetParticipant(ctx *gin.Context) {
	var Req model.GetParticipantReq
	if er := ctx.BindQuery(&Req); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := Req.Validate(); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	speakers, er := h.confSrv.GetParticipants(Req.TalkId, Req.Page, Req.PageSize)
	if er != nil {
		ctx.String(http.StatusUnprocessableEntity, er.Error())
		return
	}

	ctx.JSONP(http.StatusOK, speakers)
}

func (h *Handlers) GetTalks(ctx *gin.Context) {
	var Req model.GetTalkReq
	if er := ctx.BindQuery(&Req); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := Req.Validate(); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	speakers, er := h.confSrv.GetTalks(Req.ConferenceId, Req.Page, Req.PageSize)
	if er != nil {
		ctx.String(http.StatusUnprocessableEntity, er.Error())
		return
	}

	ctx.JSONP(http.StatusOK, speakers)
}

func (h *Handlers) GetConferences(ctx *gin.Context) {

	var Req model.GetConferenceReq
	if er := ctx.BindQuery(&Req); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := Req.Validate(); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	speakers, er := h.confSrv.GetConferences(Req.Page, Req.PageSize)
	if er != nil {
		ctx.String(http.StatusUnprocessableEntity, er.Error())
		return
	}

	ctx.JSONP(http.StatusOK, speakers)
}

func (h *Handlers) GetEditHistory(ctx *gin.Context) {

	var Req model.GetEditHistoryReq
	if er := ctx.BindQuery(&Req); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := Req.Validate(); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	speakers, er := h.confSrv.GetEditHistory(Req.ConferenceId, Req.Page, Req.PageSize)
	if er != nil {
		ctx.String(http.StatusUnprocessableEntity, er.Error())
		return
	}

	ctx.JSONP(http.StatusOK, speakers)
}
