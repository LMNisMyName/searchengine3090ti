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

func DeleteCollection(c *gin.Context) {
	var deleteColltVar DeleteColltParam
	colltId := c.Param("collt")
	if len(colltId) == 0 {
		SendResponse(c, errno.ParamErr, nil)
	}
	tmpI, err := strconv.Atoi(colltId)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	deleteColltVar.ColltID = int64(tmpI)

	claim := jwt.ExtractClaims(c)
	UserId := claim[constants.IdentityKey].(int64)
	deleteColltVar.UserID = UserId

	req := &collectionModel.DeleteColltRequest{
		UserId:  deleteColltVar.UserID,
		ColltId: deleteColltVar.ColltID,
	}
	err = rpc.DeleteCollection(c, req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	SendResponse(c, errno.Success, nil)
}
