package main

import (
	"gin-blog/model"
	"gin-blog/routes"
)

func main() {
	// 数据库初始化
	model.InitDb()

	// 路由路径初始化
	routes.InitRouter()
}
