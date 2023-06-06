package service

import (
	"quest/model"
	"strconv"
)

type UserService struct {
	UserModel model.UserModel
}

// 获取用户登录信息
func (service UserService) GetUserLoginInfo(name string, pass string) (*model.User, string, int64, error) {
	users, err := service.UserModel.GetUserList(name, pass, 0, 1)
	if err != nil {
		return nil, "", 0, err
	}
	if len(users) != 1 {
		return nil, "", 0, nil
	}
	user := users[0]
	token, exp, err := GenerateToken(&user)
	if err != nil {
		return nil, "", 0, err
	}

	return &user, token, exp, nil
}

// 获取用户列表
func (service UserService) GetUserList(name string, page int, num int) ([]model.User, error) {
	users, err := service.UserModel.GetUserList(name, "", page, num)
	if err != nil {
		return make([]model.User, 0), err
	}
	if users == nil {
		return make([]model.User, 0), nil
	}
	return users, nil
}

// 获取用户
func (service UserService) GetUser(id int64) (*model.User, error) {
	var user = model.User{
		Id: id,
	}
	err := service.UserModel.GetUser(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 删除用户
func (service UserService) DelUser(id string) bool {
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return false
	}
	sub := model.User{
		Id: uid,
	}
	state := service.UserModel.DelUser(&sub)
	return state
}
