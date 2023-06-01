package model

import "xorm.io/xorm"

type Exam struct {
	ID        int     `json:"id"`
	User      int     `json:"uid"`
	Subject   int     `json:"sid"`
	Questions string  `json:"questions"`
	Answers   string  `json:"answers"`
	Score     float32 `json:"score"`
}

type ExamModel struct {
	DB *xorm.Engine
}

func (model ExamModel) GetExam(id int) (Exam, error) {
	query := "SELECT id, uid, sid, questions, answers, score FROM exam WHERE id = ?"
	var exam Exam
	row, err := model.DB.Where(query, id).Get(&exam)
	if err != nil {
		return Exam{}, err
	}
	if row {
		return exam, nil
	}
	return Exam{}, nil
}
