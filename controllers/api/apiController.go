package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//把方法挂在结构体下面,方便以后继承
type ApiController struct{}

func (con ApiController) Snatch(c *gin.Context) {
	c.String(http.StatusOK, "fuck")
}

func (con ApiController) Open(c *gin.Context) {

}

func (con ApiController) GetWalletList(c *gin.Context) {

}
