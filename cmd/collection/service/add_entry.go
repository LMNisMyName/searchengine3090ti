package service

import (
	"context"
	"searchengine3090ti/cmd/collection/dal/db"
	"searchengine3090ti/kitex_gen/collectionModel"
	"searchengine3090ti/pkg/errno"
)

type AddEntryService struct {
	ctx context.Context
}

func NewAddEntryService(ctx context.Context) *AddEntryService {
	return &AddEntryService{ctx: ctx}
}

func (a *AddEntryService) AddEntry(req *collectionModel.AddEntryRequest) error {
	collts, err := db.GetColletction(a.ctx, req.UserId, req.ColltId)
	if err != nil {
		return err
	}
	if len(collts) == 0 {
		return errno.CollectionNotExitErr
	}
	return db.AddEntry(a.ctx, req.UserId, req.ColltId, req.NewEntry)
}
