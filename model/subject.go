package model

import (
	"xorm.io/xorm"
)

type Subject struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type SubjectModel struct {
	DB *xorm.Engine
}

// 获取科目
func (model SubjectModel) GetSubject(name string) (*Subject, error) {
	subject := &Subject{Name: name}

	has, err := model.DB.Get(subject)
	if !has {
		if err == nil {
			return nil, nil
		}
		return nil, err
	}
	return subject, nil
}

// 获取科目列表
func (model SubjectModel) GetSubjectList() ([]Subject, error) {
	var subjects []Subject
	err := model.DB.Find(&subjects)
	if err != nil {
		return nil, err
	}
	return subjects, nil
}

// 添加科目
func (model SubjectModel) AddSubject(subject *Subject) bool {
	_, err := model.DB.Insert(subject)
	return err == nil
}

// 编辑科目
func (model SubjectModel) EditSubject(subject *Subject) bool {
	if subject.Id == 0 {
		return false
	}
	_, err := model.DB.ID(subject.Id).Update(subject)
	return err == nil
}

// 删除科目
func (model SubjectModel) DelSubject(subject *Subject) bool {
	if subject.Id == 0 {
		return false
	}

	// 科目下有题目的话不允许删除
	qm := QuestionModel{DB: model.DB}
	list, _ := qm.GetQuestionList(subject.Id, 1, 1)
	if list != nil && len(list) > 0 {
		return false
	}

	_, err := model.DB.Delete(subject)
	return err == nil
}
