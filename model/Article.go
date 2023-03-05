package model

import (
	"gin-blog/utils/messages"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// CreateArt 新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return messages.ERROR // 500
	}
	return messages.SUCCSE
}

// GetCateArt 查询分类下的所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int) {
	var cateArtList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid =?", id).Find(&cateArtList).Error
	if err != nil {
		return nil, messages.ERROR_CATE_NOT_EXIST
	}
	return cateArtList, messages.SUCCSE
}

// GetArtInfo 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, messages.ERROR_ART_NOT_EXIST
	}
	return art, messages.SUCCSE
}

// GetArt 查询文章列表
func GetArt(pageSize int, pageNum int) ([]Article, int) {
	var articleList []Article
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, messages.ERROR
	}
	return articleList, messages.SUCCSE
}

// EditArt 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&art).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return messages.ERROR
	}
	return messages.SUCCSE
}

// DeleteArt 删除文章
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ? ", id).Delete(&art).Error
	if err != nil {
		return messages.ERROR
	}
	return messages.SUCCSE
}
