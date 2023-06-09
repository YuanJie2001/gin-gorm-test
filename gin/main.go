package main

import (
	"yqc-portal/gin/database"
	"yqc-portal/gin/router"
	"yqc-portal/resource"
)

func main() {
	// 初始化配置文件读取
	resource.InitConf("application-db.yml")
	// 初始化数据库连接
	database.MysqlInit()
	// 初始化路由
	router.InitRouter()
}
