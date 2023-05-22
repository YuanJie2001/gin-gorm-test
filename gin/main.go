package main

import (
	"yqc-portal/gin/database"
	"yqc-portal/gin/resource"
	"yqc-portal/gin/router"
)

func main() {
	// 初始化配置文件读取
	resource.InitConf("application-db.yaml")
	// 初始化数据库连接
	database.MysqlInit()
	// 初始化路由
	router.InitRouter()
}
