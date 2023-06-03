package controller

import (
	"quest/global"
	"quest/model"
	"quest/service"
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
	subs, err := sc.QuestionService.GetQuestionList()
	if err != nil {
		global.ReturnError(ctx, global.Errors.ParamIllegalError)
		return
	}
	ctx.JSON(200, questionListResponse{List: subs, Time: time.Now().Unix()})
}
