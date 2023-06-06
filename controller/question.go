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

// 导入题目
func (qc QuestionController) ImportQuestion(ctx *gin.Context) {
	// TODO
}

// 添加题目
func (qc QuestionController) AddQuestion(ctx *gin.Context) {
	var form model.Question
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if form.Subject == 0 {
		global.ReturnMessage(ctx, false, "科目编号不能为空")
		return
	}
	if form.Level == 0 {
		form.Level = 1
	}
	if form.Type == 0 {
		form.Type = 1
	}
	if form.Score == 0 {
		form.Score = 1.0
	}
	if len(form.Question) == 0 {
		global.ReturnMessage(ctx, false, "题干不能为空")
		return
	}
	if len(form.Answer) == 0 {
		global.ReturnMessage(ctx, false, "答案不能为空")
		return
	}
	state := qc.QuestionService.QuestionModel.AddQuestion(&form)
	ctx.JSON(200, addResponse{State: state, Time: time.Now().Unix()})
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

// 编辑题目
func (qc QuestionController) EditQuestion(ctx *gin.Context) {
	id := ctx.Param("id")
	if len(id) == 0 {
		global.ReturnMessage(ctx, false, "题目编号不能为空")
		return
	}
	qid, _ := strconv.ParseInt(id, 10, 64)
	var form model.Question
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if form.Subject == 0 {
		global.ReturnMessage(ctx, false, "科目编号不能为空")
		return
	}
	if form.Level == 0 {
		form.Level = 1
	}
	if form.Type == 0 {
		form.Type = 1
	}
	if form.Score == 0 {
		form.Score = 1.0
	}
	if len(form.Question) == 0 {
		global.ReturnMessage(ctx, false, "题干不能为空")
		return
	}
	if len(form.Answer) == 0 {
		global.ReturnMessage(ctx, false, "答案不能为空")
		return
	}
	form.Id = qid
	state := qc.QuestionService.QuestionModel.EditQuestion(&form)
	ctx.JSON(200, addResponse{State: state, Time: time.Now().Unix()})
}
