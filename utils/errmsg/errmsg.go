package errmsg

import (
	
)

const (
	SUCCESS	=	200
	ERROR	=	500

	// 1000... User model Error
	ERROR_USERNAME_USED		=	1001
	ERROR_PASSWORD_WRONG	=	1002
	ERROR_USER_NOT_EXIST	=	1003
	ERROR_TOKEN_EXIST		=	1004
	ERROR_TOKEN_RUNTIME		=	1005
	ERROR_TOKEN_WRONG		=	1006
	ERROR_TOKEN_TYPE_WRONG	=	1007
	ERROR_USER_NO_RIGNT		=	1008
	// 2000... Category model Error
	ERROR_CATENAME_USED		=	2001
	ERROR_CATE_NOT_EXIST	=	2002

	// 3000... Article model Error
	ERROR_ART_NOT_EXIST		=	3001
)

var CodeMsg = map[int]string{
	SUCCESS	:	"OK",
	ERROR	:	"FAIL",

	ERROR_USERNAME_USED		:	"UserName has been Exist!",
	ERROR_PASSWORD_WRONG	:	"Password incorrect!",
	ERROR_USER_NOT_EXIST	:	"User is not Exist!",
	ERROR_TOKEN_EXIST		:	"Token not Exist!",
	ERROR_TOKEN_RUNTIME		:	"Token has been Expired!",
	ERROR_TOKEN_WRONG		:	"Token incorrext!",		
	ERROR_TOKEN_TYPE_WRONG	:	"Token type incorrect!",
	ERROR_USER_NO_RIGNT		:	"Permission denied!",

	ERROR_CATENAME_USED		:	"Category has been Exist!",
	ERROR_CATE_NOT_EXIST	:	"Category is not Exist!",

	ERROR_ART_NOT_EXIST		:	"Article is not Exist!",

}

func GetErrMsg(code int) string {
	return CodeMsg[code]
}