package main

import (
	"Ai/model"
	"Ai/routes"
)

func main() {
	// 引用数据库
	model.InitDb()
	routes.InitRouter()
}
