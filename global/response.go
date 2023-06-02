package global

import "github.com/gin-gonic/gin"

func ReturnError(ctx *gin.Context, err CustomError) {
	ctx.JSON(200, err)
	ctx.Abort()
}
