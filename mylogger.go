package alog

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

type LogLevel int64

const (
	UNKNOWN LogLevel = iota

	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	CRIT
	ALRT
	EMER
	INVADE
)

//等级	配置	     释义	                               控制台颜色
//0    UNKNOWN   未知                                    灰白色
//1	   TRAC	   用户级基本输出	                             绿色
//2	   DEBG	   用户级调试	                                 绿色
//3	   INFO	   用户级重要	                                 天蓝色
//4	   WARN	   用户级警告	                                 黄色
//5	   EROR	   用户级错误	                                 红色
//6	   CRIT	   系统级危险，比如权限出错，访问异常等	         蓝色
//7	   ALRT	   系统级警告，比如数据库访问异常，配置文件出错等	 紫色
//8	   EMER	   系统级紧急，比如磁盘出错，内存异常，网络不可用等	 红色底
//9    INVADE  黑客入侵                                    红黑色

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)

	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warn":
		return WARN, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	case "crit":
		return CRIT, nil
	case "alrt":
		return ALRT, nil
	case "emer":
		return EMER, nil
	case "invade":
		return INVADE, nil
	default:
		err := errors.New("无效的日志级别错误")
		return UNKNOWN, err
	}
	//return UNKNOWN, err
}

func getLogString(lv LogLevel) string {

	switch lv {
	case DEBUG:
		return "Debug"
	case TRACE:
		return "Trace"
	case INFO:
		return "Info"
	case WARN:
		return "Warn"
	case ERROR:
		return "Error"
	case FATAL:
		return "Fatal"
	case CRIT:
		return "Crit"
	case ALRT:
		return "Alrt"
	case EMER:
		return "Emer"
	case INVADE:
		return "Invde"
	default:
		return "Unkown"
	}

}

//获取日志文件所在的位置的函数
func getInfo(skip int) (funcName, fileName string, lineNo int) {

	//PC调用函数的相关信息
	//file调用runtime.Caller(n)所在目录和文件名
	//line为调用 runtime.Caller(n)所在的行
	//OK如果无法获得信息，ok会被设为false
	//skip:是要提升的堆栈帧数，0:当前函数，1:上一层函数
	pc, file, lineNo, ok := runtime.Caller(skip)
	//fmt.Println("pc, file, lineNo, ok=", pc, file, lineNo, ok)
	if !ok {
		fmt.Println("runtime.Caller(n) fail")
		return
	}
	funcName = runtime.FuncForPC(pc).Name() //获取函数名
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]

	return
}
