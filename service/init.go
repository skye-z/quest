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
			Name: "admin",
			Pass: "21232f297a57a5a743894a0e4a801fc3",
		}
		userModel.AddUser(&adminUser)
		global.Set("basic.install", "1")
	}
	log.Println("[DB] loading completed")
}
