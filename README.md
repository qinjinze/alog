# alog


日志在linux、windows终端以及IDEA终端实现彩色打印，以及实现日志写入数据库、写入文件、并且实现websocket实时跟踪某个用户或者设备日志。

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

如果需要生成日志文件或者把数据写入数据库、文件、websocket发送到客户端，需要将配置文件复制到自己项目中，配置文件目前位于alog/conf/app.conf

go get github.com/wenyusun/alog
or
go install github.com/wenyusun/alog@latest

cd alog/cmd

go run main.go

