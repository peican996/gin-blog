package routes

import (
	"gin-blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	routers := gin.Default()

	router := routers.Group("/api/v1")
	{
		router.GET("hello", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "OK",
			})
		})
	}

	routers.Run()
}
