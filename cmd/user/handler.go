package main

import (
	"context"
	"searchengine3090ti/cmd/user/pack"
	"searchengine3090ti/cmd/user/service"
	"searchengine3090ti/kitex_gen/userModel"
	"searchengine3090ti/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
//注册
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userModel.CreateUserRequest) (resp *userModel.CreateUserResponse, err error) {
	resp = new(userModel.CreateUserResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
//批量获取用户信息
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *userModel.MGetUserRequest) (resp *userModel.MGetUserResponse, err error) {
	resp = new(userModel.MGetUserResponse)

	if len(req.UserIds) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewMGetUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Users = users
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
//校验用户
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userModel.CheckUserRequest) (resp *userModel.CheckUserResponse, err error) {
	resp = new(userModel.CheckUserResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
