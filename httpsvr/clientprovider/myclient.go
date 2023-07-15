package clientprovider

import (
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	etcd "github.com/kitex-contrib/registry-etcd"
	"io/ioutil"
	"log"
	"net/http"
)

var Clients = make(map[string]genericclient.Client)

// name servicename
func GetCli(serviceName string) genericclient.Client {
	value, exist := Clients[serviceName]
	if exist {
		return value
	} else {
		InitCli(serviceName)
		return Clients[serviceName]
	}
}

func InitCli(serviceName string) {
	//Todo need to get port and idlcontext by servicename

	url := "http://13.72.82.105:8888/"
	reqUrl := fmt.Sprintf("%sgetIdl?serviceName=%s", url, serviceName)
	// 发送HTTP GET请求给idl管理平台 获取idl文件的相对路径
	resp, err := http.Get(reqUrl)
	if err != nil {
		fmt.Printf("Error getting response for %s: %s\n", url, err.Error())
		return
	}
	defer resp.Body.Close()

	// 读取响应正文的内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body for %s: %s\n", url, err.Error())
	}
	//fmt.Println("idlPath", string(body))

	idlPath := string(body)
	idlPath = "../idl/student.thrift"
	p, err := generic.NewThriftFileProvider(idlPath)
	if err != nil {
		panic(err)
	}
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	cli, err := genericclient.NewClient(serviceName, g, client.WithResolver(r))
	if err != nil {
		panic(err)
	}
	Clients[serviceName] = cli

}
