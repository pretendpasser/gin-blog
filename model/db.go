package model

import (
	"fmt"
	"log"
	"time"
	"blog/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db	*gorm.DB
	err	error
)

func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
		))
	if err != nil {
		log.Println("Attach database fail!", err)
	}

	db.SingularTable(true)
	db.AutoMigrate(&User{}, &Article{}, &Category{})

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(10*time.Second)
	
	//db.Close()
}