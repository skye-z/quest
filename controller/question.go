package controller

import (
	"quest/global"
	"quest/model"
	"quest/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	QuestionService service.QuestionService
}

type questionListResponse struct {
	List []model.Question `json:"list"`
	Time int64            `json:"time"`
}

// 获取题目列表
func (qc QuestionController) GetQuestionList(ctx *gin.Context) {
	sid := ctx.Query("sid")
	page := ctx.Query("page")
	num := ctx.Query("number")
	if len(sid) == 0 || len(page) == 0 {
		global.ReturnError(ctx, global.Errors.ParamEmptyError)
		return
	}
	iSid, err1 := strconv.ParseInt(sid, 10, 64)
	iPage, err2 := strconv.Atoi(page)
	if err1 != nil || err2 != nil {
		global.ReturnError(ctx, global.Errors.ParamIllegalError)
		return
	}
	var list []model.Question
	if len(num) == 0 {
		var err error
		list, err = qc.QuestionService.GetQuestionList(iSid, iPage, 10)
		if err != nil {
			global.ReturnError(ctx, global.Errors.UnexpectedError)
			return
		}
	} else {
		iNum, err := strconv.Atoi(page)
		if err != nil {
			global.ReturnError(ctx, global.Errors.ParamIllegalError)
			return
		}
		list, err = qc.QuestionService.GetQuestionList(iSid, iPage, iNum)
		if err != nil {
			global.ReturnError(ctx, global.Errors.UnexpectedError)
			return
		}
	}
	ctx.JSON(200, questionListResponse{List: list, Time: time.Now().Unix()})
}

// 删除题目
func (qc QuestionController) DelQuestion(ctx *gin.Context) {
	qid := ctx.Param("id")
	if len(qid) == 0 {
		global.ReturnMessage(ctx, false, "题目编号不能为空")
		return
	}

	state := qc.QuestionService.DelQuestion(qid)
	ctx.JSON(200, addResponse{State: state, Time: time.Now().Unix()})
}
