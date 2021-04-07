package model

import (
	"blog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Category	Category	`gorm:"foreignkey:Cid"`
	Title		string	`gorm:type:varchar(100);not null" json:"title"`
	Cid			int		`gorm:type:int;not null" json:"cid"`
	Desc		string	`gorm:type:varchar(200)" json:"desc"`
	Content		string	`gorm:type:longtext" json:"content"`
	Img			string	`gorm:type:varchar(100)" json:"img"`
}

// Add User
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// Find All Articals in Category
func GetCategoryArticles(id int, pageSize int, pageNum int) ([]Article, int) {
	var cateArtList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?",id).Find(&cateArtList).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST
	}
	return cateArtList, errmsg.SUCCESS
}

// Find single Article
func GetArticle(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ERROR_ART_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

// Find Article List
func GetArticles (pageSize int, pageNum int) ([]Article, int) {
	var articles []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return articles, errmsg.SUCCESS
}

// Edit Article
func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := db.Model(&article).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// Delete Article
func DeleteArticle(id int) int {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
