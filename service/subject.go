package service

import (
	"quest/model"
	"strconv"
)

type SubjectService struct {
	SubjectModel model.SubjectModel
}

// 获取指定科目
func (service SubjectService) GetSubject(name string) *model.Subject {
	sub, _ := service.SubjectModel.GetSubject(name)
	return sub
}

// 获取科目列表
func (service SubjectService) GetSubjectList() ([]model.Subject, error) {
	subs, err := service.SubjectModel.GetSubjectList()
	if err != nil {
		return make([]model.Subject, 0), err
	}
	if subs == nil {
		return make([]model.Subject, 0), nil
	}
	return subs, nil
}

// 添加科目
func (service SubjectService) AddSubject(name string, tag string) bool {
	sub := model.Subject{
		Name: name,
		Tag:  tag,
	}
	state := service.SubjectModel.AddSubject(&sub)
	return state
}

// 删除科目
func (service SubjectService) DelSubject(id string) bool {
	sid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return false
	}
	sub := model.Subject{
		Id: sid,
	}
	state := service.SubjectModel.DelSubject(&sub)
	return state
}
