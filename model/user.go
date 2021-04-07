package model

import (
	"log"
	"blog/utils/errmsg"

	"encoding/base64"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	gorm.Model
	Username	string	`gorm:"type:varchar(20);not null" json:"username"`
	Password	string	`gorm:"type:varchar(20);not null" json:"password"`
	Role		int		`gorm:"type:int" json:"role"`
}

// Find if User Exist
func CheckUser(name string) int {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	 if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED	//1001
	 }
	 return errmsg.SUCCESS
}

// Add User
func CreateUser(data *User) int {
	//data.Password = ScryptPw(data.Password)
	//data.BeforeSave()
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// Find User List
func GetUsers (pageSize int, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil
	}
	return users
}

// Edit User
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// Delete User
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// Encrypt Password
func(u *User) BeforeSave() {
	u.Password = ScryptPw(u.Password)
}

func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{1,7,28,39,66,91,124,225}
	HashPw, err :=scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Println(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// Login Verify
func CheckLogin(username string, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 0 {
		return errmsg.ERROR_USER_NO_RIGNT
	}
	return errmsg.SUCCESS
}