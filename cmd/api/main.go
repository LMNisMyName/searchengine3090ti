package main

import (
	"context"
	"net/http"
	"searchengine3090ti/cmd/api/handlers"
	"searchengine3090ti/cmd/api/idgen"
	"searchengine3090ti/cmd/api/rpc"
	"searchengine3090ti/kitex_gen/userModel"
	"searchengine3090ti/pkg/constants"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.InitRPC()
	idgen.Init()
}

func main() {
	Init()
	r := gin.New()
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		SendCookie: true,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVar handlers.UserParam
			if err := c.ShouldBind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			return rpc.CheckUser(context.Background(), &userModel.CheckUserRequest{UserName: loginVar.UserName, Password: loginVar.PassWord})
		},
		// 验证用户
		// Authorizator: func(data interface{}, c *gin.Context) bool {
		// 	if v, ok := data.(*User); ok && v.UserName == "admin" {
		// 	  return true
		// 	}

		// 	return false
		// },
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	user1 := r.Group("/user")
	user1.POST("/login", authMiddleware.LoginHandler) //登录
	user1.POST("/register", handlers.Register)        //注册
	user1.Use(authMiddleware.MiddlewareFunc())
	{
		user1.POST("/refresh", authMiddleware.RefreshHandler) //续签
		user1.POST("/logout", authMiddleware.LogoutHandler)   //注销
	}

	collection1 := r.Group("/collection")
	collection1.Use(authMiddleware.MiddlewareFunc())
	{
		collection1.GET("/", handlers.MGetCollection)
		collection1.GET("/:collt", handlers.GetCollection)
		collection1.POST("/create", handlers.CreateCollection)
		collection1.GET("/delete/:collt", handlers.DeleteCollection)
		collection1.POST("/:collt/setname", handlers.SetName)
		collection1.POST("/:collt/add", handlers.AddEntry)
		collection1.GET("/:collt/delete", handlers.DeleteEntry)
	}

	search1 := r.Group("/search")
	search1.GET("/add", handlers.Add)
	search1.POST("/query", handlers.Query)
	search1.POST("/imgquery", handlers.ImgQuery)
	search1.POST("/wd2imgquery", handlers.Wd2ImgQuery)
	search1.GET("/relatedsearch", handlers.RelatedSearch)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
	//https
	// if err := http.ListenAndServeTLS(
	// 	":8080",
	// 	constants.certCrtPath,
	// 	constants.certKeyPath,
	// 	r); err != nil {
	// 	klog.Fatal(err)
	// }
}
