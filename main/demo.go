package main

import (
	"github.com/abnerxc/xcore"
)

func main() {
	xcore.Bootstrap("dev")
	data := make(map[string]interface{})
	for i := 0; i < 10; i++ {
		data["ip"] = "10.10.10.1"
		data["header"] = "mobile"
		data["cnt"] = i
		xcore.G_LOG.WithFields(data).Info("99999999999")
	}
}
