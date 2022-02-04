package handler

import (
	"conference/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handlers) CreateConference(c *gin.Context) {

	var confReq model.CreateConferenceReq
	if er := c.BindJSON(&confReq); er != nil {
		c.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := confReq.Validate(); er != nil {
		c.String(http.StatusBadRequest, er.Error())
		return
	}

	conf, er := h.confSrv.CreateConference(model.CreateConferenceReq{})
	if er != nil {
		c.String(http.StatusUnprocessableEntity, er.Error())
		return
	}
	c.JSONP(http.StatusCreated, conf)
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
}

func (h *Handlers) AddSpeaker(ctx *gin.Context) {
	var Req model.AddSpeakerReq
	if er := ctx.BindJSON(&Req); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := Req.Validate(); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}
}

func (h *Handlers) GetSpeaker(ctx *gin.Context)    {}
func (h *Handlers) RemoveSpeaker(ctx *gin.Context) {}

func (h *Handlers) AddParticipant(ctx *gin.Context) {

	var Req model.AddParticipantReq
	if er := ctx.BindJSON(&Req); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := Req.Validate(); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}
}
func (h *Handlers) RemoveParticipant(ctx *gin.Context) {}
func (h *Handlers) GetParticipant(ctx *gin.Context)    {}

func (h *Handlers) GetTalks(ctx *gin.Context)       {}
func (h *Handlers) GetConferences(ctx *gin.Context) {}
func (h *Handlers) GetEditHistory(ctx *gin.Context) {}
