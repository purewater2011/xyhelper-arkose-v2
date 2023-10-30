package config

import (
    "os"
    "strconv"
// 	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
// 	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

var (
	PORT    = 8080
	PROXY   = ""
	Cache   = gcache.New()
	AUTHKEY = os.Getenv("AUTHKEY")
	HAR_FILE_PATH = "./temp/request.har"
	WAIT = "300"
)

func init() {
// 	ctx := gctx.GetInitCtx()
	port, _ := strconv.Atoi(os.Getenv("ARKOSE_PORT"))
	if port > 0 {
		PORT = port
	}
	proxy := os.Getenv("PROXY")
	if proxy != "" {
		PROXY = proxy
	}
	harFilePath := os.Getenv("HAR_FILE_PATH")
    if gfile.Exists(harFilePath) {
        HAR_FILE_PATH = harFilePath
    }
    wait := os.Getenv("WAIT")
    if (wait != "") && (wait != "0") {
        WAIT = wait
    }
}
