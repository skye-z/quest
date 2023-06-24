package service

import (
	"quest/model"
	"strconv"
)

type ExamService struct {
	ExamModel model.ExamModel
}

func (service ExamService) GetExamList(sid int64) ([]model.Exam, error) {
	subs, err := service.ExamModel.GetExamList(sid)
	if err != nil {
		return make([]model.Exam, 0), err
	}
	if subs == nil {
		return make([]model.Exam, 0), nil
	}
	return subs, nil
}

// 删除考试
func (service ExamService) DelExam(id string) bool {
	eid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return false
	}
	sub := model.Exam{
		Id: eid,
	}
	state := service.ExamModel.DelExam(&sub)
	return state
}
