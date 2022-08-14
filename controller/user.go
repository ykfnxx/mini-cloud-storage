package controller

import (
	"file-store/lib"
	"file-store/model"
	"file-store/util"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PrivateInfo struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
}

type LoginInfo struct {
	Username string
	Password string
}

//登录成功获取的QQ用户信息
type QUserInfo struct {
	Nickname    string
	FigureUrlQQ string `json:"figureurl_qq"`
}

func Login_pwd(c *gin.Context) {
	// fmt.Println(c.PostForm("username"))
	// fmt.Println(c.PostForm("password"))

	username := c.PostForm("username")
	password := c.PostForm("password")

	hashToken := util.EncodeMd5("token" + string(time.Now().Unix()) + username)

	//存入redis
	if err := lib.SetKey(hashToken, username, 24*3600); err != nil {
		fmt.Println("Redis Set Err:", err.Error())
		return
	}
	//设置cookie
	c.SetCookie("Token", hashToken, 3600*24, "/", "172.19.43.107", false, true)
	if ok := model.QueryUserExists(username, password); ok { //用户存在直接登录
		//登录成功重定向到首页
		c.Redirect(http.StatusMovedPermanently, "/cloud/index")
	} else {
		model.CreateUser(password, username, "nil")
		//登录成功重定向到首页
		c.Redirect(http.StatusMovedPermanently, "/cloud/index")
	}
}

//登录页
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
	// log.Fatalln(c.PostForm("username"))
}

//退出登录
func Logout(c *gin.Context) {
	token, err := c.Cookie("Token")
	if err != nil {
		fmt.Println("cookie", err.Error())
	}

	if err := lib.DelKey(token); err != nil {
		fmt.Println("Del Redis Err:", err.Error())
	}

	c.SetCookie("Token", "", 0, "/", "172.19.43.107", false, false)
	c.Redirect(http.StatusFound, "/")
}
