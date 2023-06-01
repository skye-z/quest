package service

import (
	"log"
	"quest/model"

	"xorm.io/xorm"
)

func InitDatabase(engine *xorm.Engine) {
	log.Println("[DB] init user table")
	err := engine.Sync2(new(model.User))
	if err != nil {
		panic(err)
	}

	log.Println("[DB] init subject table")
	err = engine.Sync2(new(model.Subject))
	if err != nil {
		panic(err)
	}

	log.Println("[DB] init question table")
	err = engine.Sync2(new(model.Question))
	if err != nil {
		panic(err)
	}

	log.Println("[DB] init exam table")
	err = engine.Sync2(new(model.Exam))
	if err != nil {
		panic(err)
	}
	log.Println("[DB] init completed")
}
