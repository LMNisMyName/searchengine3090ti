package handlers

import (
	"net/http"
	"searchengine3090ti/pkg/errno"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}

//Search struct
//搜索请求
type QueryParam struct {
	QueryText  string `json:"querytext"`  //用户在搜索框中输入的要查询的关键词
	FilterText string `json:"filtertext"` //用户请求过滤的关键词
	Page       int64  `json:"page"`       //用户请求的页码
	Limit      int64  `json:"limit"`      //每页显示的请求数目
	Order      int64  `json:"order"`      //排序方式
}
type imgQueryParam struct {
	QueryText  string `json:"querytext"`  //用户在搜索框中输入的要查询的关键词
	FilterText string `json:"filtertext"` //用户请求过滤的关键词
	Page       int64  `json:"page"`       //用户请求的页码
	Limit      int64  `json:"limit"`      //每页显示的请求数目
	Order      int64  `json:"order"`      //排序方式
}
type wd2imgQueryParam struct {
	QueryText  string `json:"querytext"`  //用户在搜索框中输入的要查询的关键词
	FilterText string `json:"filtertext"` //用户请求过滤的关键词
	Page       int64  `json:"page"`       //用户请求的页码
	Limit      int64  `json:"limit"`      //每页显示的请求数目
	Order      int64  `json:"order"`      //排序方式
}

//添加索引请求
type AddParam struct {
	Id   int64  `json:"id"`
	Text string `json:"text"`
	Url  string `json:"url"`
}

//相关搜索请求
type RelatedSearchParam struct {
	QueryText string `json:"querytext"`
}

//User struct
//用户创建请求
type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

//Collection struct
//收藏夹创建请求
type CreateColltParam struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
}

//获取收藏夹请求
type GetColltParam struct {
	UserID  int64 `json:"user_id"`
	ColltID int64 `json:"collt_id"`
}

//获取用户收藏夹请求
type MGetColltParam struct {
	UserID int64 `json:"user_id"`
}

//删除收藏夹
type DeleteColltParam struct {
	UserID  int64 `json:"user_id"`
	ColltID int64 `json:"collt_id"`
}

//添加收藏
type AddEntryParam struct {
	UserID   int64 `json:"user_id"`
	ColltID  int64 `json:"collt_id"`
	NewEntry int64 `json:"entry"`
}

//删除收藏
type DeleteEntryParam struct {
	UserID  int64 `json:"user_id"`
	ColltID int64 `json:"collt_id"`
	Entry   int64 `json:"entry"`
}

//设置收藏夹名
type SetNameParam struct {
	UserID  int64  `json:"user_id"`
	ColltID int64  `json:"collt_id"`
	NewName string `json:"newname"`
}
