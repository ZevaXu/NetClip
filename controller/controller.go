package controller

import (
	"NetClip2/dao"
	"NetClip2/model"
	"log"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

//登录页面
func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

//添加用户
func AddUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{Username: username, Password: password}

	u := dao.Mgr.SearchUser(username)

	//数据库不存在用户，添加用户
	//存在，验证密码
	if u.Username == "" {
		dao.Mgr.AddUser(&user)
		log.Println("新增加用户:", username)
		url := "/" + username + "/" + password
		log.Println("url:", url)
		c.Redirect(302, url)
	} else {
		if u.Password == password {
			log.Println("Login:", username)
			url := "/" + username + "/" + password
			c.Redirect(302, url)
		} else {
			log.Println("密码不正确:", username)
			c.Redirect(302, "/login")
		}
	}
}

//进入用户首页
func UserIndex(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")

	u := dao.Mgr.SearchUser(username)

	if u.Username == "" {
		log.Println("用户不存在:", username)
		c.Redirect(302, "/login")
	} else {
		if u.Password == password {
			log.Println("密码正确，进入剪贴板", username)
			//这行改
			// c.HTML(200, "index.html", gin.H{
			// 	"username": username,
			// 	"password": password,
			// 	"text":     u.Text,
			// })
			c.HTML(200, "index.html", gin.H{
				"username":       username,
				"password":       password,
				"clipText":       u.Cliptext,
				"clipTextLength": utf8.RuneCountInString(u.Cliptext),
				"startAt":        u.CreatedAt.Format("01-02 15:04:05"),
				"updateAt":       u.UpdatedAt.Format("01-02 15:04:05"),
			})
			log.Println("统计:", u.CreatedAt, u.UpdatedAt, len(u.Cliptext))

		} else {
			log.Println("密码不正确:", username)
			c.Redirect(302, "/login")
		}
	}
}

//删除用户
func DeleteUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	log.Println("删除用户:", username)

	u := dao.Mgr.SearchUser(username)
	if u.Password == password {
		dao.Mgr.DeleteUser(username)
		log.Println("删除用户:", username)
		c.Redirect(302, "/login")
	} else {
		log.Println("密码不正确:", username)
		c.Redirect(302, "/login")
	}
}

//更新用户内容
func UpdateUserContent(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	clipText := strings.TrimSpace(c.PostForm("clipText"))
	clipExpire := c.PostForm("clipExpire")
	log.Println("更新用户:", username)
	log.Println("更新用户内容:", clipText)
	log.Println("过期时间：", clipExpire)

	var expire int
	switch clipExpire {
	case "一天":
		expire = 1
	case "一周":
		expire = 3
	case "三天":
		expire = 7
	case "一月":
		expire = 30
	}
	log.Println("过期时间：", expire)

	u := dao.Mgr.SearchUser(username)

	if u.Password == password {
		u.Cliptext = clipText
		u.Expired = expire
		dao.Mgr.UpdateUserContent(u)
		log.Printf("用户:%s，内容:%s", username, clipText)
		c.Redirect(302, "/"+username+"/"+password)
	} else {
		log.Println("密码不正确:", username)
		c.Redirect(302, "/login")
	}
}
