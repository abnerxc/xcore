package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//日志自定义格式
type LogFormatter struct{}

//格式详情
func (s *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	var file string
	var line int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		line = entry.Caller.Line
	}
	level := strings.ToUpper(entry.Level.String())

	content, _ := json.Marshal(entry.Data)

	msg := fmt.Sprintf("%s [%s] [GOID:%d] [%s:%d] #msg:%s #content:%v\n", timestamp, level, getGID(), file, line, entry.Message, string(content))
	return []byte(msg), nil
}

// 获取当前协程id
func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
