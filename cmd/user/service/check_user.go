package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"searchengine3090ti/cmd/user/dal/db"
	"searchengine3090ti/kitex_gen/userModel"
	"searchengine3090ti/pkg/errno"
)

type CheckUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

// CheckUser check user info
func (s *CheckUserService) CheckUser(req *userModel.CheckUserRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.UserName
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.UserNotExistErr
	}
	u := users[0]
	if u.Password != passWord {
		return 0, errno.LoginErr
	}
	return int64(u.ID), nil
}
