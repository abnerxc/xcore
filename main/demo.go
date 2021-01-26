package main

import (
	"github.com/abnerxc/xcore"
)

func main() {
	xcore.Bootstrap("dev")
	data := make(map[string]interface{})
	data["ip"] = "10.10.10.1"
	data["url"] = "http://www.baidu.com"
	data["method"] = "GET"
	data["proto"] = "http"
	data["request"] = "ffffff"
	data["header"] = "mobile"
	xcore.G_LOG.WithFields(data).Info()
}
