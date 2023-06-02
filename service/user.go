package service

import (
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

	return &user, token, exp, nil
}

func (service UserService) GetUserList(name string, page int, num int) ([]model.User, error) {
	return service.UserModel.GetUserList(name, "", page, num)
}

func (service UserService) GetUser(id int) (*model.User, error) {
	var user = model.User{}
	err := service.UserModel.GetUser(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
