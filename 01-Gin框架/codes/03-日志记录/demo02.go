package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	// 以JSON格式为输出，代替默认的ASCII格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 以Stdout为输出，代替默认的stderr
	logrus.SetOutput(os.Stdout)
	// 设置日志等级
	logrus.SetLevel(logrus.WarnLevel)
}

func main() {

	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	logrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}

// 运行结果
// {"level":"warning","msg":"The group's number increased tremendously!","number":122,"omg":true,"time":"2020-10-09T19:16:01+08:00"}
// {"level":"fatal","msg":"The ice breaks!","number":100,"omg":true,"time":"2020-10-09T19:16:01+08:00"}

