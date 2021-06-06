package alog

import (
	"fmt"
	//json "github.com/json-iterator/go"
	"encoding/json"
	"github.com/liquidMetal/file"
	"strconv"
)

type Log struct {
	//[log]
	//#如果以下不知道如何填写默认即可
	//#on表示日志只输出到log文件，off日志输出到控制台，all表示日志输出到日志文件和控制台
	Valve string
	//#日志存储路径和名字
	TimeFormat string
	//#最终输出日志打印为 print20201103-150405.log
	FilePath   string
	FileName   string
	PrintLevel string
	//#最终输出错误日志为 error20201103-150405.log
	ErrorFilePath string
	ErrorFileName string
	//#大于等于error
	PrintErrorLevel string
	//#通道最大容量
	ChanMaxSize int64
	//文件最大尺寸
	MaxFileSize int64
	test        int64
}

/*
type Element struct {
	Debug   string
	Trace   string
	Info    string
	Warn    string
	Error   string
	Fatal   string
	Crit    string
	Alrt    string
	Emer    string
	Invade  string
	Unknown string
}*/

var logData Log

//var ele *Element

func init() {
	//fmt.Println("=========================start")
	//返回切片
	filePath := "./conf/config.conf"
	msg, err := file.ReadFile(filePath, "")
	if err != nil {
		fmt.Println("File read failed. err=", err)
	}
	//fmt.Println("msg=", msg)
	//设计思路：主要用于配置文件经常变动场景，例如：接口类型的通信，双方未来都不知道未来配置有什么变化，只需要修改配置文件，不需要修改程序，把程序重新运行即可
	//如果开启一个协程周期性的比较文件修改时间，如果有变化则重新读取一次文件，连重行运行程序都不需要
	//因为不用像传统那样传递key才能取到值，当然也能兼容传统类型，通过传递键取值

	for _, v := range msg {
		//fmt.Printf("sn=%d\t topic=%s\t key=%s\t value= %s\n", k, v.Topic, v.Key, v.Value)
		if v.Topic == "log" {
			switch v.Key {
			case "valve":
				logData.Valve = v.Value
			case "timeFormat":
				logData.TimeFormat = v.Value
			case "filePath":
				logData.FilePath = v.Value
			case "fileName":
				logData.FileName = v.Value
			case "printLevel":
				logData.PrintLevel = v.Value
			case "errorFilePath":
				logData.ErrorFilePath = v.Value
			case "errorFileName":
				logData.ErrorFileName = v.Value
			case "printErrorLevel":
				logData.PrintErrorLevel = v.Value
			case "chanMaxSize":
				logData.ChanMaxSize, err = strconv.ParseInt(v.Value, 10, 64)
				//fmt.Println("err=", err)
			case "maxFileSize":
				logData.MaxFileSize, err = strconv.ParseInt(v.Value, 10, 64)
				//fmt.Println("err=", err)
			}

		}
	}

	jsonFormat, err := json.Marshal(logData)
	fmt.Println("配置读取完毕，JSON格式！==>End of reading", string(jsonFormat))
	//	fmt.Println("配置读取完毕，返回结构体！==>End of reading", logData)

}

func Config() Log {
	return logData
}
