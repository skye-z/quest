package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"

	"quest/global"
	"quest/router"
	"quest/service"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

//go:embed page/error/* page/dist/*
var f embed.FS

func main() {
	fmt.Print(`┌────────────────────────────┐
│    ____               __   │
│   / __ \__ _____ ___ / /_  │
│  / /_/ / // / -_|_-</ __/  │
│  \___\_\_,_/\__/___/\__/   │
│                            │
│   github.com/skye-z/quest  │
└────────────────────────────┘
Version: ` + global.Version + "\n\n")
	log.Println("[Core] release model")
	// 关闭调试
	gin.SetMode(gin.ReleaseMode)
	// 初始化系统配置
	global.InitConfig()
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
	// 加载静态资源
	fp, _ := fs.Sub(f, "page/dist")
	r.StaticFS("/app", http.FS(fp))
	// 加载模版页面
	templ := template.Must(template.New("").ParseFS(f, "page/error/*.html"))
	r.SetHTMLTemplate(templ)
	// 配置404
	r.NoRoute(func(ctx *gin.Context) {
		data, err := f.ReadFile("page/dist/index.html")
		if err != nil {
			ctx.HTML(http.StatusOK, "404.html", gin.H{
				"title": "404",
			})
			return
		}
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

	// 注册路由
	router.RegisterRoute(r, engine)
	return r
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = global.GetString("basic.port")
	}
	if port == "" {
		port = "12999"
	}
	return port
}
