package service

import (
	"context"
	"searchengine3090ti/cmd/user/dal/db"
	"searchengine3090ti/cmd/user/pack"
	"searchengine3090ti/kitex_gen/userModel"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *MGetUserService) MGetUser(req *userModel.MGetUserRequest) ([]*userModel.User, error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return pack.Users(modelUsers), nil
}
