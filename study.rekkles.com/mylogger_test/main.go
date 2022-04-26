package main

import (
	"study.rekkles.com/mylogger"
)

var log mylogger.Logger

func main() {
	log = mylogger.NewConsolelog("Info")
	log = mylogger.NewFileLogger("Info", "./", "test.log", 10*1024*1024)
	for {
		log.Debug("This is Debug log")
		log.Info("This is Info log")
		log.Warning("this is worning log")
		id := 100010
		name := "rekkles"
		log.Error("this is error log, id:%d, name: %s", id, name)
		log.Faltal("this is faltal log")
	}
}
