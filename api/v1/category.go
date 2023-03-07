package v1

import (
	"gin-blog/model"
	"gin-blog/utils/messages"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddCateGory 添加分类
func AddCateGory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == messages.SUCCSE {
		model.CreateCategory(&data)
	}
	if code == messages.ERROR_USERNAME_USED {
		code = messages.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": messages.GetErrMsg(code),
	})
}

// GetCategory 查询单个分类
func GetCategory(c *gin.Context) {
	data := model.GetCategory(c.Param("name"))
	code = messages.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": messages.GetErrMsg(code),
	})
}

// GetCategories 查询分类列表
func GetCategories(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total := model.GetCategories(pageSize, pageNum)
	code = messages.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": messages.GetErrMsg(code),
	})
}

// EditCategory 编辑分类
func EditCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == messages.SUCCSE {
		model.EditCategoryInfo(int(data.ID), &data)
	}
	if code == messages.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": messages.GetErrMsg(code),
	})
}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	code = model.DeleteCategory(c.Param("name"))
	if code == messages.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": messages.GetErrMsg(code),
	})
}
