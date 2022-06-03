package idgen

import (
	"context"
	"fmt"
	searchapi "search/kitex_gen/SearchApi"
	"web/rpc"
)

var ID int32

func Init() {
	ret, _ := rpc.SearchClient.QueryIDNumber(context.Background(), &searchapi.QueryIDNumberRequest{})
	ID = ret.Number
	fmt.Println("Current records number: ", ID)
}

func GetID() int32 {
	ID++
	return ID
}
