package service

import (
	"quest/model"
)

type SubjectService struct {
	SubjectModel model.SubjectModel
}

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
