package apis

import (
	"fmt"
	"ginDemo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//登陆接口
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取原始请求参数
		req := c.MustGet(ALL_PARAMES).(gin.H)
		name, _ := req["name"].(string)
		pwd, _ := req["pwd"].(string)

		var user models.User
		user.Name = name
		ok := models.FindUserByName(&user)
		if !ok {
			c.JSON(http.StatusOK, BaseRes{400, "用户不存在"})
			return
		}
		fmt.Println("登陆", user)
		if strings.Compare(user.Pwd, pwd) != 0 {
			c.JSON(http.StatusOK, BaseRes{400, "密码错误"})
			return
		}
		c.JSON(http.StatusOK, BaseRes{200, "登陆成功"})
	}
}

//注册
func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.MustGet(ALL_PARAMES).(gin.H)
		name, _ := req["name"].(string)
		pwd, _ := req["pwd"].(string)
		//todo 此处需要验证账号密码是否合法, 去除前后空格, 然后检查中间有没有空格,
		// 然后检查是不是英文或者数字或者特殊符号等, 可以写一个全局正则表达式变量验证

		var user models.User
		user.Name = name
		ok := models.FindUserByName(&user)
		if ok {
			c.JSON(http.StatusOK, BaseRes{400, "用户已经存在, 请返回登陆"})
			return
		}

		user.Name = name
		user.Pwd = pwd
		create := models.DbOrm.Create(&user)
		if create.Error != nil {
			c.JSON(http.StatusOK, BaseRes{400, "注册失败!"})
			return
		}
		c.JSON(200, BaseRes{200, "注册成功!"})
	}
}

//修改
func Modify() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.MustGet(ALL_PARAMES).(gin.H)
		name, _ := req["name"].(string)
		pwd, _ := req["pwd"].(string)
		nickname, _ := req["nickname"].(string)

		var user models.User
		user.Name = name
		ok := models.FindUserByName(&user)
		if !ok {
			c.JSON(http.StatusOK, BaseRes{400, "用户不存在"})
			return
		}
		if strings.Compare(user.Pwd, pwd) != 0 {
			c.JSON(200, BaseRes{400, "密码错误"})
			return
		}

		user.NickName = nickname
		update := models.DbOrm.Model(&user).Update("NickName", nickname)
		if update.Error != nil {
			fmt.Println(update.Error)
			c.JSON(http.StatusOK, BaseRes{400, "修改失败!" + fmt.Sprintf("%v", update.Error)})
			return
		}
		c.JSON(200, BaseRes{200, "修改成功!"})
	}
}
