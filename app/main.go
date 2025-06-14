package main

import (
	"base"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// 初始化配置
	base.InitConfig("./conf.yml")

	// http服务
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello go-lift")
	})

	// 启动服务
	http.ListenAndServe(":"+strconv.Itoa(base.Conf.Http.Port), nil)
}
