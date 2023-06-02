package controller

import (
	"quest/model"
	"quest/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

type LoginRequest struct {
	Name string `json:"name"`
	Pass string `json:"Pass"`
}

type LoginResponse struct {
	User  *model.User `json:"user"`
	Token string      `json:"token"`
}

func (uc UserController) Login(ctx *gin.Context) {
	var form LoginRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, token, err := uc.UserService.GetUserLoginInfo(form.Name, form.Pass)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, LoginResponse{User: user, Token: token})
}
