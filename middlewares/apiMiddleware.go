package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//中间件，Group、GET、POST等方法中可以加入多个函数用于做中间件
type ApiMiddleware struct{}

//模拟判断用户是否登陆的中间件
func (con ApiMiddleware) ApiMiddleware(c *gin.Context) {
	logined := false
	if logined {
		c.Next() //继续执行后面的函数
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "not login",
		})
		c.Abort() //中断执行
	}
}
