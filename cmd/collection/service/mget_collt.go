package service

import (
	"context"
	"searchengine3090ti/cmd/collection/dal/db"
	"searchengine3090ti/cmd/collection/pack"
	"searchengine3090ti/kitex_gen/collectionModel"
)

type MGetColltService struct {
	ctx context.Context
}

func NewMGetUserService(ctx context.Context) *MGetColltService {
	return &MGetColltService{ctx: ctx}
}

func (m *MGetColltService) MGetColletction(req *collectionModel.MGetColltResquest) ([]*collectionModel.Collection, error) {
	collts, err := db.MGetColletction(m.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.Collections(collts), nil
}
