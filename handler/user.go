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

func (h *Handlers) RegisterUser(ctx *gin.Context) {

	var Req model.UserSignUpReq
	if er := ctx.BindJSON(&Req); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := Req.Validate(); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	resp, er := h.userSrv.RegisterUser(Req)
	if er != nil {
		ctx.String(http.StatusUnprocessableEntity, er.Error())
		return
	}

	ctx.JSONP(http.StatusCreated, resp)
}

func (h *Handlers) LogOutUser(ctx *gin.Context) {
	var Req model.UserLogOutReq
	if er := ctx.BindQuery(&Req); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := Req.Validate(); er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
		return
	}

	if er := h.userSrv.LogOut(Req.RefreshToken, Req.AccessToken); er != nil {
		ctx.String(http.StatusUnprocessableEntity, er.Error())
		return
	}

	ctx.String(http.StatusOK, "successful")
}
