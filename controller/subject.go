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
