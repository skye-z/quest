package service

import (
	"quest/model"
	"strconv"
)

type QuestionService struct {
	QuestionModel model.QuestionModel
}

// 获取题目列表
func (service QuestionService) GetQuestionList(sid int64, keyword string, page int, num int) ([]model.Question, error) {
	subs, err := service.QuestionModel.GetQuestionList(sid, keyword, page, num)
	if err != nil {
		return make([]model.Question, 0), err
	}
	if subs == nil {
		return make([]model.Question, 0), nil
	}
	return subs, nil
}

// 获取题目数量
func (service QuestionService) GetQuestionNumber(sid int64, keyword string) (int64, error) {
	number, err := service.QuestionModel.GetQuestionNumber(sid, keyword)
	if err != nil {
		return 0, err
	}
	return number, nil
}

// 删除题目
func (service QuestionService) DelQuestion(id string) bool {
	qid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return false
	}
	sub := model.Question{
		Id: qid,
	}
	state := service.QuestionModel.DelQuestion(&sub)
	return state
}
