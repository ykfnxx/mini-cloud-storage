package controller

import (
	"file-store/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	openId, _ := c.Get("openId")
	//获取用户信息
	user := model.GetUserInfo(openId)
	//获取用户仓库信息
	userFileStore := model.GetUserFileStore(user.Id)
	//获取用户文件数量
	fileCount := model.GetUserFileCount(user.FileStoreId)
	//获取用户文件夹数量
	fileFolderCount := model.GetUserFileFolderCount(user.FileStoreId)
	//获取用户文件使用明细数量
	fileDetailUse := model.GetFileDetailUse(user.FileStoreId)
	fmt.Println(openId)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"user":            user,
		"currIndex":       "active",
		"userFileStore":   userFileStore,
		"fileCount":       fileCount,
		"fileFolderCount": fileFolderCount,
		"fileDetailUse":   fileDetailUse,
	})
}
