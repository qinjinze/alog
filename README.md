简体中文:

1.支持在linux、windows终端以及IDEA终端打印彩色日志，其它vscode等没有测试过

2.实现日志写入数据库，这里实现的写入mysql,postgresql，对中大型项目不推荐使mysql存储日志，下一个版本会增加写入ES数据库

3.实现写入文件根据大小和写入日志行数进行切割，可以自定如果不自定会使用默认参数

4.并且实现websocket实时跟踪某个用户或者设备日志，web客户端需要自己实现，如有需求本人也可以提供

如果需要生成日志文件或者把数据写入数据库、文件、websocket发送到客户端，需要将配置文件复制到自己项目中，配置文件目前位于alog/conf/app.conf

5.日志等级配置

| 等级  | 配置      | 释义                       | 控制台颜色 |
| --- | ------- | ------------------------ | ----- |
| 0   | UNKNOWN | 未知                       | 绿蓝色   |
| 1   | TRACE   | 用户级基本输出                  | 灰白色   |
| 2   | DEBG    | 用户级调试                    | 灰白色   |
| 3   | INFO    | 用户级重要                    | 天蓝色   |
| 4   | WARN    | 用户级警告                    | 黄色    |
| 5   | EROR    | 用户级错误                    | 红色    |
| 6   | FATAL   | 用户级基本输出                  | 粉色    |
| 7   | CRIT    | 系统级危险，比如权限出错，访问异常等       | 粉色    |
| 8   | ALRT    | 系统级警告，比如数据库访问异常，配置文件出错等  | 粉色    |
| 9   | EMER    | 系统级紧急，比如磁盘出错，内存异常，网络不可用等 | 粉色    |
| 10  | INVADE  | 黑客入侵                     | 绿色    |

go get github.com/qinjinze/alog

 或
go install github.com/qinjinze/alog@latest

cd alog/cmd

go run main.go

English:

1. Supports printing color logs on Linux, Windows, and IDEA terminals, and has not been tested on other vscodes

2. Implement log writing to the database, which involves writing to MySQL and PostgreSQL. It is not recommended to use MySQL to store logs for medium to large projects，The next version will add log writing to the ES database

3. Implement file splitting based on size and number of log lines written, which can be customized. If not customized, default parameters will be used

4. And achieve real-time tracking of a user or device log through websocket. The web client needs to implement it themselves, and I can also provide it if needed

If you need to generate a log file or write data to a database, file, or websocket to send to the client, you need to copy the configuration file to your own project. The configuration file is currently located in the directory/conf/app. conf

5. Log level configuration

| Level | Configuration | Definition                                                                                           | Color      |
| ----- | ------------- | ---------------------------------------------------------------------------------------------------- | ---------- |
| 0     | UNKNOWN       | unknown                                                                                              | Green blue |
| 1     | TRACE         | User level basic<br> output                                                                          | Grey white |
| 2     | DEBG          | User level debugging                                                                                 | Grey white |
| 3     | INFO          | User level importance                                                                                | Sky blue   |
| 4     | WARN          | User level warning                                                                                   | yellow     |
| 5     | EROR          | User level error                                                                                     | red        |
| 6     | FATAL         | User level basic<br> output                                                                          | Pink       |
| 7     | CRIT          | System level hazards,<br> such as permission errors, abnormal access, etc                            | Pink       |
| 8     | ALRT          | System level<br> warnings, such as abnormal database access, configuration file errors, etc          | Pink       |
| 9     | EMER          | System level<br> emergencies, such as disk errors, memory anomalies, network unavailability,<br> etc | Pink       |
| 10    | INVADE        | Hacker intrusion                                                                                     | green      |

go get github.com/qinjinze/alog 

or  
go install github.com/qinjinze/alog@latest

cd alog/cmd

go run main.go