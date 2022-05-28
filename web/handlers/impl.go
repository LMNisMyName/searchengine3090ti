package handlers

import (
	"context"
	searchapi "search/kitex_gen/SearchApi"
	"strconv"
	"web/rpc"

	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	text := ctx.Query("text")
	url := ctx.Query("url")
	resp, err := rpc.SearchClient.Add(context.Background(), &searchapi.AddRequest{Id: int32(id), Text: text, Url: url})
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, *resp)
}

func Query(ctx *gin.Context) {
	queryText := ctx.PostForm("QueryText")
	filterText := ctx.PostForm("FilterText")
	page, _ := strconv.Atoi(ctx.PostForm("page"))
	limit, _ := strconv.Atoi(ctx.PostForm("limit"))
	order, _ := strconv.Atoi(ctx.PostForm("order"))
	resp, err := rpc.SearchClient.Query(context.Background(),
		&searchapi.QueryRequest{QueryText: queryText,
			FilterText: filterText,
			Page:       int32(page),
			Limit:      int32(limit),
			Order:      int32(order)})
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, *resp)
}

func RelatedSearch(ctx *gin.Context) {
	queryText := ctx.Query("text")
	resp, err := rpc.SearchClient.RelatedQuery(context.Background(), &searchapi.RelatedQueryRequest{QueryText: queryText})
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, *resp)
}
