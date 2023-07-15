// Code generated by hertz generator.

package main

import (
	"time"

	"github.com/SYuan03/Day3/HertzServer/myclient"
	"github.com/cloudwego/hertz/pkg/app/server"

	"fmt"
	"os"
)

var ServiceNameList = [2]string{"StudentServiceA", "StudentServiceB"}
var ThriftPathList = [2]string{"./idl/StudentServiceA.thrift", "./idl/StudentServiceB.thrift"}
var HostPortList = [2]string{"127.0.0.1:9990", "127.0.0.1:9991"}

func main() {
	go func() {
		for {
			for i := 0; i < 2; i++ {
				currentPath, err := os.Getwd()
				if err != nil {
					fmt.Println("无法获取当前路径:", err)
					return
				}
				fmt.Println("当前路径:", currentPath)
				myclient.InitCli(ServiceNameList[i], ThriftPathList[i], HostPortList[i])
			}
			time.Sleep(time.Second * 10)
		}
	}()
	h := server.Default()

	register(h)
	h.Spin()
}
