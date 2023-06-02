package controller

import (
	"quest/global"
	"quest/model"
	"quest/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

type loginRequest struct {
	Name string `json:"name"`
	Pass string `json:"Pass"`
}

type loginResponse struct {
	User   *model.User `json:"user"`
	Token  string      `json:"token"`
	Expire int64       `json:"expire"`
}

func (uc UserController) Login(ctx *gin.Context) {
	var form loginRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, token, exp, err := uc.UserService.GetUserLoginInfo(form.Name, form.Pass)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, loginResponse{User: user, Token: token, Expire: exp})
}

type LoginRequest struct {
	Name string `json:"name"`
	Pass string `json:"Pass"`
}

func (uc UserController) GetList(ctx *gin.Context) {
	user, state := ctx.Keys["user"].(model.User)
	if !state {
		global.ReturnError(ctx, global.Errors.TokenIllegalError)
		return
	}
	if user.Name != "admin" {
		global.ReturnError(ctx, global.Errors.PermissionDeniedError)
		return
	}
	name := ctx.Query("name")
	page := ctx.Query("page")
	if len(page) == 0 {
		global.ReturnError(ctx, global.Errors.ParamEmptyError)
		return
	}
	iPage, err := strconv.Atoi(page)
	if err != nil {
		global.ReturnError(ctx, global.Errors.ParamIllegalError)
		return
	}
	num := ctx.Query("number")
	var users []model.User
	if len(num) == 0 {
		users, err = uc.UserService.GetUserList(name, iPage, 10)
	} else {
		iNum, err := strconv.Atoi(page)
		if err != nil {
			global.ReturnError(ctx, global.Errors.ParamIllegalError)
			return
		}
		users, err = uc.UserService.GetUserList(name, iPage, iNum)
	}
	if err != nil {
		global.ReturnError(ctx, global.Errors.UnexpectedError)
		return
	}
	ctx.JSON(200, users)
}
