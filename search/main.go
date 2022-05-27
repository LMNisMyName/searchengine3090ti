package main

import (
	"log"
	"net"
	"search/dal"
	searchapi "search/kitex_gen/SearchApi/search"
	"search/relatedsearch"
	"search/tokenizer"

	"github.com/cloudwego/kitex/server"
)

func init() {
	dal.Init()
	tokenizer.Init()
	relatedsearch.Init()
}

func main() {
	//将search服务注册到相关端口
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:2022")
	if err != nil {
		panic(err)
	}
	svr := searchapi.NewServer(new(SearchImpl),
		server.WithServiceAddr(addr))
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
