package main

import (
	"github.com/abnerxc/xcore"
)

func main() {
	xcore.Bootstrap("dev")
	data := make(map[string]interface{})
	for i := 0; i < 1000; i++ {
		data["ip"] = "10.10.10.1"
		data["url"] = "http://www.baidu.com"
		data["method"] = "GET"
		data["proto"] = "http"
		data["request"] = "ffffff"
		data["header"] = "mobile"
		data["cnt"] = i
		xcore.G_LOG.WithFields(data).Info()
	}
}
