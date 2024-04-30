package alog

import (
	"crypto/tls"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	//"github.com/qinjinze/alog/model"
	//"github.com/qinjinze/alog/utils"
	"alog/model"
	"alog/utils"
	"github.com/wonderivan/logger"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type Logfile struct {
	fileObj    *os.File
	errFileObj *os.File
}
type LogDataList struct {
	LogConfigInfo LogConfig
	PlatformLog   []model.PlatformLog
	UserLog       []model.UserLog
	DeviceLog     []model.DeviceLog
}
type LogDataTrace struct {
	LogConfigInfo LogConfig
	PlatformLog   model.PlatformLog
	UserLog       model.UserLog
	DeviceLog     model.DeviceLog
}

// var Orm = orm.NewOrm()
var static_mu *sync.Mutex = new(sync.Mutex)

// var logChan = make(chan *logMsg, logDatas.ChanMaxSize)
// var logChanFormat = make(chan *logMsg, logDatas.ChanMaxSize)
var logChanList = make(chan LogDataList, 1000)
var logChanTrace = make(chan LogDataTrace, 512)
var logChanFileList = make(chan LogDataList, 1000)
var logDbList LogDataList
var logFileList LogDataList
var logDataTrace LogDataTrace
var logFile Logfile

// var flag = true
var logRowCount int = 0
var errorLogRowCount int = 0
var logSize int64 = 512 * 1024 * 1024
var errorLogSize int64 = 10 * 1024 * 1024
var isSpiteLog = false
var isSpiteErrorLog = false
var TraceIdPeriod = make(map[string]time.Time)
var isDbAndFile = false
var client MQTT.Client
var qos *int
var retained *bool
var topic *string

var connClient = make(map[string]*websocket.Conn)
var isWebsocket = true

// 根据指定的日志文件路径和文件名打开日志文件
func (config *LogConfig) InitLogConfig() error {
	logger.Info("InitLogConfig初始化日志数据，config=%+v", config)
	logger.Info("isWebsocket=", isWebsocket)
	model.InitModel(config.DbUserName, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	var err error
	if config.IsFile {
		if utils.MaxLogSize != 0 {
			config.MaxLogSize = utils.MaxLogSize * 1024 * 1024
			logSize = config.MaxLogSize
		} else {
			if config.MaxLogSize != 0 {
				logSize = config.MaxLogSize * 1024 * 1024 //MB
			}
		}

		if utils.MaxLogSize != 0 {
			config.MaxLogLines = utils.MaxLogLines
		} else {
			if config.MaxLogLines == 0 {
				config.MaxLogLines = 100000
			}
		}

		if utils.LogPath != "" {
			config.LogPath = utils.LogPath
		} else {
			if config.LogPath == "" {
				config.LogPath = "./"
			}
		}
		if utils.LogName != "" {
			config.LogName = utils.LogName
		} else {
			if config.LogName == "" {
				config.LogName = "log.log"
			}
		}

		//日志文件都已经打开了
		logFile.fileObj, err = config.openLogFile()
		if err != nil {
			logger.Error("open log file failed,err=", err)
			return err
		}
		go config.checkLog()

	}

	if config.IsError {
		if utils.MaxErrorLogSize != 0 {
			config.MaxLogSize = utils.MaxErrorLogSize * 1024 * 1024
			errorLogSize = utils.MaxErrorLogSize
		} else {
			if config.MaxErrorLogSize != 0 {
				errorLogSize = config.MaxErrorLogSize * 1024 * 1024 //MB
			}
		}

		if utils.MaxErrorLogLines != 0 {
			config.MaxErrorLogLines = utils.MaxErrorLogLines
		} else {
			if config.MaxErrorLogLines == 0 {
				config.MaxErrorLogLines = 100000
			}
		}

		if utils.ErrorLogPath != "" {
			config.ErrorLogPath = utils.ErrorLogPath
		} else {
			if config.ErrorLogPath == "" {
				config.ErrorLogPath = "./"
			}
		}

		if utils.ErrorLogName != "" {
			config.ErrorLogName = utils.ErrorLogName
		} else {
			if config.ErrorLogName == "" {
				config.ErrorLogName = "error.log"
			}
		}

		logFile.errFileObj, err = config.openEerrorLogFile()
		if err != nil {
			logger.Error("open log file failed,err=", err)
			return err
		}
		go config.checkEerrorLog()
	}

	if (config.IsError || config.IsFile) && config.SaveDbType != "" {
		isDbAndFile = true
		go writeLogToDbAndFile()
		//go config.connectMqtt()
		if isWebsocket {
			isWebsocket = false
			go websocketListen()
		}

	} else {
		if config.IsError || config.IsFile {
			go writeLogToFile()
			//go config.connectMqtt()
			if isWebsocket {
				isWebsocket = false
				go websocketListen()
			}
		}

		if config.SaveDbType != "" {
			go writeLogToDb()
			//go config.connectMqtt()
			if isWebsocket {
				isWebsocket = false
				go websocketListen()
			}
		}
	}
	if config.TimeFormat != "" {
		TimeFormat = config.TimeFormat
	}
	config.LevelInt = ParseLogLevel("debug")
	logger.Info("InitLogConfig初始化日志数据完毕，config=%+v", config)
	return nil

}

// 连接mqtt
func (config LogConfig) connectMqtt() {
	imei := "869825040349147"
	connOpts := MQTT.NewClientOptions().AddBroker(utils.MqttHost).SetClientID(imei).SetAutoReconnect(false)
	username := "addddddvvvW123777"
	password := "dsfgyhyjj34679*^M1232"
	if username != "" {
		connOpts.SetUsername(username)
		if password != "" {
			connOpts.SetPassword(password)
		}
	}

	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)

	client = MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Error("客户端连接失败：", token.Error())
		var i int
		for {
			logger.Info("客户端继续连接MQTT服务器...：", i)
			token := client.Connect()
			if token.Wait() && token.Error() == nil {
				break
			}
			time.Sleep(60 * time.Second)
			i++
		}
		// return
	}
	//client.Publish(*topic, byte(*qos), *retained, "hello world,I am a alog client")
}

func (config *LogConfig) openLogFile() (*os.File, error) {
	fullFileName := path.Join(config.LogPath, config.LogName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error("open log file failed,err=", err)
		return fileObj, err
	}
	return fileObj, nil
}
func (config *LogConfig) openEerrorLogFile() (*os.File, error) {
	errorLogName := path.Join(config.ErrorLogPath, config.ErrorLogName)
	//logger.Error("errorLogName=", errorLogName)
	errFileObj, err := os.OpenFile(errorLogName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error("open log file failed,err=", err)
		return errFileObj, err
	}
	return errFileObj, nil
}

// 检查log是否需要切割
func checkSize(isErrorLog bool, file *os.File) bool {

	fileInfo, err := file.Stat()

	if err != nil {
		fmt.Printf("get file info failed,err= %v\n", err)
		return false
	}
	//如果当前文件大小大于设置最大值则返回真
	//Size()  按字节返回文件大小bytes
	if isErrorLog {
		return fileInfo.Size() >= errorLogSize
	}

	return fileInfo.Size() >= logSize
}

// 切割文件
func spiteFile(file *os.File, filePath string) (*os.File, error) {

	//需要切割日志文件

	//1、备份一下
	nowStr := time.Now().Format("20060102-150405")
	fileInfo, err := file.Stat()
	if err != nil {
		logger.Error("get file failed,err=", err)
		return nil, err
	}

	logName := path.Join(filePath, fileInfo.Name()) //把文件名和路径拼接获得完整路径

	newName := nowStr + logName //再拼接一个文件备份的名字

	//2、关闭当前文件
	//flag = false
	static_mu.Lock()
	err = file.Close()
	if err != nil {
		logger.Error("file.Close().err=", err)
	}
	err = os.Rename(logName, newName)
	if err != nil {
		logger.Error("os.Rename err=", err)
	}
	static_mu.Unlock()
	//3、打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error("打开文件错误：", err)
		return nil, err
	}

	return fileObj, nil
}

// 记录日志的方法
func (config LogConfig) writeLog(msg, level string, levelInt int) {

	switch config.LogType {
	case "user":
		message := model.UserLog{
			Level:     level,
			LevelInt:  levelInt,
			Function:  config.Function,
			RequestId: config.RequestId,
			Token:     config.Token,
			UserName:  config.UserName,
			Content:   msg,
			LogTime:   time.Now(),

		}
		logDbList.UserLog = append(logDbList.UserLog, message)
		logDbList.LogConfigInfo = config

		if IsTrace {
			for _, s := range TraceIdList {
				if s == message.UserName {
					logDataTrace.UserLog = message
					logDataTrace.LogConfigInfo = config
					logChanTrace <- logDataTrace
					break
				}
			}
		}
		if (config.IsFile || config.IsError) && !isDbAndFile {
			logFileList.UserLog = append(logFileList.UserLog, message)
			logFileList.LogConfigInfo = config
		}

	case "platform":

		message := model.PlatformLog{
			Level:     level,
			LevelInt:  levelInt,
			Function:  config.Function,
			RequestId: config.RequestId,
			Token:     config.Token,
			UserName:  config.UserName,
			Content:   msg,
			LogTime:   time.Now(),

		}

		logDbList.PlatformLog = append(logDbList.PlatformLog, message)
		logDbList.LogConfigInfo = config

		if IsTrace {
			for _, s := range TraceIdList {
				if s == message.UserName {
					logDataTrace.PlatformLog = message
					logDataTrace.LogConfigInfo = config
					logChanTrace <- logDataTrace
					break
				}
			}
		}
		if (config.IsFile || config.IsError) && !isDbAndFile {
			logFileList.PlatformLog = append(logFileList.PlatformLog, message)
			logFileList.LogConfigInfo = config
		}
	case "device":
		message := model.DeviceLog{
			Level:     level,
			LevelInt:  levelInt,
			Function:  config.Function,
			RequestId: config.RequestId,
			Token:     config.Token,
			UserName:  config.UserName,
			Content:   msg,
			LogTime:   time.Now(),

		}

		if IsTrace {
			for _, s := range TraceIdList {
				if s == message.UserName {
					logDataTrace.DeviceLog = message
					logDataTrace.LogConfigInfo = config
					logChanTrace <- logDataTrace
					break
				}
			}
		}

		logDbList.DeviceLog = append(logDbList.DeviceLog, message)
		logDbList.LogConfigInfo = config

		if (config.IsFile || config.IsError) && !isDbAndFile {
			logFileList.DeviceLog = append(logFileList.DeviceLog, message)
			logFileList.LogConfigInfo = config
		}
	default:
		message := model.PlatformLog{
			Level:     level,
			LevelInt:  levelInt,
			Function:  config.Function,
			RequestId: config.RequestId,
			Token:     config.Token,
			UserName:  config.UserName,
			Content:   msg,
			LogTime:   time.Now(),

		}

		if IsTrace {
			for _, s := range TraceIdList {

				if s == message.UserName {
					logDataTrace.PlatformLog = message
					logDataTrace.LogConfigInfo = config
					logChanTrace <- logDataTrace
					break
				}
			}
		}

		logDbList.PlatformLog = append(logDbList.PlatformLog, message)
		logDbList.LogConfigInfo = config

		if (config.IsFile || config.IsError) && !isDbAndFile {
			logFileList.PlatformLog = append(logFileList.PlatformLog, message)
			logFileList.LogConfigInfo = config
		}
	}

	//避免通道满或者异常导致无法写入
	if config.SaveDbType != "" || isDbAndFile {
		if len(logChanList) < 2 || len(logDbList.PlatformLog) > 999 {

			select {
			//正常写入
			case logChanList <- logDbList:

				logDbList = LogDataList{}
			default:
				fmt.Printf("LogChanList is full, cap=%d, len=%d\n", cap(logChanList), len(logDbList.PlatformLog))
				//	当无法写入的日志的时候，丢掉日志
			}

		}

	}
	if (config.IsFile || config.IsError) && !isDbAndFile {

		if len(logChanList) < 2 || len(logFileList.PlatformLog) > 999 {
			select {
			//正常写入
			case logChanFileList <- logFileList:
				logDbList = LogDataList{}

			default:
				fmt.Printf("logChanFileList is full, cap=%d, len=%d\n", cap(logChanFileList), len(logFileList.PlatformLog))

			}

		}

	}

}

// 后台写日志
func writeLogToFile() {
	for {
		//如果获取不到日志则logTmp := <-f.logChan管道阻塞了
		select {
		case logTmpList := <-logChanFileList:
			row := len(logTmpList.PlatformLog)
			//把所有日志保存文件
			if logTmpList.LogConfigInfo.IsFile {
				if isSpiteLog {
					time.Sleep(time.Second * 2)
				}
				for _, log := range logTmpList.PlatformLog {
					_, err := fmt.Fprintf(logFile.fileObj, log.Content)
					if err != nil {
						logger.Error("写入log文件失败，err=", err)
						time.Sleep(time.Second * 1)
						logFile.fileObj.Close()
						logTmpList.LogConfigInfo.openLogFile()
						_, err := fmt.Fprintf(logFile.fileObj, log.Content)
						if err != nil {
							logger.Error("关闭log后再打开log文件，再次写入log文件失败，放弃写入，err=", err)
						} else {
							logRowCount++
						}

					} else {
						logRowCount++
					}
					if logRowCount >= logTmpList.LogConfigInfo.MaxLogLines {
						//logFile.fileObj.Close()
						if !isSpiteLog {
							isSpiteLog = true
							newFile, err := spiteFile(logFile.fileObj, logTmpList.LogConfigInfo.LogPath) //切割错误的日志文件
							isSpiteLog = false
							if err != nil {
								logger.Error("切割错误日志文件出错，err=", err)
								return
							}
							logFile.fileObj = newFile
						}

					}

				}

			}

			if row >= 1000 {
				if checkSize(false, logFile.fileObj) {
					if !isSpiteLog {
						isSpiteLog = true
						newFile, err := spiteFile(logFile.fileObj, logTmpList.LogConfigInfo.LogPath) //切割错误的日志文件
						isSpiteLog = false
						if err != nil {
							logger.Error("切割错误日志文件出错，err=", err)
							return
						}
						logFile.fileObj = newFile
					}

				}
			}

			//把所有错误保存文件
			if logTmpList.LogConfigInfo.IsError {
				if isSpiteErrorLog {
					time.Sleep(time.Second * 2)
				}
				tmpCount := 0
				for _, log := range logTmpList.PlatformLog {
					if log.LevelInt >= ERROR {

						//如果记录的日志级别大于ERROR级别，我还要在err日志文件中再记录一遍
						_, err := fmt.Fprintf(logFile.errFileObj, log.Content)
						if err != nil {
							logger.Error("写入error.log文件失败，err=", err)
							time.Sleep(time.Second * 1)
							logFile.errFileObj.Close()
							logTmpList.LogConfigInfo.openEerrorLogFile()
							_, err := fmt.Fprintf(logFile.errFileObj, log.Content)
							if err != nil {
								logger.Error("关闭错误日志log文件后再次打开，再次写入log文件失败，放弃写入，err=", err)
								continue
							} else {
								errorLogRowCount++
								tmpCount++
							}
						} else {
							errorLogRowCount++
							tmpCount++
						}
						if errorLogRowCount >= logTmpList.LogConfigInfo.MaxErrorLogLines {
							if !isSpiteErrorLog {
								isSpiteErrorLog = true
								newFile, err := spiteFile(logFile.errFileObj, logTmpList.LogConfigInfo.LogPath) //切割错误的日志文件
								isSpiteErrorLog = false
								if err != nil {
									logger.Error("切割错误日志文件出错，err=", err)
									return
								}
								logFile.errFileObj = newFile
							}

						}
					}
				}
				if tmpCount >= 1000 {
					if checkSize(true, logFile.errFileObj) {
						if !isSpiteErrorLog {
							isSpiteErrorLog = true
							newFile, err := spiteFile(logFile.errFileObj, logTmpList.LogConfigInfo.LogPath) //切割错误的日志文件
							isSpiteErrorLog = false
							if err != nil {
								logger.Error("切割错误日志文件出错，err=", err)
								return
							}
							logFile.errFileObj = newFile
						}

					}
				}

			}

		case logTmpTrace := <-logChanTrace:
			//把日志发送到mqtt中
			//client.Publish("sub_"+logTmpTrace.PlatformLog.UserName, byte(*qos), *retained, logTmpTrace.PlatformLog.Content)
			if logTmpTrace.LogConfigInfo.UserName != "" {
				if conn, ok := connClient[logTmpTrace.LogConfigInfo.UserName]; ok {
					err := conn.WriteMessage(websocket.TextMessage, []byte(logTmpTrace.PlatformLog.Content))
					if err != nil {
						logger.Error("write message error,err=", err, "userName=", logTmpTrace.LogConfigInfo.UserName)
					}
				} else {
					logger.Error("connClient not found, userName=", logTmpTrace.LogConfigInfo.UserName)
				}
			}
		default:
			//如果从管道取不出来日志则休息500毫秒
			time.Sleep(time.Millisecond * 500)

		}

	}

}

func writeLogToDb() {

	for {
		//logTmpList := <-logChanList管道阻塞了
		select {
		case logTmpList := <-logChanList:

			row := len(logTmpList.PlatformLog)

			db := model.Db.CreateInBatches(logTmpList.PlatformLog, row)
			if db.Error != nil {
				logger.Error("db.CreateInBatches error:", db.Error)
				logger.Error("db.CreateInBatches RowsAffected:", db.RowsAffected)
			}

		case logTmpTrace := <-logChanTrace:

			if logTmpTrace.LogConfigInfo.UserName != "" {
				if conn, ok := connClient[logTmpTrace.LogConfigInfo.UserName]; ok {
					err := conn.WriteMessage(websocket.TextMessage, []byte(logTmpTrace.PlatformLog.Content))
					if err != nil {
						logger.Error("write message error,err=", err, "userName=", logTmpTrace.LogConfigInfo.UserName)
					}
				} else {
					logger.Error("connClient not found, userName=", logTmpTrace.LogConfigInfo.UserName)
				}
			}
		default:
			//如果从管道取不出来日志则休息500毫秒
			time.Sleep(time.Millisecond * 500)
		}

	}

}
func writeLogToDbAndFile() {

	for {

		//如果获取不到日志则logTmp := <-f.logChan管道阻塞了
		select {
		case logTmpList := <-logChanList:
			row := len(logTmpList.PlatformLog)
			//保存日志到数据库
			db := model.Db.CreateInBatches(logTmpList.PlatformLog, row)
			if db.Error != nil {
				logger.Error("db.CreateInBatches error:", db.Error)
				logger.Error("db.CreateInBatches RowsAffected:", db.RowsAffected)
			}

			//把所有日志保存文件
			if logTmpList.LogConfigInfo.IsFile {
				if isSpiteLog {
					time.Sleep(time.Second * 2)
				}
				for _, log := range logTmpList.PlatformLog {
					_, err := fmt.Fprintf(logFile.fileObj, log.Content)
					if err != nil {
						logger.Error("写入log文件失败，err=", err)
						time.Sleep(time.Second * 1)
						logFile.fileObj.Close()
						logTmpList.LogConfigInfo.openLogFile()
						_, err := fmt.Fprintf(logFile.fileObj, log.Content)
						if err != nil {
							logger.Error("关闭log后再打开log文件，再次写入log文件失败，放弃写入，err=", err)
						} else {
							logRowCount++
						}

					} else {
						logRowCount++
					}
					if logRowCount >= logTmpList.LogConfigInfo.MaxLogLines {

						if !isSpiteLog {
							isSpiteLog = true
							newFile, err := spiteFile(logFile.fileObj, logTmpList.LogConfigInfo.LogPath) //切割错误的日志文件
							isSpiteLog = false
							if err != nil {
								logger.Error("切割错误日志文件出错，err=", err)
								return
							}
							logFile.fileObj = newFile
						}

					}

				}

			}

			if row >= 1000 {
				if checkSize(false, logFile.fileObj) {
					if !isSpiteLog {
						isSpiteLog = true
						newFile, err := spiteFile(logFile.fileObj, logTmpList.LogConfigInfo.LogPath) //切割错误的日志文件
						isSpiteLog = false
						if err != nil {
							logger.Error("切割错误日志文件出错，err=", err)
							return
						}
						logFile.fileObj = newFile
					}

				}
			}

			//把所有错误保存文件
			if logTmpList.LogConfigInfo.IsError {
				if isSpiteErrorLog {
					time.Sleep(time.Second * 2)
				}
				tmpCount := 0
				for _, log := range logTmpList.PlatformLog {
					if log.LevelInt >= ERROR {

						//如果记录的日志级别大于ERROR级别，我还要在err日志文件中再记录一遍
						_, err := fmt.Fprintf(logFile.errFileObj, log.Content)
						if err != nil {
							logger.Error("写入error.log文件失败，err=", err)
							time.Sleep(time.Second * 1)
							logFile.errFileObj.Close()
							logTmpList.LogConfigInfo.openEerrorLogFile()
							_, err := fmt.Fprintf(logFile.errFileObj, log.Content)
							if err != nil {
								logger.Error("关闭错误日志log文件后再次打开，再次写入log文件失败，放弃写入，err=", err)
								continue
							} else {
								errorLogRowCount++
								tmpCount++
							}
						} else {
							errorLogRowCount++
							tmpCount++
						}
						if errorLogRowCount >= logTmpList.LogConfigInfo.MaxErrorLogLines {
							if !isSpiteErrorLog {
								isSpiteErrorLog = true
								newFile, err := spiteFile(logFile.errFileObj, logTmpList.LogConfigInfo.LogPath) //切割错误的日志文件
								isSpiteErrorLog = false
								if err != nil {
									logger.Error("切割错误日志文件出错，err=", err)
									return
								}
								logFile.errFileObj = newFile
							}

						}
					}
				}
				if tmpCount >= 1000 {
					if checkSize(true, logFile.errFileObj) {
						if !isSpiteErrorLog {
							isSpiteErrorLog = true
							newFile, err := spiteFile(logFile.errFileObj, logTmpList.LogConfigInfo.LogPath) //切割错误的日志文件
							isSpiteErrorLog = false
							if err != nil {
								logger.Error("切割错误日志文件出错，err=", err)
								return
							}
							logFile.errFileObj = newFile
						}

					}
				}
			}

		case logTmpTrace := <-logChanTrace:
			if logTmpTrace.LogConfigInfo.UserName != "" {
				if conn, ok := connClient[logTmpTrace.LogConfigInfo.UserName]; ok {
					err := conn.WriteMessage(websocket.TextMessage, []byte(logTmpTrace.PlatformLog.Content))
					if err != nil {
						logger.Error("write message error,err=", err, "userName=", logTmpTrace.LogConfigInfo.UserName)
					}
				} else {
					logger.Error("connClient not found, userName=", logTmpTrace.LogConfigInfo.UserName)
				}
			}

		default:
			//如果从管道取不出来日志则休息500毫秒
			time.Sleep(time.Millisecond * 500)

		}

	}

}

// 检测条数是否达到用户设置的最大值，如果是检测文件大小则1分钟检测一次
func (config *LogConfig) checkLog() {
	for {
		//检测条数是否达到用户设置的最大值，如果是检测文件大小则1分钟检测一次
		if checkSize(false, logFile.fileObj) {
			if !isSpiteLog {
				isSpiteLog = true
				newFile, err := spiteFile(logFile.fileObj, config.LogPath) //切割正常的日志文件
				isSpiteLog = false
				if err != nil {
					logger.Error("切割日志文件出错", err)
					return
				}
				logFile.fileObj = newFile
			}

		}
		//判断是否有按日切割文件
		if config.LogDaily && time.Now().Hour() == 23 && time.Now().Minute() == 59 {
			time.Sleep(time.Second * time.Duration(59-time.Now().Second()))
			if !isSpiteLog {
				isSpiteLog = true
				newFile, err := spiteFile(logFile.fileObj, config.LogPath) //切割错误的日志文件
				isSpiteLog = false
				if err != nil {
					logger.Error("切割日志文件出错，err=", err)
					return
				}
				logFile.fileObj = newFile
			}

		}
		//判断实时追踪日志是否有过有效期，如果有则清理掉
		for _, id := range TraceIdList {
			compareTime(id)
		}
		if len(TraceIdList) > 0 {
			IsTrace = true
		} else {
			IsTrace = false
		}
		time.Sleep(time.Minute * 1)
	}
}

// 比较两个时间是否过期，如果过期则删除
func compareTime(id string) {
	// 获取当前时间
	now := time.Now()
	logger.Info("now:", now, TraceIdPeriod[id], !now.Before(TraceIdPeriod[id]), id)
	if !now.Before(TraceIdPeriod[id]) {
		for i, traceId := range TraceIdList {
			if id == traceId {
				TraceIdList = append(TraceIdList[:i], TraceIdList[i+1:]...)
				logger.Info("删除过期的实时日志id:", id)
				delete(TraceIdPeriod, id)
				delete(connClient, id)
				break
			}
		}
	}
}

func (config *LogConfig) checkEerrorLog() {
	for {
		//检测条数是否达到用户设置的最大值，如果是检测文件大小则1分钟检测一次
		if checkSize(true, logFile.errFileObj) {
			if !isSpiteErrorLog {
				isSpiteErrorLog = true
				newFile, err := spiteFile(logFile.errFileObj, config.ErrorLogPath) //切割错误的日志文件
				isSpiteErrorLog = false
				if err != nil {
					logger.Error("切割错误日志文件出错，err=", err)
					return
				}
				logFile.errFileObj = newFile
			}

		}
		if config.ErrorDaily && time.Now().Hour() == 23 && time.Now().Minute() == 59 {
			time.Sleep(time.Second * time.Duration(59-time.Now().Second()))
			if !isSpiteErrorLog {
				isSpiteErrorLog = true
				newFile, err := spiteFile(logFile.errFileObj, config.ErrorLogPath) //切割错误的日志文件
				isSpiteErrorLog = false
				if err != nil {
					logger.Error("切割错误日志文件出错，err=", err)
					return
				}
				logFile.errFileObj = newFile
			}

		}

		time.Sleep(time.Minute * 1)
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 实时追踪日志,通过websocket发送给客户端
// 启动websockt服务
func websocketListen() {
	logger.Info("启动websocket服务...")
	//http.HandleFunc("/", index)
	http.HandleFunc("/logWebsocket", handleConnections)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

type Client struct {
	conn *websocket.Conn
}

var clients = make(map[*Client]string)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}


// 处理连接请求
func handleConnections(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	sn := r.URL.Query().Get("sn")

	logger.Info("跟踪用户或设备:", sn)
	connClient[sn] = conn
	flag := false
	for _, s := range TraceIdList {
		if s == sn {
			flag = true
			break
		}
	}
	if !flag {
		TraceIdList = append(TraceIdList, sn)
	}

	for {
		mt, message, err := conn.ReadMessage()

		if err != nil {
			logger.Info("客户端断开连接，clientId:", sn, "sn:", sn)
			delete(connClient, sn)
			break
		}

		logger.Info("收到实时日志：", mt, string(message))
	}
}

// 非格式化打印日志
func (config *LogConfig) console(levelInt int, level string, format interface{}, a ...interface{}) {

	if levelInt >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprintln(format, abc[:len(abc)-1])

		funcName, fileName, lineNo := getInfo(3)

		if config.IsConsole {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s", now, level, fileName, funcName, lineNo, msg)
			message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s", now, level, fileName, funcName, lineNo, msg)
			config.writeLog(message, level, levelInt)
		} else {
			message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s", now, level, fileName, funcName, lineNo, msg)
			config.writeLog(message, level, levelInt)
		}
	}
}

// 未知级别日志
func (config *LogConfig) Unknown(format interface{}, a ...interface{}) {
	//write(DEBUG, format, a...)
	//config.console(UNKNOWN, "Unknown", format, a...)
	if UNKNOWN >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprint(format, abc[:len(abc)-1])
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s", now, "Unknown", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Cyan(message)
			config.writeLog(message, "Unknown", UNKNOWN)
		} else {
			config.writeLog(message, "Unknown", UNKNOWN)
		}
	}
}

// 用户级调试日志
func (config *LogConfig) Debug(format interface{}, a ...interface{}) {
	//write(DEBUG, format, a...)
	//config.console(DEBUG, "Debug", format, a...)
	if DEBUG >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprint(format, abc[:len(abc)-1])
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s", now, "Debug", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.White(message)
			config.writeLog(message, "Debug", DEBUG)
		} else {
			config.writeLog(message, "Debug", DEBUG)
		}
	}
}

// 用户级信息日志
func (config *LogConfig) Info(format interface{}, a ...interface{}) {
	//write(INFO, format, a...)
	//config.console(INFO, "Info", format, a...)
	if INFO >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprint(format, abc[:len(abc)-1])
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s", now, "Info", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Blue(message)
			config.writeLog(message, "Info", DEBUG)
		} else {
			config.writeLog(message, "Info", DEBUG)
		}
	}
}

// 用户级警告
func (config *LogConfig) Warn(format interface{}, a ...interface{}) {
	//write(WARN, format, a...)
	//config.console(WARN, "Warn", format, a...)
	if WARN >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprintf("%s%s", format, abc[:len(abc)-1])
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s", now, "Warn", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Yellow(message)
			config.writeLog(message, "Warn", WARN)
		} else {
			config.writeLog(message, "Warn", WARN)
		}
	}
}

// 用户级错误
func (config *LogConfig) Error(format interface{}, a ...interface{}) {

	//write(ERROR, format, a...)
	//config.console(ERROR, "Error", format, a...)
	if ERROR >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprint(format, abc[:len(abc)-1])
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s", now, "Error", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Red(message)
			config.writeLog(message, "Error", ERROR)
		} else {
			config.writeLog(message, "Error", ERROR)
		}
	}
}

// 致命错误
func (config *LogConfig) Fatal(format interface{}, a ...interface{}) {

	//write(FATAL, format, a...)
	//config.console(FATAL, "Fatal", format, a...)
	if FATAL >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprint(format, abc[:len(abc)-1])
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s", now, "Fatal", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Magenta(message)
			config.writeLog(message, "Fatal", FATAL)
		} else {
			config.writeLog(message, "Fatal", FATAL)
		}
	}
}

// 系统级危险，比如权限出错，访问异常等
func (config *LogConfig) Crit(format interface{}, a ...interface{}) {
	//write(CRIT, format, a...)
	//config.console(CRIT, "Crit", format, a...)
	if CRIT >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprint(format, abc[:len(abc)-1])
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s", now, "Crit", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Magenta(message)
			config.writeLog(message, "Crit", CRIT)
		} else {
			config.writeLog(message, "Crit", CRIT)
		}
	}
}

// 系统级警告，比如数据库访问异常，配置文件出错等
func (config *LogConfig) Alrt(format interface{}, a ...interface{}) {

	//write(ALRT, format, a...)
	//config.console(ALRT, "Alrt", format, a...)
	if ALRT >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprint(format, abc[:len(abc)-1])
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s", now, "Alrt", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Magenta(message)
			config.writeLog(message, "Alrt", ALRT)
		} else {
			config.writeLog(message, "Alrt", ALRT)
		}
	}
}

// 系统级紧急，比如磁盘出错，内存异常，网络不可用等
func (config *LogConfig) Emer(format interface{}, a ...interface{}) {

	//write(EMER, format, a...)
	//config.console(EMER, "Emer", format, a...)
	if EMER >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprint(format, abc[:len(abc)-1])
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s", now, "Emer", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Magenta(message)
			config.writeLog(message, "Emer", EMER)
		} else {
			config.writeLog(message, "Emer", EMER)
		}
	}
}

// 入侵警告
func (config *LogConfig) Invade(format interface{}, a ...interface{}) {
	//write(INVADE, format, a...)
	//config.console(INVADE, "Invade", format, a...)
	if INVADE >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprint(format, abc[:len(abc)-1])
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s", now, "Invade", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Green(message)
			config.writeLog(message, "Invade", INVADE)
		} else {
			config.writeLog(message, "Invade", INVADE)
		}
	}
}

// 格式化打印日志
func (config *LogConfig) consoleFormat(levelInt int, level, format string, a ...interface{}) {

	if levelInt >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)
		funcName, fileName, lineNo := getInfo(2)

		if config.IsConsole {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s \n", now, level, fileName, funcName, lineNo, msg)
			message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s \n", now, level, fileName, funcName, lineNo, msg)
			config.writeLog(message, level, levelInt)
		} else {
			message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s \n", now, level, fileName, funcName, lineNo, msg)
			config.writeLog(message, level, levelInt)
		}
	}
}

// Unknown Format
func (config *LogConfig) Uf(format string, a ...interface{}) {
	//writeFormat(UNKNOWN, format, a...)
	//config.consoleFormat(UNKNOWN, "Unknown", format, a...)
	if UNKNOWN >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s \n", now, "Unknown", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Cyan(message)
			config.writeLog(message, "Unknown", UNKNOWN)
		} else {
			config.writeLog(message, "Unknown", UNKNOWN)
		}
	}
}

// Debug Format
func (config *LogConfig) Df(format string, a ...interface{}) {
	//writeFormat(DEBUG, format, a...)
	//config.consoleFormat(DEBUG, "Debug", format, a...)
	if DEBUG >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s \n", now, "Debug", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.White(message)
			config.writeLog(message, "Debug", DEBUG)
		} else {
			config.writeLog(message, "Debug", DEBUG)
		}
	}
}

// Info Format
func (config *LogConfig) If(format string, a ...interface{}) {

	//writeFormat(INFO, format, a...)
	//config.consoleFormat(INFO, "Info", format, a...)
	if INFO >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s \n", now, "Info", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Blue(message)
			config.writeLog(message, "Info", INFO)
		} else {
			config.writeLog(message, "Info", INFO)
		}
	}
}

// Warn Format
func (config *LogConfig) Wf(format string, a ...interface{}) {
	//writeFormat(WARN, format, a...)
	//config.consoleFormat(WARN, "Warn", format, a...)
	if WARN >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s \n", now, "Warn", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Yellow(message)
			config.writeLog(message, "Warn", WARN)
		} else {
			config.writeLog(message, "Warn", WARN)
		}
	}
}

// Error Format
func (config *LogConfig) Errf(format string, a ...interface{}) {

	//writeFormat(ERROR, format, a...)
	//config.consoleFormat(ERROR, "Error", format, a...)
	if ERROR >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s \n", now, "Error", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Red(message)
			config.writeLog(message, "Error", ERROR)
		} else {
			config.writeLog(message, "Error", ERROR)
		}
	}
}

// Fatal Format
func (config *LogConfig) Ff(format string, a ...interface{}) {

	//writeFormat(FATAL, format, a...)
	//config.consoleFormat(FATAL, "Fatal", format, a...)
	if FATAL >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s \n", now, "Fatal", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Magenta(message)
			config.writeLog(message, "Fatal", FATAL)
		} else {
			config.writeLog(message, "Fatal", FATAL)
		}
	}
}

// Crit Format  系统级危险，比如权限出错，访问异常等
func (config *LogConfig) Cf(format string, a ...interface{}) {
	//writeFormat(CRIT, format, a...)
	//config.consoleFormat(CRIT, "Crit", format, a...)
	if CRIT >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s \n", now, "Crit", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Magenta(message)
			config.writeLog(message, "Crit", CRIT)
		} else {
			config.writeLog(message, "Crit", CRIT)
		}
	}
}

// Alrt Format  系统级警告，比如数据库访问异常，配置文件出错等
func (config *LogConfig) Af(format string, a ...interface{}) {

	//writeFormat(ALRT, format, a...)
	//config.consoleFormat(ALRT, "Alrt", format, a...)
	if ALRT >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s \n", now, "Alrt", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Magenta(message)
			config.writeLog(message, "Alrt", ALRT)
		} else {
			config.writeLog(message, "Alrt", ALRT)
		}
	}
}

// 系统级紧急，比如磁盘出错，内存异常，网络不可用等
func (config *LogConfig) EmerF(format string, a ...interface{}) {
	//writeFormat(EMER, format, a...)
	//config.consoleFormat(EMER, "Emer", format, a...)
	if EMER >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s \n", now, "Emer", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Magenta(message)
			config.writeLog(message, "Emer", EMER)
		} else {
			config.writeLog(message, "Emer", EMER)
		}
	}
}

// Invade Format，入侵警告
func (config *LogConfig) InvadeF(format string, a ...interface{}) {
	//writeFormat(INVADE, format, a...)
	//config.consoleFormat(INVADE, "Invade", format, a...)
	if INVADE >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)
		funcName, fileName, lineNo := getInfo(2)
		message := fmt.Sprintf("%s [%s] [%s=>%s:%d] %s \n", now, "Invade", fileName, funcName, lineNo, msg)
		if config.IsConsole {
			color.Green(message)
			config.writeLog(message, "Invade", INVADE)
		} else {
			config.writeLog(message, "Invade", INVADE)
		}
	}
}
