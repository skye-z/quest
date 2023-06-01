package service

import (
	"quest/model"
)

type UserService struct {
	UserModel model.UserModel
}

func (us UserService) GetUser(id int) (model.User, error) {
	user, err := us.UserModel.GetUser(id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
