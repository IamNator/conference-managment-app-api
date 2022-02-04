package handler

import (
	"conference/service"
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	LoginUser(*gin.Context)
	CreateConference(*gin.Context)
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
