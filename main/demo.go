package main

import (
	"github.com/abnerxc/xcore"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	xcore.Bootstrap("dev")
	xcore.G_LOG.Out = os.Stdout
	xcore.G_LOG.Formatter = &logrus.JSONFormatter{}
	xcore.G_LOG.WithFields(logrus.Fields{}).Info("A group of walrus emerges from the ocean")
}
