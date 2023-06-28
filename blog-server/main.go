package main

import (
	"server/core"
	_ "server/docs"
	"server/flag"
	"server/global"
	"server/routers"
)

// @title server API文档
// @description server项目API文档
// @version 1
// @host localhost:8080
// @BasePath /

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接mysql数据库
	global.DB = core.InitGorm()
	// 连接redis
	global.Redis = core.ConnectRedis() // 默认连接0号数据库 也可以指定连接的数据库ConnectRedisDB(数据库编号)

	// 命令行参数绑定 go run main.go -db 进行数据表的创建
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	flag.SwitchOption(option)
	addr := global.Config.System.Addr()
	global.Log.Infof("server运行在：%s", addr)

	global.Router = routers.InitRouter()

	err := global.Router.Run(addr)
	if err != nil {
		global.Log.Fatalln(err.Error())
	}
}
