package model

import "xorm.io/xorm"

type Subject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type SubjectModel struct {
	DB *xorm.Engine
}

func (model SubjectModel) GetSubject(id int) (Subject, error) {
	query := "SELECT id, name, tag FROM subject WHERE id = ?"
	var subject Subject
	row, err := model.DB.Where(query, id).Get(&subject)
	if err != nil {
		return Subject{}, err
	}
	if row {
		return subject, nil
	}
	return Subject{}, nil
}
