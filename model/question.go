package model

import "xorm.io/xorm"

type Question struct {
	ID       int     `json:"id"`
	Subject  int     `json:"sid"`
	Level    int     `json:"level"`
	Type     int     `json:"type"`
	Question string  `json:"question"`
	Options  string  `json:"options"`
	Answer   string  `json:"answer"`
	Score    float32 `json:"score"`
}

type QuestionModel struct {
	DB *xorm.Engine
}

func (model QuestionModel) GetQuestion(id int) (Question, error) {
	query := "SELECT id, sid, level, type, question, options, answer, score FROM question WHERE id = ?"
	var question Question
	row, err := model.DB.Where(query, id).Get(&question)
	if err != nil {
		return Question{}, err
	}
	if row {
		return question, nil
	}
	return Question{}, nil
}
