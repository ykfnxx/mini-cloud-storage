package model

import (
	"file-store/model/mysql"
	"time"
)

//用户表
type User struct {
	Id           int
	Password     string
	FileStoreId  int
	UserName     string
	RegisterTime time.Time
	ImagePath    string
}

//创建用户并新建文件仓库
func CreateUser(password, username, image string) {
	user := User{
		Password:     password,
		FileStoreId:  0,
		UserName:     username,
		RegisterTime: time.Now(),
		ImagePath:    image,
	}
	mysql.DB.Create(&user)

	fileStore := FileStore{
		UserId:      user.Id,
		CurrentSize: 0,
		MaxSize:     1048576,
	}
	mysql.DB.Create(&fileStore)

	user.FileStoreId = fileStore.Id
	mysql.DB.Save(&user)
}

//查询判断用户是否存在
func QueryUserExists(username string, password string) bool {
	var user User
	mysql.DB.Find(&user, "user_name = ?", username)
	if user.Id == 0 || user.Password != password {
		return false
	}
	return true
}

//根据openId查询用户
func GetUserInfo(openId interface{}) (user User) {
	mysql.DB.Find(&user, "user_name = ?", openId)
	return
}
