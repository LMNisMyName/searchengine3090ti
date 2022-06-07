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

func GetCollection(c *gin.Context) {
	var getColltVar GetColltParam
	colltID := c.Param("collt")
	if len(colltID) == 0 {
		SendResponse(c, errno.ParamErr, nil)
	}
	tmpI, err := strconv.Atoi(colltID)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	getColltVar.ColltID = int64(tmpI)

	claim := jwt.ExtractClaims(c)
	UserId := claim[constants.IdentityKey].(int64)
	getColltVar.UserID = UserId

	req := &collectionModel.GetColltRequest{
		UserId:  getColltVar.UserID,
		ColltId: getColltVar.ColltID,
	}
	name, entry, err := rpc.GetCollection(c, req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	SendResponse(c, errno.Success,
		map[string]interface{}{
			constants.Name:  name,
			constants.Entry: entry,
		},
	)
}