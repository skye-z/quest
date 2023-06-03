package service

import (
	"quest/model"
)

type QuestionService struct {
	QuestionModel model.QuestionModel
}

func (service QuestionService) GetQuestionList() ([]model.Question, error) {
	subs, err := service.QuestionModel.GetQuestionList()
	if err != nil {
		return make([]model.Question, 0), err
	}
	if subs == nil {
		return make([]model.Question, 0), nil
	}
	return subs, nil
}
