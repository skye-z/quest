package model

import "xorm.io/xorm"

type Exam struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Subject   int64  `json:"sid"`
	Questions string `json:"questions"`
}

type ExamModel struct {
	DB *xorm.Engine
}

// 获取考试
func (model ExamModel) GetExam(id int64) (*Exam, error) {
	exam := &Exam{Id: id}

	has, err := model.DB.Get(exam)
	if !has {
		if err == nil {
			return nil, nil
		}
		return nil, err
	}
	return exam, nil
}

// 获取考试列表
func (model ExamModel) GetExamList() ([]Exam, error) {
	var exams []Exam
	err := model.DB.Find(&exams)
	if err != nil {
		return nil, err
	}
	return exams, nil
}

// 添加考试
func (model ExamModel) AddExam(exam *Exam) bool {
	_, err := model.DB.Insert(exam)
	return err == nil
}

// 编辑考试
func (model ExamModel) EditExam(exam *Exam) bool {
	if exam.Id == 0 {
		return false
	}
	_, err := model.DB.ID(exam.Id).Update(exam)
	return err == nil
}

// 删除考试
func (model ExamModel) DelExam(exam *Exam) bool {
	if exam.Id == 0 {
		return false
	}
	_, err := model.DB.Delete(exam)
	return err == nil
}
