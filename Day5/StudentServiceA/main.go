package main

import (
	"log"
	"net"

	demo "github.com/SYuan03/Day3/KitexServer/kitex_gen/demo/studentservice"
	"github.com/cloudwego/kitex/server"
)

func main() {
	// NEW
	handler := &StudentServiceImpl{}
	
	handler.InitDB()

	addr, _ := net.ResolveTCPAddr("tcp", ":9990")
	svr := demo.NewServer(handler, server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
