package handler

import (
	"conference/service"
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Login(ctx *gin.Context)
	RegisterUser(ctx *gin.Context)
	LogOut(ctx *gin.Context)

	CreateConference(ctx *gin.Context)
	UpdateConference(ctx *gin.Context)
	CreateTalk(ctx *gin.Context)
	UpdateTalk(ctx *gin.Context)
	AddSpeaker(ctx *gin.Context)
	RemoveSpeaker(ctx *gin.Context)
	AddParticipant(ctx *gin.Context)
	RemoveParticipant(ctx *gin.Context)
	GetTalks(ctx *gin.Context)
	GetConferences(ctx *gin.Context)
	GetEditHistory(ctx *gin.Context)
}

type Handlers struct {
	userSrv service.IUserService
	confSrv service.IConferenceService
}

func NewHandler(uSrv service.IUserService, cSrv service.IConferenceService) *Handlers {
	return &Handlers{
		userSrv: uSrv,
		confSrv: cSrv,
	}
}
