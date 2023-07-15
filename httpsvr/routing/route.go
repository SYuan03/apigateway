package routing

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"nju/apigw/clientprovider"
	"strings"
)

func ProvideService(serviceName string, methodName string, c *app.RequestContext, ctx context.Context) (string, error) {
	data, err := json.Marshal(c.Request)
	if err != nil {
		return "", err
	}

	dataString := string(data)

	Cli := clientprovider.GetCli(serviceName)
	resp, err := Cli.GenericCall(ctx, methodName, dataString)
	if err != nil {
		return "", err
	}

	realResp := resp.(string)
	realResp = strings.ReplaceAll(realResp, "\"", "")
	return realResp, nil
}
