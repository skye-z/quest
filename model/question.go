package model

import "xorm.io/xorm"

type Question struct {
	Id       int64  `json:"id"`
	Subject  int64  `json:"sid"`
	Level    int    `json:"level"`
	Type     int    `json:"type"`
	Question string `json:"question"`
	Options  string `json:"options"`
	Answer   string `json:"answer"`
}

type QuestionModel struct {
	DB *xorm.Engine
}

// 获取题目列表
func (model QuestionModel) GetQuestionList(sid int64, keyword string, page int, num int) ([]Question, error) {
	var questions []Question
	if sid < 0 {
		var cache *xorm.Session
		if len(keyword) > 0 {
			cache = model.DB.Where("question LIKE ? OR options LIKE ?", keyword, keyword, keyword)
		}
		err := cache.Limit(page*num, (page-1)*num).Find(&questions)
		if err != nil {
			return nil, err
		}
	} else {
		var cache *xorm.Session
		if len(keyword) > 0 {
			cache = model.DB.Where("subject = ? AND ( question LIKE ? OR options LIKE ?)", sid, keyword, keyword, keyword)
		} else {
			cache = model.DB.Where("subject = ?", sid)
		}
		err := cache.Limit(page*num, (page-1)*num).Find(&questions)
		if err != nil {
			return nil, err
		}
	}
	return questions, nil
}

// 获取题目数量
func (model QuestionModel) GetQuestionNumber(sid int64, keyword string) (int64, error) {
	var num int64
	var err error
	var question Question
	if sid < 0 {
		var cache *xorm.Session
		if len(keyword) > 0 {
			cache = model.DB.Where("question LIKE ? OR options LIKE ?", keyword, keyword, keyword)
		}
		num, err = cache.Count(question)
	} else {
		var cache *xorm.Session
		if len(keyword) > 0 {
			cache = model.DB.Where("subject = ? AND ( question LIKE ? OR options LIKE ?)", sid, keyword, keyword, keyword)
		} else {
			cache = model.DB.Where("subject = ?", sid)
		}
		num, err = cache.Count(question)
	}
	if err != nil {
		return 0, err
	}
	return num, nil
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
