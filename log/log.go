package log

import (
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

func NewLfsHook(filePath string) logrus.Hook {
	infoWriter, err := rotate.New(
		// 分割后的文件名称
		filePath+".%Y%m%d.info.log",
		//// 生成软链，指向最新日志文件
		rotate.WithLinkName(filePath),
		// 设置日志切割时间间隔(1天)
		rotate.WithRotationTime(24*time.Hour),
		// 设置最大保存时间(30天)
		rotate.WithMaxAge(30*24*time.Hour),
		rotate.WithRotationSize(5),
	)
	errorWriter, err := rotate.New(
		// 分割后的文件名称
		filePath+".%Y%m%d.error.log",
		//// 生成软链，指向最新日志文件
		rotate.WithLinkName(filePath),
		// 设置日志切割时间间隔(1天)
		rotate.WithRotationTime(24*time.Hour),
		// 设置最大保存时间(30天)
		rotate.WithMaxAge(30*24*time.Hour),
	)

	if err != nil {
		logrus.Errorf("config logger error: %v", err)
	}

	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: infoWriter,
		logrus.InfoLevel:  infoWriter,
		logrus.WarnLevel:  errorWriter,
		logrus.ErrorLevel: errorWriter,
		logrus.FatalLevel: errorWriter,
		logrus.PanicLevel: errorWriter,
	}
	lfsHook := lfshook.NewHook(writeMap, new(LogFormatter))

	return lfsHook
}
