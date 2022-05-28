package main

import (
	"web/handlers"
	"web/rpc"

	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.Init()
}

func main() {
	Init()
	router := gin.Default()
	//注册search服务的路由和处理函数
	search := router.Group("/search")
	search.GET("/add", handlers.Add)
	search.POST("/query", handlers.Query)
	search.GET("/relatedsearch", handlers.RelatedSearch)

	err := router.Run() //默认监听8080端口
	if err != nil {
		panic(err)
	}
}
