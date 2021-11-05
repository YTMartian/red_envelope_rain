package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"red_envelope/routers"
	"red_envelope/utils"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

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
