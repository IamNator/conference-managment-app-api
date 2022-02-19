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

func keyFuncDefault(c *gin.Context) string {
	return c.GetHeader("Authorization")
}

func keyFuncIPAddr(c *gin.Context) string {
	if c.ClientIP() == "" {
		return c.GetHeader("Authorization")
	}
	return c.ClientIP()
}

func errorHandler(c *gin.Context) {
	c.String(http.StatusTooManyRequests, "Too many requests")
}

func Run(h handler.IHandler, port string) error {

	if !strings.Contains(port, ":") {
		port = ":" + port
	}

	storeRateLimit := GinRateLimit.InMemoryStore(1, 5)
	storeRateLimitIP := GinRateLimit.InMemoryStore(1, 5)

	rateLimitMiddleware := GinRateLimit.RateLimiter(keyFuncDefault, errorHandler, storeRateLimit)
	rateLimitMiddlewareIP := GinRateLimit.RateLimiter(keyFuncIPAddr, errorHandler, storeRateLimitIP)

	router := gin.Default()
	router.GET("", rateLimitMiddlewareIP, func(ctx *gin.Context) {
		docs := os.Getenv("API_DOCS_URL")
		ctx.Data(http.StatusOK, "text/html", []byte(fmt.Sprintf(`<html><head><a href="%s">click here for docs</a></head></html>`, docs)))
	})

	routerV1 := router.Group("/v1")

	userRouter := routerV1.Group("/user")
	confRouter := routerV1.Group("/conference")

	userRouter.POST("/register", rateLimitMiddlewareIP, h.RegisterUser)
	userRouter.POST("/login", rateLimitMiddlewareIP, h.LoginUser)
	userRouter.POST("/logout", rateLimitMiddleware, h.LogOutUser)

	confRouter.GET("/history", rateLimitMiddlewareIP, h.GetEditHistory)

	confRouter.POST("", rateLimitMiddleware, h.CreateConference)
	confRouter.PUT("", rateLimitMiddleware, h.UpdateConference)
	confRouter.GET("", rateLimitMiddlewareIP, h.GetConferences)

	confTalkRouter := confRouter.Group("/talk")
	confTalkRouter.POST("", rateLimitMiddleware, h.CreateTalk)
	confTalkRouter.PUT("", rateLimitMiddleware, h.UpdateTalk)
	confTalkRouter.GET("", rateLimitMiddlewareIP, h.GetTalks)

	confTalkSpeakerRouter := confTalkRouter.Group("/speaker")
	confTalkSpeakerRouter.POST("", rateLimitMiddleware, h.AddSpeaker)
	confTalkSpeakerRouter.GET("", rateLimitMiddlewareIP, h.GetSpeaker)
	confTalkSpeakerRouter.DELETE("", rateLimitMiddleware, h.RemoveSpeaker)

	confTalkParticipantRouter := confTalkRouter.Group("/participant")
	confTalkParticipantRouter.POST("", rateLimitMiddleware, h.AddParticipant)
	confTalkParticipantRouter.GET("", rateLimitMiddlewareIP, h.GetParticipant)
	confTalkParticipantRouter.DELETE("", rateLimitMiddleware, h.RemoveParticipant)

	if er := router.Run(port); er != nil {
		return er
	}

	return nil
}
