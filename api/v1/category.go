package v1

import (
	"strconv"
	"net/http"
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
)

// Add Category
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		code = errmsg.ERROR_CATENAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":	code,
		"data":		data,
		"message":	errmsg.GetErrMsg(code),
	})
}
// Find All Articals in Category

// Find Category list
func GetCategorys(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetCategorys(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"status":	code,
		"data":		data,
		"message":	errmsg.GetErrMsg(code),
	})
}

// Edit Category
func EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.EditCategory(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK,gin.H{
		"status":	code,
		"message":	errmsg.GetErrMsg(code),
	})
}

// Delete Category
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":	code,
		"message":	errmsg.GetErrMsg(code),		
	})
}
