package model

import (
	"gin-blog/utils/messages"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategory 分类校验
func CheckCategory(name string) (code int) {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return messages.ERROR_USERNAME_USED
	}
	return messages.SUCCSE
}

// CreateCategory 创建分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return messages.ERROR
	}
	return messages.SUCCSE
}

// GetCategory 获取单个分类
func GetCategory(name string) Category {
	var category Category
	err = db.Where("name = ?", name).Find(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return category
	}
	return category
}

// GetCategories 获取用户信息
func GetCategories(pageSize int, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

// EditCategoryInfo 用户信息编辑
func EditCategoryInfo(id int, data *Category) int {
	var category Category
	var maps = make(map[string]any)
	maps["name"] = data.Name
	err = db.Model(&category).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return messages.ERROR
	}
	return messages.SUCCSE
}

// DeleteCategory 删除用户
func DeleteCategory(name string) int {
	var category Category
	err = db.Where("name = ?", name).Delete(&category).Error
	if err != nil {
		return messages.ERROR
	}
	return messages.SUCCSE
}
