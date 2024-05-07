package main

import (
	"github.com/qinjinze/alog"
	//"alog"
	"time"
)

// 测试自己写的日志库
func main() {
	// 初始化日志配置,可以不先填写，使用默认配置
	log := alog.LogConfig{
		IsConsole:    true,
		Level:        "debug",
		Color:        true,
		IsFile:       true,
		LogName:      "log.log",
		LogPath:      "./",
		IsError:      false,
		ErrorLogPath: "errorlog.log",
		ErrorLogName: "./errorlog.log",
		SaveDbType:   "mysql",
		LogType:      "platform", //user,platform、device
		UserName:     "",         //用户名
		RequestId:    "",         //每次API请求生成唯一ID
		Page:         "",         //页面
		Api:          "",         //api名称
		Function:     "",         //方法函数
		Seller:       "",
		SellerId:     "",
		Token:        "",
		DbHost:       "127.0.0.1",
		DbPort:       "3306",
		DbUserName:   "",
		DbPassword:   "",
		DbName:       "",
	}

	log.InitLogConfig()
	id := 0
	//name := "loge"
	//alog.TimeFormat = "2006-01-02 15:04:05.999999999 -0700 MST"
	alog.IsColor = true
	//utils.IsConsole = false
	for {
		id++
		//other.Hello()
		//alog.Debug("这是一条Debug日志%s", name)
		//alog.Info("这是一条Info日志%s", name)
		//alog.Warn("这是一条Warning日志%s", name)
		//alog.Error("这是一条Error日志", id, name)
		//alog.Fatal("这是一条Fatal日志", name)
		//alog.Unknown("亮色字体", 1, time.Now())
		//fmt.Println()
		//alog.Debug("白色字体", 2, time.Now())
		//fmt.Println()
		//alog.Info("蓝色字体", 3, time.Now())
		//fmt.Println()
		//alog.Warn("黄色字体", 4, time.Now())
		//fmt.Println()
		//alog.Error("红色字体", 5, time.Now())
		//fmt.Println()
		//alog.Fatal("粉色字体", 6, time.Now())
		//fmt.Println()
		//alog.Invade("绿色字体", 6, time.Now())
		//fmt.Println()
		//alog.Uf("亮色字体%d, %s", 1, time.Now())
		//alog.Df("白色字体%d, %s", 2, time.Now())
		//alog.If("蓝色字体%d, %s", 3, time.Now())
		//alog.Wf("黄色字体%d, %s", 4, time.Now())
		//alog.Errf("红色字体%d, %s", 5, time.Now())
		//alog.Ff("粉色字体%d, %s", 6, time.Now())
		//alog.InvadeF("绿色字体%d, %s", 6, time.Now())
		//content, err := Hello()
		//alog.If("格式化打印id=%d,name=%s", content, err)
		//logger.Info("测试日志库", id, name)
		//log.Info("写入文件中", id, time.Now())
		log.UserName = "admin"
		log.RequestId = "123456"
		log.Page = "index"
		log.Api = "/login"
		log.Function = "login"
		log.LogType = "platform"
		log.Uf("亮色字体%d, %s", id, time.Now())
		//log.Df("白色字体%d, %s", 2, time.Now())
		//log.If("蓝色字体%d, %s", 3, time.Now())
		//log.Wf("黄色字体%d, %s", 4, time.Now())
		//log.Errf("红色字体%d, %s", 5, time.Now())
		//log.Ff("粉色字体%d, %s", 6, time.Now())
		//log.Unknown("亮色字体", 1, time.Now())
		////fmt.Println()
		//log.Debug("白色字体", 2, time.Now())
		////fmt.Println()
		//log.Info("蓝色字体", 3, time.Now())
		////fmt.Println()
		//log.Warn("黄色字体", 4, time.Now())
		////fmt.Println()
		//log.Error("红色字体", 5, time.Now())
		////fmt.Println()
		//log.Fatal("粉色字体", 6, time.Now())
		////fmt.Println()
		//log.Invade("绿色字体", 6, time.Now())
		//fmt.Println()
		//id++
		//log.Info("写入文件中", id, name)
		//time.Sleep(1 * time.Second)
		//break
		//time.Sleep(3 * time.Second)

	}
	//logger.Info("测试日志库")
	//log:=	alog.LogConfig{
	//		IsConsole:     false,
	//		Level:         1,
	//		Color:         false,
	//		IsFile:        false,
	//		FileName:      false,
	//		FilePath:      false,
	//		IsError:       false,
	//		ErrorFilePath: false,
	//		ErrorFileName: false,
	//	}
	//	log.

}
