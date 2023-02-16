package slog

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Log struct {
	Driver string
	Path   string
	Output *os.File
}

// 实例化log对象
func New() *Log {
	path := "$HOME/.log/"
	driver := "file"
	dirPath := filepath.Dir(path)

	//目录不存在创建
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, os.ModeDir|0755)
	}

	//%DAY%替换
	path = strings.ReplaceAll(path, "%DAY%", time.Now().Format("2006-01-02"))

	//打开文件
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return &Log{
			Driver: driver,
			Path:   "",
			Output: os.Stdout,
		}
	}
	return &Log{
		Driver: driver,
		Path:   path,
		Output: file,
	}
}

func (log *Log) write(level, msg string, args ...interface{}) {
	str := ""
	if len(args) > 0 {
		content, _ := json.Marshal(args)
		str = fmt.Sprintf("%v  [%s]  %s  %s\n", time.Now().Format("2006-01-02T13:04:05.999999"), level, msg, content)
	} else {
		str = fmt.Sprintf("%v  [%s]  %s\n", time.Now().Format("2006-01-02T13:04:05.999999"), level, msg)
	}
	log.Output.Write([]byte(str))
}

func (log *Log) Info(msg string, args ...interface{}) {
	log.write("INFO", msg, args...)
}

func (log *Log) Error(msg string, args ...interface{}) {
	log.write("ERROR", msg, args...)
}

func (log *Log) Debug(msg string, args ...interface{}) {
	log.write("DEBUG", msg, args...)
}

func (log *Log) Notice(msg string, args ...interface{}) {
	log.write("NOTICE", msg, args...)
}

func (log *Log) Warn(msg string, args ...interface{}) {
	log.write("WARN", msg, args...)
}
