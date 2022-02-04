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
