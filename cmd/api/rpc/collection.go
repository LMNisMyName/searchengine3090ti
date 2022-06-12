package rpc

import (
	"context"
	"searchengine3090ti/kitex_gen/collectionModel"
	"searchengine3090ti/kitex_gen/collectionModel/collectionservice"
	"searchengine3090ti/pkg/constants"
	"searchengine3090ti/pkg/errno"
	"searchengine3090ti/pkg/middleware"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var collectionClient collectionservice.Client

func initCollectionRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := collectionservice.NewClient(
		constants.CollectionServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithMiddleware(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	collectionClient = c
}

func CreateCollection(ctx context.Context, req *collectionModel.CreateColltRequest) error {
	resp, err := collectionClient.CreateCollection(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func DeleteCollection(ctx context.Context, req *collectionModel.DeleteColltRequest) error {
	resp, err := collectionClient.DeleteCollection(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func GetCollection(ctx context.Context, req *collectionModel.GetColltRequest) (string, []int64, error) {
	resp, err := collectionClient.GetCollection(ctx, req)
	if err != nil {
		return "", nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return "", nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Name, resp.Entry, nil
}

func MGetCollection(ctx context.Context, req *collectionModel.MGetColltResquest) ([]string, []int64, error) {
	resp, err := collectionClient.MGetCollection(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	names := []string{}
	colltIds := []int64{}
	for _, collt := range resp.Collections {
		names = append(names, collt.Name)
		colltIds = append(colltIds, collt.ColltId)
	}
	return names, colltIds, nil
}

func AddEntry(ctx context.Context, req *collectionModel.AddEntryRequest) error {
	resp, err := collectionClient.AddEntry(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func DeleteEntry(ctx context.Context, req *collectionModel.DeleteEntryRequest) error {
	resp, err := collectionClient.DeleteEntry(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func SetName(ctx context.Context, req *collectionModel.SetNameRequest) error {
	resp, err := collectionClient.SetName(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}
