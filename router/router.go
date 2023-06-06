package router

import (
	"log"
	"net/http"
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
	route.Static("/app", "./view")
	route.LoadHTMLGlob("./view/error/*.html")
	route.GET("/", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/app"
		route.HandleContext(ctx)
	})
	route.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "404.html", gin.H{
			"title": "404",
		})
	})
}

func addPublicRoute(route *gin.Engine, engine *xorm.Engine, uc controller.UserController) {
	log.Println("[Core] public route registered")
	route.GET("/api/init", func(ctx *gin.Context) {
		appName := global.GetString("basic.name")
		info := map[string]string{
			"name":    appName,
			"version": global.Version,
		}
		ctx.JSON(200, info)
	})
	// 登录
	route.POST("/api/user/login", uc.Login)
}

func addPrivateRoute(route gin.IRoutes, engine *xorm.Engine, uc controller.UserController) {
	log.Println("[Core] private route registered")

	subjectModel := model.SubjectModel{DB: engine}
	subjectService := service.SubjectService{SubjectModel: subjectModel}
	sc := controller.SubjectController{SubjectService: subjectService}

	questionModel := model.QuestionModel{DB: engine}
	questionService := service.QuestionService{QuestionModel: questionModel}
	qc := controller.QuestionController{QuestionService: questionService}

	examModel := model.ExamModel{DB: engine}
	examService := service.ExamService{ExamModel: examModel}
	ec := controller.ExamController{ExamService: examService}

	// 管理-获取用户列表
	route.GET("/api/user/list", uc.GetUserList)
	// 管理-添加用户
	// route.POST("/api/user/add", uc.GetUserList)
	// 管理-删除用户
	route.DELETE("/api/user/:id", uc.DelUser)
	// 管理-编辑用户
	// route.POST("/api/user/:id", uc.GetUserList)
	// 获取科目列表
	route.GET("/api/subject/list", sc.GetSubjectList)
	// 管理-添加科目
	route.POST("/api/subject/add", sc.AddSubject)
	// 管理-删除科目
	route.DELETE("/api/subject/:id", sc.DelSubject)
	// 管理-编辑科目
	// route.POST("/api/subject/:id", sc.AddSubject)
	// 获取题目列表
	route.GET("/api/question/list", qc.GetQuestionList)
	// 管理-导入题目
	// route.POST("/api/question/import", qc.GetQuestionList)
	// 管理-添加题目
	// route.POST("/api/question/add", qc.GetQuestionList)
	// 管理-删除题目
	route.DELETE("/api/question/:id", qc.DelQuestion)
	// 管理-编辑题目
	// route.POST("/api/question/:id", sc.AddSubject)
	// 获取考试列表
	route.GET("/api/exam/list", ec.GetExamList)
	// 管理-添加考试
	// route.POST("/api/exam/add", ec.GetQuestionList)
	// 管理-删除考试
	route.DELETE("/api/exam/:id", ec.DelExam)
	// 管理-编辑考试
	// route.POST("/api/exam/:id", ec.AddSubject)
}
