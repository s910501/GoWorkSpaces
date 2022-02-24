package main

import (
	"study.rekkles.com/mylogger"
)

var log mylogger.Logger

func main() {
	// fileObj, err := os.OpenFile("xx.log", os.O_CREATE|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Printf("open file failed, err: %v", err)
	// 	return
	// }
	// log.SetOutput(fileObj)
	// for {
	// 	log.Println("this is test log")
	// 	time.Sleep(time.Second * 3)

	// }
	log = mylogger.NewConsolelog("debug")
	log = mylogger.NewFileLogger("debug", "./", "filelog.txt", 10*1024)
	for {
		username := "rekkles"
		log.Error("This is Error log")
		log.Trace("This is TraceError log")
		log.Debug("This is Debug log user:%s", username)
		log.Info("This is Info log")
		log.Warning("This is Warning log")
		log.Faltal("This is Faltal log")
		// time.Sleep(time.Second)
		// nowStr := time.Now().Format("20160102150405000")
		// fmt.Println(nowStr)
	}

}
