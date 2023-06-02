package global

import (
	"github.com/gin-gonic/gin"
)

func ReturnError(ctx *gin.Context, err CustomError) {
	ctx.JSON(200, err)
	ctx.Abort()
}

func ReturnSuccess(ctx *gin.Context, obj any) {
	ctx.JSON(200, obj)
	ctx.Abort()
}
