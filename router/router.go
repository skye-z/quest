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
	// 登录
	route.POST("/user/login", uc.Login)
}

func addPrivateRoute(route gin.IRoutes, engine *xorm.Engine, uc controller.UserController) {
	log.Println("[Core] private route registered")

	subjectModel := model.SubjectModel{DB: engine}
	subjectService := service.SubjectService{SubjectModel: subjectModel}
	sc := controller.SubjectController{SubjectService: subjectService}

	// 管理-获取用户列表
	route.GET("/user/list", uc.GetUserList)
	// 获取科目列表
	route.GET("/subject/list", sc.GetSubjectList)
	// 获取题目列表
	// route.GET("/question/list", uc.GetList)
	// 获取考试列表
	// route.GET("/exam/list", uc.GetList)
}
