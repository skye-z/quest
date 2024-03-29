package router

import (
	"log"
	"quest/controller"
	"quest/global"
	"quest/model"
	"quest/service"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func RegisterRoute(route *gin.Engine, engine *xorm.Engine) {
	// 公用路由
	addCommonRoute(route, engine)
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

func addCommonRoute(route *gin.Engine, engine *xorm.Engine) {
	route.GET("/", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/app"
		route.HandleContext(ctx)
	})
}

func addPublicRoute(route *gin.Engine, engine *xorm.Engine, uc controller.UserController) {
	log.Println("[Core] public route registered")
	route.GET("/api/init", func(ctx *gin.Context) {
		appName := global.GetString("basic.name")
		appMask := global.GetString("basic.mask")
		info := map[string]string{
			"name":    appName,
			"mask":    appMask,
			"version": global.Version,
		}
		ctx.JSON(200, info)
	})
	// 登录
	route.POST("/api/user/login", uc.Login)
}

func addPrivateRoute(route gin.IRoutes, engine *xorm.Engine, uc controller.UserController) {
	log.Println("[Core] private route registered")

	sys := controller.SystemController{}

	subjectModel := model.SubjectModel{DB: engine}
	subjectService := service.SubjectService{SubjectModel: subjectModel}
	sc := controller.SubjectController{SubjectService: subjectService}

	questionModel := model.QuestionModel{DB: engine}
	questionService := service.QuestionService{QuestionModel: questionModel}
	qc := controller.QuestionController{QuestionService: questionService}

	examModel := model.ExamModel{DB: engine}
	examService := service.ExamService{ExamModel: examModel}
	ec := controller.ExamController{ExamService: examService}

	// 管理-获取系统信息
	route.GET("/api/sys/info", sys.GetSystemInfo)
	// 管理-获取系统占用情况
	route.GET("/api/sys/use", sys.GetSystemUse)
	// 管理-获取用户列表
	route.GET("/api/user/list", uc.GetUserList)
	// 管理-添加用户
	route.POST("/api/user/add", uc.AddUser)
	// 管理-删除用户
	route.DELETE("/api/user/:id", uc.DelUser)
	// 管理-编辑用户
	route.POST("/api/user/:id", uc.EditUser)
	// 获取科目列表
	route.GET("/api/subject/list", sc.GetSubjectList)
	// 管理-添加科目
	route.POST("/api/subject/add", sc.AddSubject)
	// 管理-删除科目
	route.DELETE("/api/subject/:id", sc.DelSubject)
	// 管理-编辑科目
	route.POST("/api/subject/:id", sc.EditSubject)
	// 获取题目列表
	route.GET("/api/question/list", qc.GetQuestionList)
	// 获取题目数量
	route.GET("/api/question/number", qc.GetQuestionnNumber)
	// 管理-添加题目
	route.POST("/api/question/add", qc.AddQuestion)
	// 管理-删除题目
	route.DELETE("/api/question/:id", qc.DelQuestion)
	// 管理-编辑题目
	route.POST("/api/question/:id", qc.EditQuestion)
	// 获取考试列表
	route.GET("/api/exam/list", ec.GetExamList)
	// 管理-添加考试
	route.POST("/api/exam/add", ec.AddExam)
	// 管理-删除考试
	route.DELETE("/api/exam/:id", ec.DelExam)
	// 管理-编辑考试
	route.POST("/api/exam/:id", ec.EditExam)
}
