package controller

import (
	"quest/global"
	"quest/model"
	"quest/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

type loginRequest struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type loginResponse struct {
	User   *model.User `json:"user"`
	Token  string      `json:"token"`
	Expire int64       `json:"expire"`
	Time   int64       `json:"time"`
}

// 登录
func (uc UserController) Login(ctx *gin.Context) {
	var form loginRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if len(form.Name) == 0 || len(form.Pass) == 0 {
		ctx.JSON(200, gin.H{"message": "用户名或密码不能为空"})
		return
	}
	if len(form.Pass) != 32 {
		ctx.JSON(200, gin.H{"message": "禁止传输明文密码"})
		return
	}
	user, token, exp, err := uc.UserService.GetUserLoginInfo(form.Name, form.Pass)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		ctx.JSON(200, gin.H{"message": "用户名或密码错误"})
		return
	}
	ctx.JSON(200, loginResponse{User: user, Token: token, Expire: exp, Time: time.Now().Unix()})
}

type userListResponse struct {
	List []model.User `json:"list"`
	Time int64        `json:"time"`
}

// 获取用户列表
func (uc UserController) GetUserList(ctx *gin.Context) {
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
		if err != nil {
			global.ReturnError(ctx, global.Errors.UnexpectedError)
			return
		}
	} else {
		iNum, err := strconv.Atoi(page)
		if err != nil {
			global.ReturnError(ctx, global.Errors.ParamIllegalError)
			return
		}
		users, err = uc.UserService.GetUserList(name, iPage, iNum)
		if err != nil {
			global.ReturnError(ctx, global.Errors.UnexpectedError)
			return
		}
	}
	ctx.JSON(200, userListResponse{List: users, Time: time.Now().Unix()})
}

// 添加用户
func (uc UserController) AddUser(ctx *gin.Context) {
	var form model.User
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if len(form.Name) == 0 {
		global.ReturnMessage(ctx, false, "用户名不能为空")
		return
	}
	user := model.User{
		Name: form.Name,
	}
	uc.UserService.UserModel.GetUser(&user)
	if user.Id != 0 {
		global.ReturnMessage(ctx, false, "用户已存在")
		return
	}
	if len(form.Nickname) == 0 {
		form.Nickname = form.Name
	}
	if len(form.Pass) == 0 {
		global.ReturnMessage(ctx, false, "密码不能为空")
		return
	}
	state := uc.UserService.UserModel.AddUser(&form)
	ctx.JSON(200, addResponse{State: state, Time: time.Now().Unix()})
}

// 删除用户
func (uc UserController) DelUser(ctx *gin.Context) {
	uid := ctx.Param("id")
	if len(uid) == 0 {
		global.ReturnMessage(ctx, false, "用户编号不能为空")
		return
	}

	state := uc.UserService.DelUser(uid)
	ctx.JSON(200, addResponse{State: state, Time: time.Now().Unix()})
}

// 编辑用户
func (uc UserController) EditUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if len(id) == 0 {
		global.ReturnMessage(ctx, false, "用户编号不能为空")
		return
	}
	uid, _ := strconv.ParseInt(id, 10, 64)
	var form model.User
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if len(form.Name) == 0 {
		global.ReturnMessage(ctx, false, "用户名不能为空")
		return
	}
	if len(form.Nickname) == 0 {
		form.Nickname = form.Name
	}
	form.Id = uid
	state := uc.UserService.UserModel.EditUser(&form)
	ctx.JSON(200, addResponse{State: state, Time: time.Now().Unix()})
}
