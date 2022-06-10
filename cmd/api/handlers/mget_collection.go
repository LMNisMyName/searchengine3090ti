package handlers

import (
	"searchengine3090ti/cmd/api/rpc"
	"searchengine3090ti/kitex_gen/collectionModel"
	"searchengine3090ti/pkg/constants"
	"searchengine3090ti/pkg/errno"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func MGetCollection(c *gin.Context) {
	var mgetColltVar MGetColltParam
	claim := jwt.ExtractClaims(c)
	UserId := int64(claim[constants.IdentityKey].(float64))
	mgetColltVar.UserID = UserId

	req := &collectionModel.MGetColltResquest{
		UserId: UserId,
	}
	names, collts, err := rpc.MGetCollection(c, req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	} else {
		SendResponse(c, errno.Success,
			map[string]interface{}{
				constants.Name:    names,
				constants.ColltID: collts,
			},
		)
	}
}
