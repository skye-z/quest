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

func (sc ExamController) GetExamList(ctx *gin.Context) {
	subs, err := sc.ExamService.GetExamList()
	if err != nil {
		global.ReturnError(ctx, global.Errors.ParamIllegalError)
		return
	}
	ctx.JSON(200, examListResponse{List: subs, Time: time.Now().Unix()})
}
