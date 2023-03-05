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
		// 用户接口
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUsers)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:name", v1.DeleteUser)

		// 分类接口
		router.POST("category/add", v1.AddCateGory)
		router.GET("category/:name", v1.GetCategory)
		router.GET("categories", v1.GetCategories)
		router.PUT("category", v1.EditCategory)
		router.DELETE("category/:name", v1.DeleteCategory)

		// 文章模块的路由接口
		router.POST("article/add", v1.AddArticle)
		router.GET("article", v1.GetArt)
		router.PUT("article/:id", v1.EditArt)
		router.DELETE("article/:id", v1.DeleteArt)
		router.GET("article/list/:id", v1.GetCateArt)
		router.GET("article/info/:id", v1.GetArtInfo)
	}

	r.Run(utils.HttpPort)
}
