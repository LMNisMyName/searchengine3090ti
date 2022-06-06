package idgen

import (
	"context"
	"fmt"
	"searchengine3090ti/cmd/api/rpc"
	searchapi "searchengine3090ti/kitex_gen/SearchApi"
)

var ID int32

func Init() {
	ret, err := rpc.SearchClient.QueryIDNumber(context.Background(), &searchapi.QueryIDNumberRequest{})
	if err != nil {
		panic(err)
	}
	ID = ret.Number
	fmt.Println("Current records number: ", ID)
}

func GetID() int32 {
	ID++
	return ID
}
