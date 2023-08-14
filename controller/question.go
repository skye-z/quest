package controller

import (
	"log"
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
	page := ctx.Query("page")
	iPage, err1 := strconv.Atoi(page)
	num := ctx.Query("number")
	iNum, err2 := strconv.Atoi(num)
	if err1 != nil || err2 != nil {
		global.ReturnError(ctx, global.Errors.ParamIllegalError)
		return
	}
	if iNum == 0 {
		iNum = 10
	}
	keyword := ctx.Query("keyword")
	if len(keyword) > 0 {
		keyword = "%" + keyword + "%"
	}

	var list []model.Question
	sid := ctx.Query("sid")
	if len(sid) == 0 {
		list, err1 = qc.QuestionService.GetQuestionList(-1, keyword, iPage, iNum)
		if err1 != nil {
			global.ReturnError(ctx, global.Errors.UnexpectedError)
			return
		}
	} else {
		iSid, _ := strconv.ParseInt(sid, 10, 64)
		list, err1 = qc.QuestionService.GetQuestionList(iSid, keyword, iPage, iNum)
		if err1 != nil {
			global.ReturnError(ctx, global.Errors.UnexpectedError)
			return
		}
	}
	ctx.JSON(200, questionListResponse{List: list, Time: time.Now().Unix()})
}

type questionNumberResponse struct {
	Number int64 `json:"num"`
	Time   int64 `json:"time"`
}

// 获取题目数量
func (qc QuestionController) GetQuestionnNumber(ctx *gin.Context) {
	var num int64
	var err error
	sid := ctx.Query("sid")
	keyword := ctx.Query("keyword")
	if len(keyword) > 0 {
		keyword = "%" + keyword + "%"
	}
	if len(sid) == 0 {
		num, err = qc.QuestionService.GetQuestionNumber(-1, keyword)
	} else {
		sub, _ := strconv.ParseInt(sid, 10, 64)
		num, err = qc.QuestionService.GetQuestionNumber(sub, keyword)
	}
	if err != nil {
		log.Println(err)
		global.ReturnError(ctx, global.Errors.UnexpectedError)
		return
	}
	ctx.JSON(200, questionNumberResponse{Number: num, Time: time.Now().Unix()})
}

// 添加题目
func (qc QuestionController) AddQuestion(ctx *gin.Context) {
	if !global.CheckAuth(ctx, true, true) {
		global.ReturnMessage(ctx, false, "权限不足")
		return
	}

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
	if len(form.Question) == 0 {
		global.ReturnMessage(ctx, false, "题干不能为空")
		return
	}
	if len(form.Answer) == 0 {
		global.ReturnMessage(ctx, false, "答案不能为空")
		return
	}
	if qc.QuestionService.QuestionModel.ExistQuestion(form.Subject, form.Question) {
		global.ReturnMessage(ctx, false, "本科目已存在此试题")
		return
	}
	state := qc.QuestionService.QuestionModel.AddQuestion(&form)
	ctx.JSON(200, addResponse{State: state, Time: time.Now().Unix()})
}

// 删除题目
func (qc QuestionController) DelQuestion(ctx *gin.Context) {
	if !global.CheckAuth(ctx, true, true) {
		global.ReturnMessage(ctx, false, "权限不足")
		return
	}

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
	if !global.CheckAuth(ctx, true, true) {
		global.ReturnMessage(ctx, false, "权限不足")
		return
	}

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
