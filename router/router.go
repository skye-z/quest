package router

import (
	"log"
	"quest/controller"
	"quest/model"
	"quest/service"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func RegisterRoute(route *gin.Engine, engine *xorm.Engine) {
	userModel := model.UserModel{DB: engine}
	userService := service.UserService{UserModel: userModel}
	userController := controller.UserController{UserService: userService}
	// 公共路由
	addPublicRoute(route, engine, userController)
	// 私有路由
	private := route.Group("").Use(service.AuthHandler())
	{
		addPrivateRoute(private, engine, userController)
	}

}

func addPublicRoute(route *gin.Engine, engine *xorm.Engine, uc controller.UserController) {
	log.Println("[Core] public route registered")
	route.POST("/user/login", uc.Login)
}

func addPrivateRoute(route gin.IRoutes, engine *xorm.Engine, uc controller.UserController) {
	log.Println("[Core] private route registered")
	route.GET("/user/list", uc.GetList)
}
