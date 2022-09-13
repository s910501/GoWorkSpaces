package myloggerChannel

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

type Logger interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Faltal(format string, a ...interface{})
	Trace(format string, a ...interface{})
}

func getInfo(skip int) (funcName, filename string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Print("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	filename = path.Base(file)
	return
}
