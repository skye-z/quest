package route

import (
	"log"

	"xorm.io/xorm"

	"quest/controller"
	"quest/model"
	"quest/service"

	"github.com/gin-gonic/gin"
)

func RegisterUser(route *gin.Engine, engine *xorm.Engine) {
	log.Println("[User] route registered")
	// 添加用户路由
	addUserRoute(route, engine)
}

func addUserRoute(route *gin.Engine, engine *xorm.Engine) {
	userModel := model.UserModel{DB: engine}
	userService := service.UserService{UserModel: userModel}
	userController := controller.UserController{UserService: userService}
	route.GET("/users/:id", userController.GetUser)
}
