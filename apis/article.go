package apis

import (
	"fmt"
	"ginDemo/models"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/mapstructure"
	"net/http"
)

//获取作品列表
func ArticleList() gin.HandlerFunc {
	return func(c *gin.Context) {
		articles := make([]models.Article, 10, 10)
		err := models.ArticleList(&articles)
		if err != nil {
			c.JSON(http.StatusOK, BaseRes{400, fmt.Sprintf("%v", err)})
			return
		}
		c.JSON(http.StatusOK, BodyRes{BaseRes{200, "Success"}, &articles})
	}
}

/*添加一个作品
请求参数:
	Desc string 	`gorm:"not null;defult:''"`
	ImgUrl string 	`gorm:"not null;defult:0"`
	Source string 	`gorm:"not null;defult:0"` //来自哪个网站
	Type int 		`gorm:"not null;defult:0"` //什么类型的内容, 0, 图片+文字 1, gif+文字 2, 视频+文字 3, 文章
*/
func AddArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.MustGet(ALL_PARAMES).(gin.H)
		var article models.Article
		err := mapstructure.Decode(p, &article)
		if err != nil {
			c.JSON(http.StatusOK, BaseRes{400, fmt.Sprintf("%v", err)})
			return
		}
		err = models.DbOrm.Create(&article).Error
		if err != nil {
			panic(err)
			c.JSON(http.StatusOK, BaseRes{400, fmt.Sprintf("%v", err)})
			return
		}
		c.JSON(http.StatusOK, BodyRes{BaseRes{200, "Success"}, &article})
	}
}
