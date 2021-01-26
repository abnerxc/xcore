package xcore

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

var path = G_APP_PATH + "/main/go.log"

func NewFileHook(logLevel *string, maxRemainCnt uint) logrus.Hook {
	write, err := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)
	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}
	logrus.SetOutput(write)

}
