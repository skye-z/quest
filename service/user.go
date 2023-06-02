package service

import (
	"quest/model"

	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	UserModel model.UserModel
}

func (service UserService) GetUserLoginInfo(name string, pass string) (*model.User, string, error) {
	users, err := service.UserModel.GetUserList(name, pass, 0, 1)
	if err != nil {
		return nil, "", err
	}
	if len(users) != 1 {
		return nil, "", nil
	}
	user := users[0]
	info := jwt.MapClaims{
		"iss": "skye-quest-auth",
		"sub": user.Name,
	}
	tc := jwt.NewWithClaims(jwt.SigningMethodES256, info)
	token, err := tc.SignedString([]byte("secret"))
	if err != nil {
		return nil, "", err
	}
	return &user, token, nil
}

func (service UserService) GetUser(id int) (*model.User, error) {
	var user = model.User{}
	err := service.UserModel.GetUser(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
