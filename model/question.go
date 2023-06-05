package model

import "xorm.io/xorm"

type Question struct {
	Id       int64   `json:"id"`
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

// 获取题目列表
func (model QuestionModel) GetQuestionList(sid int, page int, num int) ([]Question, error) {
	var questions []Question
	err := model.DB.Where("sid = ?", sid).Limit(num, page).Find(&questions)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

// 添加题目
func (model QuestionModel) AddQuestion(question *Question) bool {
	_, err := model.DB.Insert(question)
	return err == nil
}

// 编辑题目
func (model QuestionModel) EditQuestion(question *Question) bool {
	if question.Id == 0 {
		return false
	}
	_, err := model.DB.ID(question.Id).Update(question)
	return err == nil
}

// 删除题目
func (model QuestionModel) DelQuestion(question *Question) bool {
	if question.Id == 0 {
		return false
	}
	_, err := model.DB.Delete(question)
	return err == nil
}
