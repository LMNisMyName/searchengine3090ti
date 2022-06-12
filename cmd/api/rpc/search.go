package rpc

import (
	"context"
	"errors"
	searchapi "searchengine3090ti/kitex_gen/SearchApi"
	"searchengine3090ti/kitex_gen/SearchApi/search"
	"searchengine3090ti/pkg/constants"
	"searchengine3090ti/pkg/middleware"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var SearchClient search.Client

func initSearchRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := search.NewClient(
		constants.SearchServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(30*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	SearchClient = c
}

// Add add search index info
func Add(ctx context.Context, req *searchapi.AddRequest) error {
	resp, err := SearchClient.Add(ctx, req)
	if err != nil {
		return err
	}
	if !resp.Status {
		return errors.New("Add index fail")
	}
	return nil
}

// Query get search info
func Query(ctx context.Context, req *searchapi.QueryRequest) (float64, int64, int64, int64, int64, []*searchapi.AddRequest, error) {
	resp, err := SearchClient.Query(ctx, req)
	if err != nil {
		return 0, 0, 0, 0, 0, nil, err
	}
	return resp.Time, resp.Total, resp.Pagecount, resp.Page, resp.Limit, resp.Contents, nil
}

// RelatedSearch get related-search info
func RelatedSearch(ctx context.Context, req *searchapi.RelatedQueryRequest) ([]string, error) {
	resp, err := SearchClient.RelatedQuery(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.RelatedTexts, nil
}

// FindID check a record if exists by ID
func FindID(ctx context.Context, req *searchapi.FindIDRequest) ([]*searchapi.AddRequest, error) {
	resp, err := SearchClient.FindID(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Entries, err
}
func ImgQuery(ctx context.Context, req *searchapi.ImgQueryRequest) (float64, int64, int64, int64, int64, []*searchapi.AddRequest, error) {
	resp, err := SearchClient.Imgquery(ctx, req)
	if err != nil {
		return 0, 0, 0, 0, 0, nil, err
	}
	return resp.Time, resp.Total, resp.Pagecount, resp.Page, resp.Limit, resp.Contents, nil
}
func Wd2imgQuery(ctx context.Context, req *searchapi.Wd2imgQueryRequest) (float64, int64, int64, int64, int64, []*searchapi.AddRequest, error) {
	resp, err := SearchClient.Wd2imgquery(ctx, req)
	if err != nil {
		return 0, 0, 0, 0, 0, nil, err
	}
	return resp.Time, resp.Total, resp.Pagecount, resp.Page, resp.Limit, resp.Contents, nil
}
