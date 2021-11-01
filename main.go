package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"red_envelope/src/routers"
)

func main() {
	r := gin.Default()

	routers.ApiRoutersInit(r)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("run error.")
	}
}
