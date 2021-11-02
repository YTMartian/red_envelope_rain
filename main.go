package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"red_envelope/routers"
	"red_envelope/utils"
)

func main() {
	var err error
	err = utils.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	r := gin.Default()
	routers.ApiRoutersInit(r)
	err = r.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
