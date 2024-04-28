package alog

import (
	//json "github.com/json-iterator/go"
	_ "github.com/alog/model"
)

var TimeFormat string = "2006-01-02 15:04:05"
var Level int = 0

var IsTrace bool = false
var TraceIdList []string = []string{}
var IsColor bool = true

//var ele *Element

type LogConfig struct {
	IsConsole        bool   `json:"IsConsole"`        // 是否输出到控制台
	TimeFormat       string `json:"TimeFormat"`       // 控制台日志输出时间格式
	Level            string `json:"Level"`            // 控制台日志输出等级
	LevelInt         int    `json:"LevelInt"`         // 控制台日志输出等级
	Color            bool   `json:"Color"`            // 控制台日志颜色开关
	IsFile           bool   `json:"IsFile"`           // 是否输出到文件   如果开启后没有输入相关对应值，则按默认值
	LogName          string `json:"FileName"`         // 初始日志文件名  	//输出日志打印为 log20201103-150405.log
	LogPath          string `json:"FilePath"`         // 初始日志文件路径 默认当前目录：./
	LogDaily         bool   `json:"Daily"`            // 是否按天生成日志文件
	MaxLogLines      int    `json:"Maxlines"`         // 日志文件最大行数      默认1000000 即100万
	MaxLogSize       int64  `json:"MaxFilesize"`      // 日志文件最大尺寸,单位MB,当文件打开超过这个尺寸时，会重新创建一个日志文件，默认512MB
	MaxFileDays      int    `json:"Maxdays"`          // 日志文件最大天数
	LogZip           bool   `json:"LogZip"`           // 是否压缩成zip格式文件
	IsError          bool   `json:"IsError"`          // 是否输出到错误日志文件 ，如果开启后没有输入相关对应值，则按默认值
	ErrorLogPath     string `json:"ErrorFilePath"`    // 初始日志文件路径   默认当前目录：./
	ErrorLogName     string `json:"ErrorFileName"`    // 初始日志文件路径 //输出错误日志为 error20201103-150405.log
	ErrorDaily       bool   `json:"ErrorDaily"`       // 是否按天生成日志文件
	MaxErrorLogLines int    `json:"MaxErrorLogLines"` // 日志文件最大行数          默认100000即10万
	MaxErrorLogSize  int64  `json:"MaxErrorLogSize"`  // 日志文件最大尺寸,单位MB,当文件打开超过这个尺寸时，会重新创建一个日志文件 默认10M
	ErrorLogZip      bool   `json:"ErrorLogZip"`      // 是否压缩成zip格式文件
	Permit           string `json:"Permit"`           // 新创建的日志文件权限属性
	SaveDbType       string `json:"SaveDbType"`       //日志保存类型保存：配置文件：mysql,redis,mongodb,postgresql,sqlite3,oracle,sqlserver,etcd
	LogType          string `json:"LogType"`          //用户日志还是管理平台日志 ：user,platform、device
	UserName         string `json:"UserName"`         //用户账号
	RequestId        string `json:"RequestId"`        //本次请求的唯一标识，即本次请求id，用于追踪请求日志
	Page             string `json:"Page"`             //请求页面
	Api              string `json:"Api"`              //请求接口
	Function         string `json:"Function"`         //记录函数名或者日志功能
	Seller           string `json:"Seller"`           //商家名称
	SellerId         string `json:"SellerId"`         //商家id
	Token            string `json:"Token"`            //用户token

}
