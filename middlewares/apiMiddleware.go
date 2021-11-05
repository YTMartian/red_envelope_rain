package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"red_envelope/utils"
)

// ApiMiddleware 中间件，Group、GET、POST等方法中可以加入多个函数用于做中间件
type ApiMiddleware struct{}

// ApiMiddleware 模拟判断用户是否登陆的中间件
func (con ApiMiddleware) ApiMiddleware(c *gin.Context) {
	login := true
	if login {
		c.Next() //继续执行后面的函数
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": utils.CODE_NOT_LOGIN_ERROR,
			"msg":  utils.MSG_NOT_LOGIN_ERROR,
		})
		c.Abort() //中断执行
	}
}
