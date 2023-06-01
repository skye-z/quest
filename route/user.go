package route

import (
	"xorm.io/xorm"

	"quest/controller"
	"quest/model"
	"quest/service"

	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	route := gin.Default()

	engine, err := xorm.NewEngine("sqlite3", "./local.db")
	if err != nil {
		panic(err)
	}
	defer engine.Close()

	userModel := model.UserModel{DB: engine}
	userService := service.UserService{UserModel: userModel}
	userController := controller.UserController{UserService: userService}
	route.GET("/users/:id", userController.GetUser)

	return route
}
