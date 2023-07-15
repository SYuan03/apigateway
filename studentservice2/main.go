package main

import (
	"Jiao-Yiyang/d3/kitex_gen/demo2/studentservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	handler := &StudentServiceImpl{}
	handler.InitDB()
	addr, _ := net.ResolveTCPAddr("tcp", ":9991")
	//svr := studentservice.NewServer(handler, server.WithServiceAddr(addr))
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	svr := studentservice.NewServer(handler, server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "studentservice2",
	}), server.WithServiceAddr(addr))
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
