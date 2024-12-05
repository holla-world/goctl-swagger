package generate

import (
	"testing"

	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
)

func TestDo(t *testing.T) {
	in := &plugin.Plugin{
		Api:         nil,
		ApiFilePath: "./api.api",
		Style:       "gozero",
		Dir:         "./",
	}

	api, err := parser.Parse(in.ApiFilePath)
	if err != nil {
		t.Fatal(err)
	}

	in.Api = api
	err = Do("post-swagger.json", "www.abc.com", "./", "https", "", "", in)
	if err != nil {
		t.Fatal(err)
	}
}
