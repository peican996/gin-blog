package routes

import (
	"gin-blog/api/v1"
	"gin-blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	//auth := routers.Group("/api/v1")
	//{
	//	// 用户模块路由,后台管理路由接口
	//	auth.GET("user/:id", v1.GetUsers)
	//	auth.PUT("user/:id", v1.EditUser)
	//	auth.DELETE("user/:id", v1.DeleteUser)
	//
	//
	//	auth.GET("hello", func(context *gin.Context) {
	//		context.JSON(http.StatusOK, gin.H{
	//			"messages": "OK",
	//		})
	//	})
	//}

	// 前台展示页面接口
	router := r.Group("/api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUsers)
		router.GET("users", v1.GetUsers)
	}

	r.Run(utils.HttpPort)
}
