package clientprovider

import (
	"encoding/json"
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
	if IdlVersion[serviceName] != idlVersion {
		IdlVersion[serviceName] = idlVersion
		UpdateCli(serviceName)
	}
	value, exist := Clients[serviceName]
	if exist {
		return value
	} else {
		UpdateCli(serviceName)
		return Clients[serviceName]
	}
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Content string `json:"content"`
}

func UpdateCli(serviceName string) {
	//Todo: need to get port and idlcontext by servicename
	url := "http://127.0.0.1:6666/idl/query?service_name=" + serviceName + "&version=" + IdlVersion[serviceName]
	//fmt.Println(url)
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

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error decoding response body:", err)
		return
	}

	content, ok := response["content"].(string)
	if !ok {
		fmt.Println("Error decoding content field")
		return
	}
	fmt.Println(content)
	err = ioutil.WriteFile("../idl/student.thrift", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	idlPath := "../idl/student.thrift"

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
