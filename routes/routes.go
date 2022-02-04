package routes

import (
	"conference/handler"
	"github.com/gin-gonic/gin"
	"strings"
)

func Run(h handler.IHandler, port string) error {
	if !strings.Contains(port, ":") {
		port = ":" + port
	}
	router := gin.Default()
	routerV1 := router.Group("/v1")

	userRouter := routerV1.Group("/user")
	confRouter := routerV1.Group("/conference")

	userRouter.POST("/register", h.RegisterUser)
	userRouter.POST("/login", h.LoginUser)
	userRouter.POST("/logout", h.LogOutUser)

	confRouter.POST("", h.CreateConference)
	confRouter.PUT("", h.UpdateConference)
	confRouter.GET("", h.GetConferences)

	confTalkRouter := confRouter.Group("/talk")
	confTalkRouter.POST("", h.CreateTalk)
	confTalkRouter.PUT("", h.UpdateTalk)
	confTalkRouter.GET("", h.GetTalks)

	confTalkSpeakerRouter := confTalkRouter.Group("/speaker")
	confTalkSpeakerRouter.POST("", h.AddSpeaker)
	confTalkSpeakerRouter.GET("", h.GetSpeaker)
	confTalkSpeakerRouter.DELETE("", h.RemoveSpeaker)

	confTalkParticipantRouter := confTalkRouter.Group("/participant")
	confTalkParticipantRouter.POST("", h.AddParticipant)
	confTalkParticipantRouter.GET("", h.GetParticipant)
	confTalkParticipantRouter.DELETE("", h.RemoveParticipant)

	if er := router.Run(port); er != nil {
		return er
	}

	return nil
}
