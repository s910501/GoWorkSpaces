package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

var (
	MaxSize = 50000
)

type FileLogger struct {
	level       LogLevel
	filepath    string
	filename    string
	fileObj     *os.File
	errFileObje *os.File
	maxFileSize int64
	logChan     chan *logMsg
}

type logMsg struct {
	level     LogLevel
	msg       string
	funName   string
	fileName  string
	timestamp string
	line      int
}

func NewFileLogger(lelevStr, fp, fn string, maxSize int64) *FileLogger {
	LogLevel, err := parseLogLevel(lelevStr)
	if err != nil {
		fmt.Print("err")
	}
	fl := &FileLogger{
		level:       LogLevel,
		filepath:    fp,
		filename:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *logMsg, MaxSize),
	}

	err = fl.initFile()
	if err != nil {
		fmt.Print("err")
	}
	return fl

}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filepath, f.filename)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}
	errfileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err:%v\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObje = errfileObj
	// 运行 多个有问题
	// for i := 0; i < 5; i++ {
	// 	go f.writeLogBackgroud()
	// }
	go f.writeLogBackgroud()
	return nil

}

func (f *FileLogger) enable(LogLevel LogLevel) bool {
	return f.level <= LogLevel
}

func (f *FileLogger) checkSzie(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return false
	}
	return fileInfo.Size() > f.maxFileSize
}

func (f *FileLogger) SplitFile(file *os.File) (*os.File, error) {
	// split log file

	nowStr := time.Now().Format("20060102150405000")
	// fmt.Println(nowStr)
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return nil, err
	}
	logName := path.Join(f.filepath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)

	// close log file
	closeerr := file.Close()
	if closeerr != nil {
		fmt.Println(err)
	}
	// fmt.Println(error)
	// rename
	// time.Sleep(time.Second * 5)
	// Windows test error
	renameerr := os.Rename(logName, newLogName)
	if renameerr != nil {
		fmt.Println(renameerr)
	}
	fmt.Println(logName, newLogName, file)
	// open new file
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %s", err)
	}
	file = fileObj
	// fmt.Println(logName)
	return fileObj, nil
}

func (f *FileLogger) writeLogBackgroud() {
	for {
		if f.checkSzie(f.fileObj) {
			// fmt.Println(222)
			newFile, err := f.SplitFile(f.fileObj)
			if err != nil {
				fmt.Print(1)
				return
			}
			f.fileObj = newFile
		}
		select {
		case logTemp := <-f.logChan:
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", logTemp.timestamp, getLogStrng(logTemp.level), logTemp.fileName, logTemp.funName, logTemp.line, logTemp.msg)
			fmt.Fprintf(f.fileObj, logInfo)
			if logTemp.level >= ERROR {
				if f.checkSzie(f.errFileObje) {
					// fmt.Println(333)
					newFile, err := f.SplitFile(f.errFileObje)
					if err != nil {
						fmt.Print(2)
						return
					}
					f.errFileObje = newFile
				}
				// if recored errr level, need add log to err log file
				fmt.Fprintf(f.errFileObje, logInfo)
			}
		default:
			// no log then sleep 500ms
			time.Sleep(time.Millisecond * 500)
		}

	}

}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format("2006-01-02 15:04:05")
		funcName, filename, lineNo := getInfo(3)
		// send log to channel
		logTemp := &logMsg{
			level:     lv,
			msg:       msg,
			funName:   funcName,
			fileName:  filename,
			timestamp: now,
			line:      lineNo,
		}
		select {
		case f.logChan <- logTemp:
		default:
			// to Null,
		}

	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	if f.enable(DEBUG) {
		f.log(DEBUG, format, a...)
	}

}
func (f *FileLogger) Trace(format string, a ...interface{}) {
	if f.enable(TRACE) {
		f.log(TRACE, format, a...)
	}
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	if f.enable(INFO) {
		f.log(INFO, format, a...)
	}
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
	if f.enable(WARNING) {
		f.log(WARNING, format, a...)
	}
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	if f.enable(ERROR) {
		f.log(ERROR, format, a...)
	}
}
func (f *FileLogger) Faltal(format string, a ...interface{}) {
	if f.enable(FALTAL) {
		f.log(FALTAL, format, a...)
	}
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObje.Close()
}
