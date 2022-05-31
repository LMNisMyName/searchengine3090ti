package handlers

import (
	"context"
	"searchengine3090ti/cmd/api/rpc"
	searchapi "searchengine3090ti/kitex_gen/SearchApi"
	"searchengine3090ti/pkg/constants"
	"searchengine3090ti/pkg/errno"

	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) {
	var searchQueryVar QueryParam
	if err := c.ShouldBind(searchQueryVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	//请求字段特殊情况处理
	//TODO
	if searchQueryVar.Page < 1 ||
		searchQueryVar.Order > 1 ||
		searchQueryVar.Order < 0 {
		SendResponse(c, errno.ParamErr, nil)
	}
	//RPC实际请求搜索
	req := &searchapi.QueryRequest{
		QueryText:  searchQueryVar.QueryText,
		FilterText: searchQueryVar.FilterText,
		Page:       searchQueryVar.Page,
		Limit:      searchQueryVar.Limit,
		Order:      searchQueryVar.Order,
	}
	time, total, pagecount, page, limit, contents, err := rpc.Query(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	SendResponse(c, errno.Success,
		map[string]interface{}{
			constants.Time:      time,
			constants.Total:     total,
			constants.PageCount: pagecount,
			constants.Page:      page,
			constants.Limit:     limit,
			constants.Contents:  contents,
		},
	)
}
