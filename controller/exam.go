package controller

import (
	"quest/global"
	"quest/model"
	"quest/service"
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

func (ec ExamController) GetExamList(ctx *gin.Context) {
	subs, err := ec.ExamService.GetExamList()
	if err != nil {
		global.ReturnError(ctx, global.Errors.ParamIllegalError)
		return
	}
	ctx.JSON(200, examListResponse{List: subs, Time: time.Now().Unix()})
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
