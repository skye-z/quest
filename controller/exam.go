package controller

import (
	"quest/global"
	"quest/model"
	"quest/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ExamController struct {
	ExamService service.ExamService
}

type examListResponse struct {
	List []model.Exam `json:"list"`
	Time int64        `json:"time"`
}

// 获取考试列表
func (ec ExamController) GetExamList(ctx *gin.Context) {
	subs, err := ec.ExamService.GetExamList()
	if err != nil {
		global.ReturnError(ctx, global.Errors.ParamIllegalError)
		return
	}
	ctx.JSON(200, examListResponse{List: subs, Time: time.Now().Unix()})
}

// 添加考试
func (ec ExamController) AddExam(ctx *gin.Context) {
	var form model.Exam
	if !check(ctx, &form) {
		return
	}
	state := ec.ExamService.ExamModel.AddExam(&form)
	ctx.JSON(200, addResponse{State: state, Time: time.Now().Unix()})
}

// 删除考试
func (ec ExamController) DelExam(ctx *gin.Context) {
	eid := ctx.Param("id")
	if len(eid) == 0 {
		global.ReturnMessage(ctx, false, "考试编号不能为空")
		return
	}

	state := ec.ExamService.DelExam(eid)
	ctx.JSON(200, addResponse{State: state, Time: time.Now().Unix()})
}

// 编辑考试
func (ec ExamController) EditExam(ctx *gin.Context) {
	id := ctx.Param("id")
	if len(id) == 0 {
		global.ReturnMessage(ctx, false, "考试编号不能为空")
		return
	}
	eid, _ := strconv.ParseInt(id, 10, 64)
	var form model.Exam
	if !check(ctx, &form) {
		return
	}
	form.Id = eid
	state := ec.ExamService.ExamModel.EditExam(&form)
	ctx.JSON(200, addResponse{State: state, Time: time.Now().Unix()})
}

// 校验数据
func check(ctx *gin.Context, form *model.Exam) bool {
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return false
	}
	if len(form.Name) == 0 {
		global.ReturnMessage(ctx, false, "考试名称不能为空")
		return false
	}
	if form.Subject == 0 {
		global.ReturnMessage(ctx, false, "科目编号不能为空")
		return false
	}
	if len(form.Questions) == 0 {
		global.ReturnMessage(ctx, false, "考题不能为空")
		return false
	}
	return true
}
