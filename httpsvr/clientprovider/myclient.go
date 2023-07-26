package clientprovider

import (
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	etcd "github.com/kitex-contrib/registry-etcd"
	"io/ioutil"
	"log"
	"net/http"
)

var Clients = make(map[string]genericclient.Client)
var IdlVersion = make(map[string]string)

func GetCli(serviceName string, idlVersion string) genericclient.Client {
	value, exist := Clients[serviceName]
	if exist {
		return value
	} else {
		IdlVersion[serviceName] = idlVersion
		UpdateCli(serviceName)
		return Clients[serviceName]
	}
}

func UpdateCli(serviceName string) {
	//Todo: need to get port and idlcontext by servicename
	url := "http://127.0.0.1:6666/idl/query?service_name=" + serviceName + "&version=" + IdlVersion[serviceName]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error getting response:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(body))
	idlPath := "../idl/student2.thrift"
	if serviceName == "studentserviceA" {
		idlPath = "../idl/student.thrift"
	}

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
	cli, err := genericclient.NewClient(serviceName, g, client.WithResolver(r), client.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer()))
	if err != nil {
		panic(err)
	}
	Clients[serviceName] = cli

}
