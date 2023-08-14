/*
常用工具

BetaX Quest
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/

package global

import (
	"quest/model"

	"github.com/gin-gonic/gin"
)

// 校验权限
func CheckAuth(ctx *gin.Context, admin bool, edit bool) bool {
	param, exists := ctx.Get("user")
	if exists {
		user := param.(model.User)
		return (edit && (user.Edit || user.Admin)) || (admin && user.Admin)
	} else {
		return false
	}
}
