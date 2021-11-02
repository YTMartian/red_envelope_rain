package utils

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//全局变量
var SnowflakeNode *snowflake.Node
var DB *gorm.DB

func Init() (err error) {
	//初始化雪花算法
	SnowflakeNode, err = snowflake.NewNode(1)
	if err != nil {
		return err
	}
	//初始化数据库
	return
}
