package routing

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"nju/apigw/clientprovider"
	"strings"
)

func ProvideService(serviceName string, methodName string, c *app.RequestContext, ctx context.Context, idlVersion string) (string, error) {
	var data interface{}

	err := json.Unmarshal(c.Request.Body(), &data)

	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	jsonString := string(jsonData)

	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return "", err
	}

	Cli, err := clientprovider.GetCli(serviceName, idlVersion)
	if err != nil {
		return "", err
	}
	resp, err := Cli.GenericCall(ctx, methodName, jsonString)
	if err != nil {
		return "", err
	}

	realResp := resp.(string)
	//maybe
	realResp = strings.ReplaceAll(realResp, "\"", "")
	return realResp, nil
}
