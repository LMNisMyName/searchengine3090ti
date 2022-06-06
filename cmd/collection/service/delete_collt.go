package service

import (
	"context"
	"searchengine3090ti/cmd/collection/dal/db"
	"searchengine3090ti/kitex_gen/collectionModel"
	"searchengine3090ti/pkg/errno"
)

type DeleteColltService struct {
	ctx context.Context
}

func NewDeleteColltService(ctx context.Context) *DeleteColltService {
	return &DeleteColltService{ctx: ctx}
}

func (d *DeleteColltService) DeleteCollection(req *collectionModel.DeleteColltRequest) error {
	collt, err := db.GetColletction(d.ctx, req.UserId, req.ColltId)
	if err != nil {
		return err
	}
	if len(collt) == 0 {
		return errno.CollectionNotExitErr
	}
	return db.DeleteCollection(d.ctx, req.UserId, req.ColltId)
}
