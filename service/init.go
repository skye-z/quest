package service

import (
	"log"
	"quest/global"
	"quest/model"

	"xorm.io/xorm"
)

func InitDatabase(engine *xorm.Engine) {
	log.Println("[DB] start loading")
	err := engine.Sync2(new(model.User))
	if err != nil {
		panic(err)
	}
	err = engine.Sync2(new(model.Subject))
	if err != nil {
		panic(err)
	}
	err = engine.Sync2(new(model.Question))
	if err != nil {
		panic(err)
	}
	err = engine.Sync2(new(model.Exam))
	if err != nil {
		panic(err)
	}
	if !global.GetBool("basic.install") {
		userModel := model.UserModel{DB: engine}
		adminUser := model.User{
			Name:     "admin",
			Nickname: "管理员",
			Admin:    true,
			Pass:     "25d55ad283aa400af464c76d713c07ad",
		}
		userModel.AddUser(&adminUser)
		global.Set("basic.install", "1")
	}
	log.Println("[DB] loading completed")
}
