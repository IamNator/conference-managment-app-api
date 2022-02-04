package handler

import (
	"conference/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handlers) LoginUser(c *gin.Context) {

	var loginReq model.UserLoginReq
	if er := c.BindJSON(&loginReq); er != nil {
		c.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := loginReq.Validate(); er != nil {
		c.String(http.StatusBadRequest, er.Error())
		return
	}

	resp, er := h.userSrv.Login(loginReq)
	if er != nil {
		c.String(http.StatusUnprocessableEntity, er.Error())
		return
	}

	c.JSONP(http.StatusOK, resp)
}
