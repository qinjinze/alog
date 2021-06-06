package alog

import (
	//"alog/conf"
	"fmt"
	"os"
	"path"
	"time"
)

type logMsg struct {
	level     LogLevel //实际打印出来的日志级别
	msg       string
	abc       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}
type Logfile struct {
	fileObj    *os.File
	errFileObj *os.File
}

var logDatas Log = logData
var logChan = make(chan *logMsg, logDatas.ChanMaxSize)
var logChanFormat = make(chan *logMsg, logDatas.ChanMaxSize)
var logFile Logfile
var flag = true

//根据指定的日志文件路径和文件名打开日志文件
func InitLogConfig() error {
	logDatas = logData
	logChan = make(chan *logMsg, logDatas.ChanMaxSize)
	logChanFormat = make(chan *logMsg, logDatas.ChanMaxSize)
	//fmt.Println("logDatas=", logData, logDatas.ChanMaxSize)

	fullFileName := path.Join(logDatas.FilePath, logDatas.FileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open log file failed,err=", err)
		return err
	}
	errorLogName := path.Join(logDatas.FilePath, logDatas.ErrorFileName)
	//fmt.Println("errorLogName=", errorLogName)
	errFileObj, err := os.OpenFile(errorLogName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open log file failed,err=", err)
		return err
	}
	//日志文件都已经打开了
	logFile.fileObj = fileObj
	logFile.errFileObj = errFileObj
	//开启5个后台的goroutine去往日志里写日志
	/*	for i := 0; i < 5; i++ {
		go f.writeLogBackgroud()
	}*/
	go writeLogFormatBackgroud()
	go writeLogBackgroud()
	return nil

}

//检查log是否需要切割
func checkSize(maxFileSize int64, file *os.File) bool {

	fileInfo, err := file.Stat()
	//fmt.Println("fileInfo.Size()", fileInfo.Size(), maxFileSize)
	if err != nil {
		fmt.Printf("get file info failed,err= %v\n", err)
		return false
	}
	//如果当前文件大小大于设置最大值则返回真

	return fileInfo.Size() >= maxFileSize

}

//切割文件
func spiteFile(file *os.File, filePath string) (*os.File, error) {

	//需要切割日志文件

	//1、备份一下
	//nowStr := time.Now().Format("20060405150405000") //2006-04-05 15:04:05 +000最后三位为毫秒
	nowStr := time.Now().Format(logDatas.TimeFormat) //2006-04-05 15:04:05 +000最后三位为毫秒
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file failed,err=%v\n", err)
		return nil, err
	}
	//path_time := filePath + nowStr
	logName := path.Join(filePath, fileInfo.Name())        //把文件名和路径拼接获得完整路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr) //再拼接一个文件备份的名字

	//2、关闭当前文件
	flag = false

	err = file.Close()
	if err != nil {
		fmt.Println("file.Close().err=", err)
	}
	err = os.Rename(logName, newLogName)
	if err != nil {
		fmt.Println("os.Rename err=", err)
	}

	//3、打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开文件错误：", err)
		return nil, err
	}

	//4、将打开的新日志文件对象赋值给f.fileObj

	flag = true
	return fileObj, nil
}

//后台写日志
func writeLogFormatBackgroud() {

	//fmt.Println("logChan================>?????????????", logFile.fileObj)
	for {

		//如果获取不到日志则logTmp := <-f.logChan管道阻塞了
		select {

		case logTmp := <-logChan:

			//把日志拼出来
			logInfo := fmt.Sprintf("[%s][%s] [%s=>%s=>%d] %s\n", logTmp.timestamp, getLogString(logTmp.level), logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)
			//写入文件
			if flag == true {
				_, err := fmt.Fprintf(logFile.fileObj, logInfo)
				if err != nil {
					fmt.Println("写入print.log文件失败，err=", err)
					return
				}

			}

			if logTmp.level >= ERROR {
				if checkSize(logDatas.MaxFileSize, logFile.errFileObj) {
					newFile, err := spiteFile(logFile.errFileObj, logDatas.ErrorFilePath) //切割错误的日志文件
					if err != nil {
						return
					}
					logFile.errFileObj = newFile
				}
				//如果记录的日志级别大于ERROR级别，我还要在err日志文件中再记录一遍
				_, err := fmt.Fprintf(logFile.errFileObj, logInfo)
				if err != nil {
					fmt.Println("写入error.log文件失败，err=", err)
					return
				}
			}
		default:
			//如果从管道取不出来日志则休息500毫秒
			time.Sleep(time.Millisecond * 500)

		}

	}

}

//记录日志的方法
func writeLogFormat(lv LogLevel, format string, a ...interface{}) {

	//判断是否是否满足写日志级别
	if enable(lv) {
		a := fmt.Sprintf("%s", a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(4)
		//先把日志发送到通道中
		//1、构造一个logMsg对象
		logTmp := logMsg{
			level:     lv,
			msg:       format,
			abc:       a,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: now.Format("2006-01-02 15:04:05.0000000 +0800"),
			line:      lineNo,
		}

		//fmt.Println("logTmp=================>", logTmp)
		//避免通道满或者异常导致无法写入
		select {
		//正常写入
		case logChan <- &logTmp:

		default:
			//	当无法写入的日志的时候，丢掉日志
		}
		//fmt.Println("logFile.fileObj=", logFile.fileObj.Name())

		if checkSize(logDatas.MaxFileSize, logFile.fileObj) {
			newFile, err := spiteFile(logFile.fileObj, logDatas.FilePath) //切割正常的日志文件
			if err != nil {
				fmt.Println("切割日志文件出错")
				return
			}

			logFile.fileObj = newFile

		}

	}
}

//记录日志的方法
func writeLog(lv LogLevel, format string, a ...interface{}) {

	//判断是否是否满足写日志级别
	if enable(lv) {

		a := fmt.Sprintln(a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(4)
		//先把日志发送到通道中
		//1、构造一个logMsg对象
		logTmpFoamt := logMsg{
			level:     lv,
			msg:       format,
			abc:       a,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: now.Format("2006-01-02 15:04:05.0000000 +0800"),
			line:      lineNo,
		}
		//fmt.Println("a=====================<<<<<<<<<<<<<<<<<<", a)
		//fmt.Println("format=================>", format)
		//避免通道满或者异常导致无法写入
		select {
		//正常写入
		case logChanFormat <- &logTmpFoamt:

		default:
			//	当无法写入的日志的时候，丢掉日志
		}
		//fmt.Println("logFile.fileObj=", logFile.fileObj.Name())

		if checkSize(logDatas.MaxFileSize, logFile.fileObj) {
			newFile, err := spiteFile(logFile.fileObj, logDatas.FilePath) //切割正常的日志文件
			if err != nil {
				fmt.Println("切割日志文件出错")
				return
			}

			logFile.fileObj = newFile

		}

	}
}

//后台写日志
func writeLogBackgroud() {

	//fmt.Println("logChan======format==========>?????????????", logFile.fileObj)
	for {

		//如果获取不到日志则logTmp := <-f.logChan管道阻塞了
		select {

		case logTmp := <-logChanFormat:

			//now := logTmp.timestamp
			//abc := fmt.Sprintf("%s", a...)
			//fmt.Println("-----------", abc)
			now := fmt.Sprintf("[%s]", logTmp.timestamp)
			leven := fmt.Sprintf("[%s]", getLogString(logTmp.level))
			//funcName, fileName, lineNo := getInfo(4)
			file := fmt.Sprintf("[%s=>%s=>%d]", logTmp.fileName, logTmp.funcName, logTmp.line)
			//	abc := fmt.Sprintln(logTmp.abc)
			fileLog := fmt.Sprintf("%s%s%s%s", now, leven, file, logTmp.msg)
			logInfo := fmt.Sprintln(fileLog, logTmp.abc)
			//fmt.Println("logInfo<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<", logInfo)
			//把日志拼出来
			//logInfo := fmt.Sprintf("[%s][%s] [%s=>%s=>%d] %s\n", logTmp.timestamp, getLogString(logTmp.level), logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)
			//写入文件
			if flag == true {
				_, err := fmt.Fprintf(logFile.fileObj, logInfo)
				if err != nil {
					fmt.Println("写入print.log文件失败，err=", err)
					return
				}

			}

			if logTmp.level >= ERROR {
				if checkSize(logDatas.MaxFileSize, logFile.errFileObj) {
					newFile, err := spiteFile(logFile.errFileObj, logDatas.ErrorFilePath) //切割错误的日志文件
					if err != nil {
						return
					}
					logFile.errFileObj = newFile
				}
				//如果记录的日志级别大于ERROR级别，我还要在err日志文件中再记录一遍
				_, err := fmt.Fprintf(logFile.errFileObj, logInfo)
				if err != nil {
					fmt.Println("写入error.log文件失败，err=", err)
					return
				}
			}
		default:
			//如果从管道取不出来日志则休息500毫秒
			time.Sleep(time.Millisecond * 500)

		}

	}

}
