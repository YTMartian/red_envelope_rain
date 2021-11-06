package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"red_envelope/configure"
	"red_envelope/middlewares"
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
	r.Use(middlewares.RecoveryMiddleWare())
	gin.SetMode(configure.GIN_MODE)
	routers.ApiRoutersInit(r)
	//s := &http.Server{
	//	Addr:           configure.LISTEN_PORT,
	//	Handler:        r,
	//	ReadTimeout:    configure.ReadTimeout,
	//	WriteTimeout:   configure.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//err = s.ListenAndServe()

	err = r.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
