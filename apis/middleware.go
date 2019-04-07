package apis

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const (
	//秘钥
	SECRET_KEY string = "HAFQWE128FAJ14F0A"
	//请过中间件取值后重新复制的参数key, 后续的都可以用这个key取出请求参数
	ALL_PARAMES string = "parameall"
)

type BaseRes struct {
	ResCode int
	ResMsg  string
}

type BodyRes struct {
	BaseRes
	Body interface{}
}

type BaseReq struct {
	Body      gin.H
	Sign      string
	TimeStamp uint64
}

//基础的api验证
func BaseVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取原始请求参数
		var req = new(BaseReq)
		err := c.ShouldBindJSON(req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, BaseRes{400, "参数解析失败!"})
			return
		}
		fmt.Println("middle BaseVerify", req)
		m := req.Body

		//以下是验证sign的流程, md5(参数form& + 时间戳 + key)
		//1,把原始参数map按key排序
		params := make([]string, 0, len(m))
		for key, _ := range m {
			params = append(params, key)
		}
		sort.Strings(params)

		//2,把原始参数map按key排序后生成form格式参数  k1=v1&k2=v2&...
		var bulder = new(strings.Builder)
		for _, key := range params {
			v, _ := m[key].(string)
			bulder.WriteString(key)
			bulder.WriteString("=")
			bulder.WriteString(v)
			bulder.WriteString("&")
		}
		bulder.WriteString(strconv.FormatUint(req.TimeStamp, 10))
		bulder.WriteString(SECRET_KEY)
		str := bulder.String()
		//fmt.Println("sign ori : ", str)

		//3,MD5指定的字符串
		hash := md5.New()
		hash.Write([]byte(str))
		sign := hex.EncodeToString(hash.Sum(nil))

		fmt.Println("sign =", sign)
		//4,对比sign
		if strings.Compare(sign, req.Sign) != 0 { //sign对不上, 中断请求
			c.AbortWithStatusJSON(http.StatusOK, BaseRes{400, "签名错误!"})
			return
		}

		c.Set(ALL_PARAMES, m)
		c.Next()
	}
}
