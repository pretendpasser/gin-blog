package model

import (
	"github.com/jinzhu/gorm"
	"blog/utils/errmsg"
)

type Category struct {
	ID		uint 	`gorm:"primary_key;auto_increment" json:"id"`
	Name	string	`gorm:type:varchar(20);not null" json:"name"`
}

// Find if Category Exist
func CheckCategory(name string) int {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	 if category.ID > 0 {
		return errmsg.ERROR_CATENAME_USED	//1001
	 }
	 return errmsg.SUCCESS
}

// Add Category
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// Find All Articals in Category

// Find Category List
func GetCategorys(pageSize int, pageNum int) []Category {
	var categorys []Category
	err := db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&categorys).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil
	}
	return categorys
}

// Edit Category
func EditCategory(id int, data *Category) int {
	var category Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&category).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// Delete Category
func DeleteCategory(id int) int {
	var category Category
	err := db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}