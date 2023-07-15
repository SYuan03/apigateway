package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
	"nju/apigw/kitex_gen/demo/studentservice"
)

func main() {
	//TODO 注册

	handler := &StudentServiceImpl{}
	handler.InitDB()

	addr, _ := net.ResolveTCPAddr("tcp", ":9990")
	//svr := studentservice.NewServer(handler, server.WithServiceAddr(addr))
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	svr := studentservice.NewServer(handler, server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "studentserviceA",
	}), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
