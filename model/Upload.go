package model

import (
	"context"
	"blog/utils"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"blog/utils/errmsg"
	"mime/multipart"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var QiniuServer = utils.QiniuServer

var ImgUrl = utils.QiniuServer

func UpLoadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:			&storage.ZoneHuadong,
		UseCdnDomains:	false,
		UseHTTPS:		false,
	}

	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	url := ImgUrl + ret.Key
	return url, errmsg.SUCCESS

}