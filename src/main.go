package main

import (
	"./routers"
	"github.com/gin-gonic/gin"
)

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()

	//加载路由
	routers.RoutersInit(r)

	//启动路由
	r.Run()
}
