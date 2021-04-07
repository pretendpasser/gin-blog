package v1

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"blog/model"
	"blog/utils/errmsg"
)

func Upload(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := model.UpLoadFile(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
		"url":url,		
	})
}