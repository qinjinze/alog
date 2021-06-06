package alog

import (
	"fmt"
	"time"
)

//判断是否需要记录该日子
func enable(logLevel LogLevel) bool {
	printLevel, _ := parseLogLevel(logData.PrintLevel)
	return logLevel >= printLevel
}

func consoleFormat(lv LogLevel, format string, a ...interface{}) {
	//fmt.Println("Format  lv==============format==============>", format)
	if enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(4)
		fmt.Printf("[%s][%s] [%s=>%s=>%d] %s \n", now, getLogString(lv), fileName, funcName, lineNo, msg)
		abc := fmt.Sprintf("%s", a...)
		//fmt.Println("-----------", abc)
		now = fmt.Sprintf("[%s]", now)
		leven := fmt.Sprintf("[%s]", getLogString(lv))
		file := fmt.Sprintf("[%s=>%s=>%d]", fileName, funcName, lineNo)
		fmt.Printf("%s%s%s%s\n", now, leven, file, format, abc)
	}
}

func Unknown(format string, a ...interface{}) {
	write(UNKNOWN, format, a...)
}

func Debug(format string, a ...interface{}) {
	write(DEBUG, format, a...)
}

func Info(format string, a ...interface{}) {
	//fmt.Println("format=", format, a)
	write(INFO, format, a...)
}

func Warn(format string, a ...interface{}) {
	write(WARN, format, a...)
}

func Error(format string, a ...interface{}) {

	write(ERROR, format, a...)
}
func Fatal(format string, a ...interface{}) {

	write(FATAL, format, a...)

}

func Crit(format string, a ...interface{}) {
	write(CRIT, format, a...)

}
func Alrt(format string, a ...interface{}) {

	write(ALRT, format, a...)

}
func Emer(format string, a ...interface{}) {

	write(EMER, format, a...)

}
func Invade(format string, a ...interface{}) {
	write(INVADE, format, a...)

}

func writeFormat(level LogLevel, format string, a ...interface{}) {
	//fmt.Println("logDatas.Valve=", logDatas.Valve)
	switch logDatas.Valve {
	case "all":
		writeLogFormat(level, format, a...)
		consoleFormat(level, format, a...)
	case "off":
		consoleFormat(level, format, a...)
	case "on":
		writeLogFormat(level, format, a...)
	default:
		consoleFormat(level, format, a...)
	}
}

func UnknownF(format string, a ...interface{}) {
	writeFormat(UNKNOWN, format, a...)
}

func DebugF(format string, a ...interface{}) {
	writeFormat(DEBUG, format, a...)
}

func InfoF(format string, a ...interface{}) {
	//fmt.Println("format=", format, a)
	writeFormat(INFO, format, a...)
}

func WarnF(format string, a ...interface{}) {
	writeFormat(WARN, format, a...)
}

func ErrorF(format string, a ...interface{}) {

	writeFormat(ERROR, format, a...)
}
func FatalF(format string, a ...interface{}) {

	writeFormat(FATAL, format, a...)

}

func CritF(format string, a ...interface{}) {
	writeFormat(CRIT, format, a...)

}
func AlrtF(format string, a ...interface{}) {

	writeFormat(ALRT, format, a...)

}
func EmerF(format string, a ...interface{}) {

	writeFormat(EMER, format, a...)

}
func InvadeF(format string, a ...interface{}) {
	writeFormat(INVADE, format, a...)

}

func write(level LogLevel, format string, a ...interface{}) {
	//fmt.Println("logDatas.Valve=", logDatas.Valve)
	switch logDatas.Valve {
	case "all":
		writeLog(level, format, a...)
		console(level, format, a...)
	case "off":
		console(level, format, a...)
	case "on":
		writeLog(level, format, a...)
	default:
		console(level, format, a...)
	}
}

func console(lv LogLevel, format string, a ...interface{}) {
	//fmt.Println("lv============================>", format)
	if enable(lv) {
		now := time.Now().Format("2006-01-02 15:04:05")
		abc := fmt.Sprintln(a...)
		///	fmt.Println("-----------", abc)
		now = fmt.Sprintf("[%s]", now)
		leven := fmt.Sprintf("[%s]", getLogString(lv))
		funcName, fileName, lineNo := getInfo(4)
		file := fmt.Sprintf("[%s=>%s=>%d]", fileName, funcName, lineNo)
		fileLog := fmt.Sprintf("%s%s%s%s", now, leven, file, format)
		//log := fmt.Sprintln(fileLog, abc)
		fmt.Printf("%s %s", fileLog, abc)
	}
}
