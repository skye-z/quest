package model

import (
	"database/sql"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserModel struct {
	DB *sql.DB
}

func (um UserModel) GetUser(id int) (User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := um.DB.QueryRow(query, id)
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
