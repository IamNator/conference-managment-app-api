package routes

import (
	"conference/handler"
	"fmt"
	"github.com/Nebulizer1213/GinRateLimit"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func keyFunc(c *gin.Context) string {
	return c.GetHeader("Authorization")
}

func errorHandler(c *gin.Context) {
	c.String(http.StatusTooManyRequests, "Too many requests")
}

func Run(h handler.IHandler, port string) error {
	if !strings.Contains(port, ":") {
		port = ":" + port
	}

	storeRateLimit := GinRateLimit.InMemoryStore(1, 5)
	rateLimitMiddleware := GinRateLimit.RateLimiter(keyFunc, errorHandler, storeRateLimit)

	router := gin.Default()
	router.Use(rateLimitMiddleware)
	router.GET("", func(ctx *gin.Context) {
		docs := os.Getenv("API_DOCS_URL")
		ctx.Data(http.StatusOK, "text/html", []byte(fmt.Sprintf(`<html><head><a href="%s">click here for docs</a></head></html>`, docs)))
	})

	routerV1 := router.Group("/v1")

	userRouter := routerV1.Group("/user")
	confRouter := routerV1.Group("/conference")

	userRouter.POST("/register", h.RegisterUser)
	userRouter.POST("/login", h.LoginUser)
	userRouter.POST("/logout", h.LogOutUser)

	confRouter.GET("/history", h.GetEditHistory)

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
