package main

import (
	"context"
	"searchengine3090ti/cmd/collection/pack"
	"searchengine3090ti/cmd/collection/service"
	"searchengine3090ti/kitex_gen/collectionModel"
	"searchengine3090ti/pkg/errno"
)

// CollectionServiceImpl implements the last service interface defined in the IDL.
type CollectionServiceImpl struct{}

// CreateCollection implements the CollectionServiceImpl interface.
func (s *CollectionServiceImpl) CreateCollection(ctx context.Context, req *collectionModel.CreateColltRequest) (resp *collectionModel.CreateColltResponse, err error) {
	resp = new(collectionModel.CreateColltResponse)

	if len(req.Name) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateColltService(ctx).CreateCollection(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetCollection implements the CollectionServiceImpl interface.
func (s *CollectionServiceImpl) GetCollection(ctx context.Context, req *collectionModel.GetColltRequest) (resp *collectionModel.GetColltResponse, err error) {
	resp = new(collectionModel.GetColltResponse)

	collts, err := service.NewGetColltService(ctx).GetColletction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	if collts != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.Success)
		resp.Name = collts[0].Name
		resp.Entry = collts[0].Entry
	}
	return resp, nil
}

// MGetCollection implements the CollectionServiceImpl interface.
func (s *CollectionServiceImpl) MGetCollection(ctx context.Context, req *collectionModel.MGetColltResquest) (resp *collectionModel.MGetColltResponse, err error) {
	resp = new(collectionModel.MGetColltResponse)
	collts, err := service.NewMGetUserService(ctx).MGetColletction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	for _, collt := range collts {
		resp.Collections = append(resp.Collections,
			&collectionModel.MGetColltResponse_Collection{
				Name:    collt.Name,
				ColltId: collt.ColltId,
			},
		)
	}
	return resp, nil
}

// DeleteCollection implements the CollectionServiceImpl interface.
func (s *CollectionServiceImpl) DeleteCollection(ctx context.Context, req *collectionModel.DeleteColltRequest) (resp *collectionModel.DeleteColltResponse, err error) {
	resp = new(collectionModel.DeleteColltResponse)
	err = service.NewDeleteColltService(ctx).DeleteCollection(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// AddEntry implements the CollectionServiceImpl interface.
func (s *CollectionServiceImpl) AddEntry(ctx context.Context, req *collectionModel.AddEntryRequest) (resp *collectionModel.AddEntryResponse, err error) {
	resp = new(collectionModel.AddEntryResponse)
	err = service.NewAddEntryService(ctx).AddEntry(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DeleteEntry implements the CollectionServiceImpl interface.
func (s *CollectionServiceImpl) DeleteEntry(ctx context.Context, req *collectionModel.DeleteEntryRequest) (resp *collectionModel.DeleteEntryResponse, err error) {
	resp = new(collectionModel.DeleteEntryResponse)
	err = service.NewDeleteEntryService(ctx).DeleteEntry(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// SetName implements the CollectionServiceImpl interface.
func (s *CollectionServiceImpl) SetName(ctx context.Context, req *collectionModel.SetNameRequest) (resp *collectionModel.SetNameResponse, err error) {
	resp = new(collectionModel.SetNameResponse)
	err = service.NewSetNameService(ctx).SetName(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
