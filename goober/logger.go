package goober

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

func Logger() *logrus.Logger {

	logger := logrus.New()
	var dev = true

	if dev {
		logger.SetOutput(os.Stdout)
	} else {
		file, _ := outputLogFile()
		logger.SetOutput(file)
	}

	var lv logrus.Level
	if dev {
		lv = logrus.DebugLevel
	} else {
		lv = logrus.InfoLevel
	}

	logger.SetLevel(lv)

	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}
func outputLogFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}

	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	return src, nil
}
