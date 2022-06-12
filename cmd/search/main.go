package main

import (
	"net"
	"searchengine3090ti/cmd/search/dal/db"
	"searchengine3090ti/cmd/search/relatedsearch"
	"searchengine3090ti/cmd/search/tokenizer"
	searchapi "searchengine3090ti/kitex_gen/SearchApi/search"
	"searchengine3090ti/pkg/constants"
	"searchengine3090ti/pkg/middleware"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	db.Init()
	tokenizer.Init()
	relatedsearch.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8887")
	if err != nil {
		panic(err)
	}
	Init()
	svr := searchapi.NewServer(new(SearchImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.SearchServiceName}),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithRegistry(r),
	)

	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
