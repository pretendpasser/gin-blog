package v1

import (
	"strconv"
	"net/http"
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
)

// Add User
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":	code,
		"data":		data,
		"message":	errmsg.GetErrMsg(code),
	})
}

// Find single User


// Find User list
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"status":	code,
		"data":		data,
		"message":	errmsg.GetErrMsg(code),
	})
}

// Edit User
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK,gin.H{
		"status":	code,
		"message":	errmsg.GetErrMsg(code),
	})
}

// Delete User
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":	code,
		"message":	errmsg.GetErrMsg(code),		
	})
}
