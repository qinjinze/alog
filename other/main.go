package main

import (
	"github.com/alog"
	"fmt"
	"time"
)

//测试自己写的日志库
func main() {
	err := alog.InitLogConfig()
	fmt.Println("err", err)
	id := 1001
	name := "loge"

	for {
		//other.Hello()
		//alog.Debug("这是一条Debug日志%s", name)
		//alog.Info("这是一条Info日志%s", name)
		//alog.Warn("这是一条Warning日志%s", name)
		alog.Error("这是一条Error日志", id, name)
		//alog.Fatal("这是一条Fatal日志", name)
		alog.Info("非格式化打印", id, name)

		time.Sleep(2 * time.Second)
	}

}
