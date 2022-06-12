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

func AddEntry(c *gin.Context) {
	var addEntryVar AddEntryParam
	colltId := c.Param("collt")
	if len(colltId) == 0 {
		SendResponse(c, errno.ParamErr, nil)
	}
	tmpI, err := strconv.Atoi(colltId)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	addEntryVar.ColltID = int64(tmpI)

	newEntry := c.PostForm("newentry")
	if len(newEntry) == 0 {
		SendResponse(c, errno.ParamErr, nil)
	}
	tmpI, err = strconv.Atoi(newEntry)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	addEntryVar.NewEntry = int64(tmpI)

	claim := jwt.ExtractClaims(c)
	userId := int64(claim[constants.IdentityKey].(float64))
	addEntryVar.UserID = userId

	//check newEntry if exits
	reqC := &searchapi.FindIDRequest{
		Ids: []int64{addEntryVar.NewEntry},
	}
	isExist, err := rpc.FindID(c, reqC)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	if len(isExist) == 0 {
		SendResponse(c, errno.ParamErr, nil)
	}

	req := &collectionModel.AddEntryRequest{
		UserId:   addEntryVar.UserID,
		ColltId:  addEntryVar.ColltID,
		NewEntry: addEntryVar.NewEntry,
	}
	err = rpc.AddEntry(c, req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}

	SendResponse(c, errno.Success, nil)
}
