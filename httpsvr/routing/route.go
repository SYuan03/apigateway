package routing

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"nju/apigw/clientprovider"
	"strings"
)

func ProvideService(serviceName string, methodName string, c *app.RequestContext, ctx context.Context) {
	data, err := json.Marshal(c.Request)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	dataString := string(data)

	Cli := clientprovider.GetCli(serviceName)
	resp, err := Cli.GenericCall(ctx, methodName, dataString)
	if err != nil {
		c.String(consts.StatusBadRequest, "Call Error")
	}

	realResp := resp.(string)
	realResp = strings.ReplaceAll(realResp, "\"", "")
	c.JSON(consts.StatusOK, realResp)
}
