package service

import (
	"quest/model"
)

type ExamService struct {
	ExamModel model.ExamModel
}

func (service ExamService) GetExamList() ([]model.Exam, error) {
	subs, err := service.ExamModel.GetExamList()
	if err != nil {
		return make([]model.Exam, 0), err
	}
	if subs == nil {
		return make([]model.Exam, 0), nil
	}
	return subs, nil
}
