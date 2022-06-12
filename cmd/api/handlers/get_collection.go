package handlers

import (
	"searchengine3090ti/cmd/api/rpc"
	searchapi "searchengine3090ti/kitex_gen/SearchApi"
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
		return
	}
	tmpI, err := strconv.Atoi(colltID)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	getColltVar.ColltID = int64(tmpI)

	claim := jwt.ExtractClaims(c)
	UserId := int64(claim[constants.IdentityKey].(float64))
	getColltVar.UserID = UserId

	req := &collectionModel.GetColltRequest{
		UserId:  getColltVar.UserID,
		ColltId: getColltVar.ColltID,
	}
	name, entry, err := rpc.GetCollection(c, req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	// rpc to get actuall entry title
	reqF := &searchapi.FindIDRequest{
		Ids: entry,
	}
	actEntry, err := rpc.FindID(c, reqF)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	} else {
		SendResponse(c, errno.Success,
			map[string]interface{}{
				constants.Name:  name,
				constants.Entry: actEntry,
			},
		)
	}
}
