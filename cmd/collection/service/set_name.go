package service

import (
	"context"
	"searchengine3090ti/cmd/collection/dal/db"
	"searchengine3090ti/kitex_gen/collectionModel"
	"searchengine3090ti/pkg/errno"
)

type SetNameService struct {
	ctx context.Context
}

func NewSetNameService(ctx context.Context) *SetNameService {
	return &SetNameService{ctx: ctx}
}

func (s *SetNameService) SetName(req *collectionModel.SetNameRequest) error {
	collts, err := db.GetColletction(s.ctx, req.UserId, req.ColltId)
	if err != nil {
		return err
	}
	if len(collts) == 0 {
		return errno.CollectionNotExitErr
	}
	for _, collt := range collts {
		if collt.Name == req.NewName {
			return errno.CollectionAlreadyExistErr
		}
	}
	return db.SetName(s.ctx, req.UserId, req.ColltId, req.NewName)
}
