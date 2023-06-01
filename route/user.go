package route

import (
	"database/sql"
	"quest/controller"
	"quest/model"
	"quest/service"

	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	route := gin.Default()

	db, err := sql.Open("mysql", "root:password@/myapp")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userModel := model.UserModel{DB: db}
	userService := service.UserService{UserModel: userModel}
	userController := controller.UserController{UserService: userService}
	route.GET("/users/:id", userController.GetUser)

	return route
}
