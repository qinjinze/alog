package utils

import (
	"github.com/beego/beego/config"

	"github.com/wonderivan/logger"
	"strconv"
)

var (
	G_server_name string //项目名称
	//G_server_addr string //服务器ip地址
	//G_server_port string //服务器端口
	G_redis_addr       string //redis ip地址
	G_redis_port       string //redis port端口
	G_redis_dbnum      int    //redis db 编号
	G_redis_password   string //redis db 密码
	G_mysql_addr       string //mysql ip 地址
	G_mysql_port       string //mysql 端口
	G_mysql_dbUserName string //mysql db name
	G_mysql_dbPassword string //mysql db 密码

	DbName   string
	MqttHost string

	TimeFormat  string
	LogPath     string
	LogName     string
	MaxLogLines int
	MaxLogSize  int64

	ErrorLogPath     string
	ErrorLogName     string
	MaxErrorLogLines int
	MaxErrorLogSize  int64

	LogLevel       string
	LogBackupCount string
	LogDir         string
)

var IsConsole bool = true

func InitinitConfigFile() {

	//从配置文件读取配置信息
	//如果项目迁移此处需要修改配置文件路径
	appconf, err := config.NewConfig("ini", "./conf/app.conf")
	if err != nil {
		logger.Error("读取配置文件错误 ./conf/app.conf,err=", err)
		appconf, err = config.NewConfig("ini", "../conf/app.conf")
		if err != nil {
			logger.Error("读取配置文件错误 ../conf/app.conf:,err=", err)
			return
		}
	}
	G_server_name = appconf.String("appname")
	G_redis_addr = appconf.String("redisaddr")
	G_redis_port = appconf.String("redisport")
	redis_port := appconf.String("redisdbnum")

	G_redis_password = appconf.String("password")
	G_mysql_addr = appconf.String("mysqladdr")
	G_mysql_addr = appconf.String("mysqladdr")
	G_mysql_port = appconf.String("mysqlport")
	G_mysql_dbUserName = appconf.String("dbUserName")
	G_mysql_dbPassword = appconf.String("mysqlPassword")

	G_redis_dbnum, err = strconv.Atoi(redis_port)
	if err != nil {
		logger.Error("strconv.Atoi(redis_port),err=", err)
		panic(err)
	}
	DbName = appconf.String("dbName")
	MqttHost = appconf.String("mqtthost")

	isConsole := appconf.String("IsConsole")
	if isConsole == "true" {
		IsConsole = true
	}

	LogPath = appconf.String("LogPath")
	LogName = appconf.String("LogName")
	ErrorLogPath = appconf.String("ErrorLogPath")
	ErrorLogName = appconf.String("ErrorLogName")
	TimeFormat = appconf.String("TimeFormat")

	LogLevel = appconf.String("LogLevel")

	LogBackupCount = appconf.String("LogBackupCount")
	LogDir = appconf.String("LogDir")
	MaxLogLinesStr := appconf.String("MaxLogLines")
	MaxLogLines, err = strconv.Atoi(MaxLogLinesStr)
	if err != nil {
		logger.Error("strconv.Atoi(MaxLogLinesStr),err=", err)
		panic("strconv.Atoi(MaxLogLinesStr),err=")
	}
	MaxLogSizeStr := appconf.String("MaxLogSize")
	MaxLogSize, err = strconv.ParseInt(MaxLogSizeStr, 10, 64)
	if err != nil {
		logger.Error("strconv.ParseInt(MaxLogSizeStr),err=", err)
		panic("strconv.ParseInt(MaxLogSizeStr, 10, 64),err=")
	}

	MaxErrorLogLinesStr := appconf.String("MaxErrorLogLines")
	MaxErrorLogLines, err = strconv.Atoi(MaxErrorLogLinesStr)
	if err != nil {
		logger.Error("strconv.Atoi(MaxErrorLogLines),err=", err)
		panic("strconv.Atoi(MaxErrorLogLines),err=")
	}
	MaxErrorLogSizeStr := appconf.String("MaxErrorLogSize")
	MaxErrorLogSize, err = strconv.ParseInt(MaxErrorLogSizeStr, 10, 64)
	if err != nil {
		logger.Error("strconv.ParseInt(MaxErrorLogSize),err=", err)
		panic("strconv.ParseInt(MaxErrorLogSize, 10, 64),err=")
	}

}
