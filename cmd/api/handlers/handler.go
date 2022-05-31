package handlers

import (
	"net/http"
	searchapi "searchengine3090ti/kitex_gen/SearchApi"
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
//TODO
//搜索请求
type QueryParam struct {
	QueryText  string `json:"querytext"`  //用户在搜索框中输入的要查询的关键词
	FilterText string `json:"filtertext"` //用户请求过滤的关键词
	Page       int32  `json:"page"`       //用户请求的页码
	Limit      int32  `json:"limit"`      //每页显示的请求数目
	Order      int32  `json:"order"`      //排序方式
}
type SearchQuery struct {
	Time      float64                 `json:"time"`
	Total     int32                   `json:"total"`
	PageCount int32                   `json:"pagecount"`
	Page      int32                   `json:"page"`
	Limit     int32                   `json:"limit"`
	Contents  []*searchapi.AddRequest `json:"contents"`
}

//添加索引请求
type AddParam struct {
	Id   int32  `json:"id"`
	Text string `json:"text"`
	Url  string `json:"url"`
}

//相关搜索请求
type RelatedSearchParam struct {
	QueryText string `json:"querytext"`
}
type RelatedQuery struct {
	RelatedText []string `json:"relatedtext"`
}

//User struct
type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
