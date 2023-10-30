package main

import (
	"xyhelper-arkose-v2/api"
	"xyhelper-arkose-v2/config"
	"xyhelper-arkose-v2/har"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

func main() {
	ctx := gctx.New()

	// 加载har文件
	loadHarFile(ctx)
	s := g.Server()
	s.SetPort(config.PORT)
	s.BindHandler("/ping", api.Ping)
	s.BindHandler("/token", api.GetToken)
	s.BindHandler("/upload", api.Upload)
	s.Run()

}

// loadHarFile 加载har文件
func loadHarFile(ctx g.Ctx) {
    harFilePath := config.HAR_FILE_PATH
	if gfile.Exists(harFilePath) {
		request, err := har.Parse(ctx, harFilePath)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
		if request == nil {
			g.Log().Error(ctx, "request is nil")
			return
		}
		config.Cache.Set(ctx, "request", request, 0)
		g.Log().Info(ctx, "Load har file success")
	}
}
