package alog

import (
	"fmt"
	"runtime"
	"strings"
)

//type LogLevel int

const (
	UNKNOWN int = iota

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

// 等级	配置	       释义	                                   控制台颜色
// 0    UNKNOWN    未知                                       灰白色
// 2	   DEBG	   用户级调试	                                 灰白色
// 3	   INFO	   用户级重要	                                 天蓝色
// 4	   WARN	   用户级警告	                                 黄色
// 5	   EROR	   用户级错误	                                 红色
// 6	   FATAL  用户级基本输出	                             粉色
// 7	   CRIT	   系统级危险，比如权限出错，访问异常等	         粉色
// 8	   ALRT	   系统级警告，比如数据库访问异常，配置文件出错等	 粉色
// 9	   EMER	   系统级紧急，比如磁盘出错，内存异常，网络不可用等	 粉色
// 10    INVADE   黑客入侵                                     绿色

// 解析日志等级:debug, trace, info, warn, error, fatal, crit, alrt, emer, invade
func ParseLogLevel(s string) int {
	s = strings.ToLower(s)

	switch s {
	case "debug":
		return DEBUG
	case "trace":
		return TRACE
	case "info":
		return INFO
	case "warn":
		return WARN
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	case "crit":
		return CRIT
	case "alrt":
		return ALRT
	case "emer":
		return EMER
	case "invade":
		return INVADE
	default:
		//err := errors.New("无效的日志级别错误")
		return UNKNOWN
	}
	//return UNKNOWN, err
}

func GetLogString(lv int) string {

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

// 获取日志文件所在的位置的函数
func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller(n) fail")
		return
	}
	funcName = runtime.FuncForPC(pc).Name() //获取函数名

	fileNameList := strings.Split(file, "src/")
	if len(fileNameList) > 1 {
		fileName = fileNameList[1]
	}
	funcName = strings.Split(funcName, ".")[1]

	return
}
