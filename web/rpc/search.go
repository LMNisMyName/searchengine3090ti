package rpc

import (
	"search/kitex_gen/SearchApi/search"

	"github.com/cloudwego/kitex/client"
)

var SearchClient search.Client

func initSearchRpc() {
	client, err := search.NewClient("example", client.WithHostPorts("127.0.0.1:2022"))
	if err != nil {
		panic(err)
	}
	SearchClient = client
}
