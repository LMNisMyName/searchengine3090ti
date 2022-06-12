package handlers

import (
	"context"
	"searchengine3090ti/cmd/api/rpc"
	"searchengine3090ti/kitex_gen/userModel"
	"searchengine3090ti/pkg/errno"

	"github.com/gin-gonic/gin"
)

//注册直接接口
func Register(c *gin.Context) {
	var registerVar UserParam
	if err := c.ShouldBind(&registerVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	err := rpc.CreateUser(context.Background(), &userModel.CreateUserRequest{
		UserName: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	} else {
		SendResponse(c, errno.Success, nil)
	}
}
