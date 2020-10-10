package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}

// 运行结果 time="2020-10-09T19:05:54+08:00" level=info msg="A walrus appears" animal=walrus