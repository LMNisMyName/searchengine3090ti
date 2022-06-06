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

func DeleteEntry(c *gin.Context) {
	var deleteEntryVar DeleteEntryParam
	colltId := c.Param("collt")
	if len(colltId) == 0 {
		SendResponse(c, errno.ParamErr, nil)
	}
	tmpI, err := strconv.Atoi(colltId)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	deleteEntryVar.ColltID = int64(tmpI)

	newEntry := c.PostForm("newentry")
	if len(newEntry) == 0 {
		SendResponse(c, errno.ParamErr, nil)
	}
	tmpI, err = strconv.Atoi(newEntry)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	deleteEntryVar.Entry = int32(tmpI)

	claim := jwt.ExtractClaims(c)
	userId := claim[constants.IdentityKey].(int64)
	deleteEntryVar.UserID = userId

	req := &collectionModel.DeleteEntryRequest{
		UserId:  deleteEntryVar.UserID,
		ColltId: deleteEntryVar.ColltID,
		Entry:   deleteEntryVar.Entry,
	}
	err = rpc.DeleteEntry(c, req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	SendResponse(c, errno.Success, nil)
}
