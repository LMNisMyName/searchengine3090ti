package pack

import (
	"errors"
	"searchengine3090ti/kitex_gen/userModel"
	"searchengine3090ti/pkg/errno"
	"time"
)

func BuildBaseResp(err error) *userModel.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *userModel.BaseResp {
	return &userModel.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
