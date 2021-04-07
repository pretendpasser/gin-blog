package v1

import (
	"net/http"
	"blog/midware"
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data model.User
	var token string
	c.ShouldBindJSON(&data)
	code := model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		token,code = midware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
		"token":token,
	})
}