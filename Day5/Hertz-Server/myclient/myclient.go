package myclient

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
)

var Clients = make(map[string]genericclient.Client)

func InitCli(serviceName string, idlPath string, hostString string) {
	p, err := generic.NewThriftFileProvider(idlPath)
	if err != nil {
		panic(err)
	}

	// g, err := generic.HTTPThriftGeneric(p)
	// if err != nil {
	// 	panic(err)
	// }

	// 构造 JSON 请求和返回类型的泛化调用
    g, err := generic.JSONThriftGeneric(p)
    if err != nil {
        panic(err)
    }

	// cli, err := genericclient.NewClient(serviceName, g, client.WithHostPorts("127.0.0.1:9990"))
	cli, err := genericclient.NewClient(serviceName, g, client.WithHostPorts(hostString))

	if err != nil {
		panic(err)
	}
	Clients[serviceName] = cli
}

func GetCli(name string) genericclient.Client {
	return Clients[name]
}
