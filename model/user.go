package model

import "xorm.io/xorm"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Pass string `json:"-"`
}

type UserModel struct {
	DB *xorm.Engine
}

func (model UserModel) GetUser(id int) (User, error) {
	query := "SELECT id, name, pass FROM user WHERE id = ?"
	var user User
	row, err := model.DB.Where(query, id).Get(&user)
	if err != nil {
		return User{}, err
	}
	if row {
		return user, nil
	}
	return User{}, nil
}
