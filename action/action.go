package action

import (
	cli "github.com/urfave/cli/v2"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/goctl-swagger/generate"
)

func Generator(ctx *cli.Context) error {
	fileName := ctx.String("filename")

	if len(fileName) == 0 {
		fileName = "rest.swagger.json"
	}

	p, err := plugin.NewPlugin()
	if err != nil {
		return err
	}
	basepath := ctx.String("basepath")
	host := ctx.String("host")
	schemes := ctx.String("schemes")
	pack := ctx.String("pack")
	response := ctx.String("response")

	return generate.Do(fileName, host, basepath, schemes, pack, response, p)
}
