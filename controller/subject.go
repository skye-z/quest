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

// 获取科目列表响应
type subjectListResponse struct {
	List []model.Subject `json:"list"`
	Time int64           `json:"time"`
}

// 获取科目列表
func (sc SubjectController) GetSubjectList(ctx *gin.Context) {
	subs, err := sc.SubjectService.GetSubjectList()
	if err != nil {
		global.ReturnError(ctx, global.Errors.ParamIllegalError)
		return
	}
	ctx.JSON(200, subjectListResponse{List: subs, Time: time.Now().Unix()})
}

// 添加科目请求
type addRequest struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

// 添加科目响应
type addResponse struct {
	State bool  `json:"state"`
	Time  int64 `json:"time"`
}

// 添加科目
func (sc SubjectController) AddSubject(ctx *gin.Context) {
	var form addRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if len(form.Name) == 0 {
		global.ReturnMessage(ctx, false, "科目名称不能为空")
		return
	}
	sub := sc.SubjectService.GetSubject(form.Name)
	if sub != nil {
		global.ReturnMessage(ctx, false, "科目已存在")
		return
	}
	if len(form.Tag) == 0 {
		form.Tag = ""
	}
	state := sc.SubjectService.AddSubject(form.Name, form.Tag)
	ctx.JSON(200, addResponse{State: state, Time: time.Now().Unix()})
}

// 删除科目
func (sc SubjectController) DelSubject(ctx *gin.Context) {
	sid := ctx.Param("id")
	if len(sid) == 0 {
		global.ReturnMessage(ctx, false, "科目编号不能为空")
		return
	}

	state := sc.SubjectService.DelSubject(sid)
	ctx.JSON(200, addResponse{State: state, Time: time.Now().Unix()})
}
