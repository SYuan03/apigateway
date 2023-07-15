package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"net/http"
	"nju/apigw/kitex_gen/demo/studentservice"
)

func main() {
	//TODO 注册
	handler := &StudentServiceImpl{}
	handler.InitDB()
	addr, _ := net.ResolveTCPAddr("tcp", ":9990")
	svr := studentservice.NewServer(handler, server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
	go http.ListenAndServe("localhost:8880", nil)
}
