package main

import (
	"fmt"
	"log"
	"os"

	"quest/config"
	"quest/route"
	"quest/service"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

func main() {
	fmt.Println(`   ____               __ 
  / __ \__ _____ ___ / /_
 / /_/ / // / -_|_-</ __/
 \___\_\_,_/\__/___/\__/ 

= Developed by SkyeZhang =`)
	log.Println("[Core] release model")
	// 关闭调试
	// gin.SetMode(gin.ReleaseMode)
	// 初始化系统配置
	config.Init()
	// 初始化数据库
	engine := loadDBEngine()
	go service.InitDatabase(engine)
	log.Println("[Core] service starting")
	// 注册路由
	r := register(engine)
	// 获取端口号
	port := getPort()
	log.Println("[Core] service started, port is", port)
	r.Run(":" + port)
	engine.Close()
}

func loadDBEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("sqlite", "./local.store")
	if err != nil {
		panic(err)
	}
	return engine
}

func register(engine *xorm.Engine) *gin.Engine {
	r := gin.Default()
	route.RegisterUser(r, engine)
	return r
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = config.GetString("basic.port")
	}
	if port == "" {
		port = "3000"
	}
	return port
}
