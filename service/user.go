package service

import (
	"log"
	"quest/model"
)

type UserService struct {
	UserModel model.UserModel
}

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

	// 测试
	state, uid, name, err := ValidateToken(token)
	log.Println("state: %b\tuserId: %v", state, uid)

	return &user, token, exp, nil
}

func (service UserService) GetUser(id int) (*model.User, error) {
	var user = model.User{}
	err := service.UserModel.GetUser(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
