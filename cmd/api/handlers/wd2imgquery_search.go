package handlers

import (
	"context"
	"searchengine3090ti/cmd/api/rpc"
	searchapi "searchengine3090ti/kitex_gen/SearchApi"
	"searchengine3090ti/pkg/constants"
	"searchengine3090ti/pkg/errno"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Wd2ImgQuery(c *gin.Context) {
	var searchQueryVar wd2imgQueryParam
	// if err := c.ShouldBind(searchQueryVar); err != nil {
	// 	SendResponse(c, errno.ConvertErr(err), nil)
	// }
	//请求字段特殊情况处理
	//TODO
	// if searchQueryVar.Page < 1 ||
	// 	searchQueryVar.Order > 1 ||
	// 	searchQueryVar.Order < 0 {
	// 	SendResponse(c, errno.ParamErr, nil)
	// }
	//RPC实际请求搜索
	searchQueryVar.QueryText = c.PostForm("QueryText")
	searchQueryVar.FilterText = c.PostForm("FilterText")
	page_int, _ := strconv.Atoi(c.PostForm("page"))
	limit_int, _ := strconv.Atoi(c.PostForm("limit"))
	order_int, _ := strconv.Atoi(c.PostForm("order"))
	searchQueryVar.Page = int64(page_int)
	searchQueryVar.Limit = int64(limit_int)
	searchQueryVar.Order = int64(order_int)
	req := &searchapi.Wd2imgQueryRequest{
		QueryText:  searchQueryVar.QueryText,
		FilterText: searchQueryVar.FilterText,
		Page:       searchQueryVar.Page,
		Limit:      searchQueryVar.Limit,
		Order:      searchQueryVar.Order,
	}
	time, total, pagecount, page, limit, contents, err := rpc.Wd2imgQuery(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	} else {
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
}
