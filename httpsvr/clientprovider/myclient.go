package clientprovider

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
)

var JsonClients = make(map[string]genericclient.Client)

// name servicename
func GetCli(serviceName string) genericclient.Client {
	return JsonClients[serviceName]
}

func InitCli(serviceName string, idlPath string, port string) {
	//Todo need to get port and idlcontext by servicename
	p, err := generic.NewThriftFileProvider(idlPath)
	if err != nil {
		panic(err)
	}
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}
	cli, err := genericclient.NewClient(serviceName, g, client.WithHostPorts(port))
	if err != nil {
		panic(err)
	}
	JsonClients[serviceName] = cli
}
