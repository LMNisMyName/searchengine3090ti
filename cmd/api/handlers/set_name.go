package handlers

import (
	"searchengine3090ti/cmd/api/rpc"
	"searchengine3090ti/kitex_gen/collectionModel"
	"searchengine3090ti/pkg/constants"
	"searchengine3090ti/pkg/errno"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func SetName(c *gin.Context) {
	var setNameVar SetNameParam
	colltId := c.Param("collt")
	if len(colltId) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	tmpI, err := strconv.Atoi(colltId)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	setNameVar.ColltID = int64(tmpI)

	newname := c.PostForm("newname")
	if len(newname) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	setNameVar.NewName = newname

	claim := jwt.ExtractClaims(c)
	UserId := int64(claim[constants.IdentityKey].(float64))
	setNameVar.UserID = UserId

	req := &collectionModel.SetNameRequest{
		UserId:  setNameVar.UserID,
		ColltId: setNameVar.ColltID,
		NewName: setNameVar.NewName,
	}
	err = rpc.SetName(c, req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	} else {
		SendResponse(c, errno.Success, nil)
	}
}
