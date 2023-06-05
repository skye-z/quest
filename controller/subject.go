package controller

import (
	"quest/global"
	"quest/model"
	"quest/service"
	"time"

	"github.com/gin-gonic/gin"
)

type SubjectController struct {
	SubjectService service.SubjectService
}

type subjectListResponse struct {
	List []model.Subject `json:"list"`
	Time int64           `json:"time"`
}

func (sc SubjectController) GetSubjectList(ctx *gin.Context) {
	subs, err := sc.SubjectService.GetSubjectList()
	if err != nil {
		global.ReturnError(ctx, global.Errors.ParamIllegalError)
		return
	}
	ctx.JSON(200, subjectListResponse{List: subs, Time: time.Now().Unix()})
}

type addRequest struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type addResponse struct {
	State bool  `json:"state"`
	Time  int64 `json:"time"`
}

func (sc SubjectController) AddSubject(ctx *gin.Context) {
	var form addRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if len(form.Name) == 0 {
		ctx.JSON(200, gin.H{"message": "科目名称不能为空"})
		return
	}
	if len(form.Tag) == 0 {
		form.Tag = ""
	}
	state := sc.SubjectService.AddSubject(form.Name, form.Tag)
	ctx.JSON(200, addResponse{State: state, Time: time.Now().Unix()})
}
