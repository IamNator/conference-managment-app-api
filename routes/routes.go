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

	userRouter.POST("/login", h.LoginUser)
	confRouter.POST("/create", h.CreateConference)

	if er := router.Run(port); er != nil {
		return er
	}

	return nil
}
