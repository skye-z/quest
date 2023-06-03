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

func (sc QuestionController) GetQuestionList(ctx *gin.Context) {
	sid := ctx.Query("sid")
	page := ctx.Query("page")
	num := ctx.Query("number")
	if len(sid) == 0 || len(page) == 0 {
		global.ReturnError(ctx, global.Errors.ParamEmptyError)
		return
	}
	iSid, err1 := strconv.Atoi(sid)
	iPage, err2 := strconv.Atoi(page)
	if err1 != nil || err2 != nil {
		global.ReturnError(ctx, global.Errors.ParamIllegalError)
		return
	}
	var err error
	var list []model.Question
	if len(num) == 0 {
		list, err = sc.QuestionService.GetQuestionList(iSid, iPage, 10)
	} else {
		iNum, err := strconv.Atoi(page)
		if err != nil {
			global.ReturnError(ctx, global.Errors.ParamIllegalError)
			return
		}
		list, err = sc.QuestionService.GetQuestionList(iSid, iPage, iNum)
	}
	if err != nil {
		global.ReturnError(ctx, global.Errors.UnexpectedError)
		return
	}
	ctx.JSON(200, questionListResponse{List: list, Time: time.Now().Unix()})
}
