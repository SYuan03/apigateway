// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	demo "github.com/SYuan03/idlmanage/biz/handler/demo"
	
)

func main() {
	// 初始化数据库
	demo.InitDB()

	h := server.New(server.WithHostPorts("127.0.0.1:6666"))

	register(h)
	h.Spin()
}


