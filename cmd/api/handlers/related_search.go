package handlers

import (
	"context"
	"searchengine3090ti/cmd/api/rpc"
	searchapi "searchengine3090ti/kitex_gen/SearchApi"
	"searchengine3090ti/pkg/constants"
	"searchengine3090ti/pkg/errno"

	"github.com/gin-gonic/gin"
)

func RelatedSearch(c *gin.Context) {
	var relatedSearchVar RelatedSearchParam
	if err := c.ShouldBind(relatedSearchVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}

	//请求字段特殊情况处理
	//TODO

	//RPC实际请求
	relatedSearchVar.QueryText = c.Query("text")
	req := &searchapi.RelatedQueryRequest{
		QueryText: relatedSearchVar.QueryText,
	}
	relatedTexts, err := rpc.RelatedSearch(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	} else {
		SendResponse(c, errno.Success,
			map[string]interface{}{
				constants.RelatedTexts: relatedTexts,
			},
		)
	}
}
