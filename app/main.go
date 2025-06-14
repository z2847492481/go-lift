package main

import (
	"base"
	"fmt"
)

func main() {
	// 初始化配置
	base.InitConfig("./conf.yml")
	fmt.Printf("Http端口号:%d\n", base.Conf.Http.Port)
	fmt.Printf("数据库地址:%s\n", base.Conf.Db.Path)
}
