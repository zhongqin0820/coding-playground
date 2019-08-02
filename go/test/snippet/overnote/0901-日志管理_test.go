package overnote

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func init() {
	// 日志格式化为JSON而不是默认的ASCII
	log.SetFormatter(&log.JSONFormatter{})

	// 输出stdout而不是默认的stderr，也可以是一个文件
	log.SetOutput(os.Stdout)

	// 只记录严重或以上警告
	log.SetLevel(log.InfoLevel)
}

func TestLogrus(t *testing.T) {

	log.WithFields(log.Fields{
		"params1": "walrus",
		"params2": 10,
	}).Info("info信息.....")

	log.WithFields(log.Fields{
		"params3": true,
		"params4": 122,
	}).Warn("warn信息.....")

	log.WithFields(log.Fields{
		"params5": true,
		"params6": 100,
	}).Fatal("fatal信息....")
}
