package main

import (
	"fmt"
	"ginDemo/apis"
	"ginDemo/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

var db *gorm.DB
var err error

func main() {
	defer models.DbOrm.Close()
	g := gin.Default()
	apiRouter(g)
	//webRouter(g)

	g.Run()
}

func apiRouter(g *gin.Engine) {
	//apiGroup := g.Group("/api")
	apiGroup := g.Group("/api", apis.BaseVerify())

	//登陆
	apiGroup.POST("/login", apis.Login())

	//注册
	apiGroup.POST("/register", apis.Register())

	//修改
	apiGroup.POST("/modify", apis.Modify())

	apiGroup.POST("/articleList", apis.ArticleList())

	apiGroup.POST("/addArticle", apis.AddArticle())
}

func webRouter(g *gin.Engine) {
	//g.Static("/static", "./static")
	//g.LoadHTMLGlob("views/*")

	//首页页面
	g.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//注册页面
	g.GET("/reg", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	//注册逻辑
	g.POST("/reg", func(c *gin.Context) {
		name := c.PostForm("name")
		pwd := c.PostForm("pwd")
		fmt.Println("name =", name, "pwd =", pwd)

		//写入数据库, 并返回登录界面或者首页
		c.HTML(200, "index.html", gin.H{
			"name": name,
		})
	})
}
