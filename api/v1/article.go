package v1

import (
	"strconv"
	"net/http"
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
)

var code int

// Add Article
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":	code,
		"data":		data,
		"message":	errmsg.GetErrMsg(code),
	})
}

// Find All Articals in Category
func GetCategoryArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := model.GetCategoryArticles(id, pageSize, pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status":	code,
		"data":		data,
		"message":	errmsg.GetErrMsg(code),
	})
}

// Find single Article
func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArticle(id)
	c.JSON(http.StatusOK,gin.H{
		"status":	code,
		"data":		data,
		"message":	errmsg.GetErrMsg(code),
	})
}

// Find Article list
func GetArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := model.GetArticles(pageSize, pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status":	code,
		"data":		data,
		"message":	errmsg.GetErrMsg(code),
	})
}

// Edit Article
func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.EditArticle(id, &data)
	c.JSON(http.StatusOK,gin.H{
		"status":	code,
		"message":	errmsg.GetErrMsg(code),
	})
}

// Delete Article
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":	code,
		"message":	errmsg.GetErrMsg(code),		
	})
}

//