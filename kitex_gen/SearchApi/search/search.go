// Code generated by Kitex v0.3.1. DO NOT EDIT.

package search

import (
	"context"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	searchapi "searchengine3090ti/kitex_gen/SearchApi"
)

func serviceInfo() *kitex.ServiceInfo {
	return searchServiceInfo
}

var searchServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Search"
	handlerType := (*searchapi.Search)(nil)
	methods := map[string]kitex.MethodInfo{
		"query":         kitex.NewMethodInfo(queryHandler, newSearchQueryArgs, newSearchQueryResult, false),
		"add":           kitex.NewMethodInfo(addHandler, newSearchAddArgs, newSearchAddResult, false),
		"relatedQuery":  kitex.NewMethodInfo(relatedQueryHandler, newSearchRelatedQueryArgs, newSearchRelatedQueryResult, false),
		"findID":        kitex.NewMethodInfo(findIDHandler, newSearchFindIDArgs, newSearchFindIDResult, false),
		"queryIDNumber": kitex.NewMethodInfo(queryIDNumberHandler, newSearchQueryIDNumberArgs, newSearchQueryIDNumberResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "searchapi",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.3.1",
		Extra:           extra,
	}
	return svcInfo
}

func queryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*searchapi.SearchQueryArgs)
	realResult := result.(*searchapi.SearchQueryResult)
	success, err := handler.(searchapi.Search).Query(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSearchQueryArgs() interface{} {
	return searchapi.NewSearchQueryArgs()
}

func newSearchQueryResult() interface{} {
	return searchapi.NewSearchQueryResult()
}

func addHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*searchapi.SearchAddArgs)
	realResult := result.(*searchapi.SearchAddResult)
	success, err := handler.(searchapi.Search).Add(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSearchAddArgs() interface{} {
	return searchapi.NewSearchAddArgs()
}

func newSearchAddResult() interface{} {
	return searchapi.NewSearchAddResult()
}

func relatedQueryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*searchapi.SearchRelatedQueryArgs)
	realResult := result.(*searchapi.SearchRelatedQueryResult)
	success, err := handler.(searchapi.Search).RelatedQuery(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSearchRelatedQueryArgs() interface{} {
	return searchapi.NewSearchRelatedQueryArgs()
}

func newSearchRelatedQueryResult() interface{} {
	return searchapi.NewSearchRelatedQueryResult()
}

func findIDHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*searchapi.SearchFindIDArgs)
	realResult := result.(*searchapi.SearchFindIDResult)
	success, err := handler.(searchapi.Search).FindID(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSearchFindIDArgs() interface{} {
	return searchapi.NewSearchFindIDArgs()
}

func newSearchFindIDResult() interface{} {
	return searchapi.NewSearchFindIDResult()
}

func queryIDNumberHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*searchapi.SearchQueryIDNumberArgs)
	realResult := result.(*searchapi.SearchQueryIDNumberResult)
	success, err := handler.(searchapi.Search).QueryIDNumber(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSearchQueryIDNumberArgs() interface{} {
	return searchapi.NewSearchQueryIDNumberArgs()
}

func newSearchQueryIDNumberResult() interface{} {
	return searchapi.NewSearchQueryIDNumberResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Query(ctx context.Context, req *searchapi.QueryRequest) (r *searchapi.QueryResponse, err error) {
	var _args searchapi.SearchQueryArgs
	_args.Req = req
	var _result searchapi.SearchQueryResult
	if err = p.c.Call(ctx, "query", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Add(ctx context.Context, req *searchapi.AddRequest) (r *searchapi.AddResponse, err error) {
	var _args searchapi.SearchAddArgs
	_args.Req = req
	var _result searchapi.SearchAddResult
	if err = p.c.Call(ctx, "add", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RelatedQuery(ctx context.Context, req *searchapi.RelatedQueryRequest) (r *searchapi.RelatedQueryResponse, err error) {
	var _args searchapi.SearchRelatedQueryArgs
	_args.Req = req
	var _result searchapi.SearchRelatedQueryResult
	if err = p.c.Call(ctx, "relatedQuery", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FindID(ctx context.Context, req *searchapi.FindIDRequest) (r *searchapi.FindIDResponse, err error) {
	var _args searchapi.SearchFindIDArgs
	_args.Req = req
	var _result searchapi.SearchFindIDResult
	if err = p.c.Call(ctx, "findID", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryIDNumber(ctx context.Context, req *searchapi.QueryIDNumberRequest) (r *searchapi.QueryIDNumberResponse, err error) {
	var _args searchapi.SearchQueryIDNumberArgs
	_args.Req = req
	var _result searchapi.SearchQueryIDNumberResult
	if err = p.c.Call(ctx, "queryIDNumber", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}