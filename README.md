**简体中文**:  
以往日志库总有这样那样痛点，所以自己写了个日志库，希望能帮到大家。

1.支持在linux、windows终端以及IDEA终端打印彩色日志，其它vscode等没有测试过

2.支持日志写入数据库，大多数公司使用mysql,默认支持mysql，没有特殊要求，推荐使用postgresql或ES数据库。如果使用默认之外的数据库，请修改model.go和file.go文件，file.go文件693行或750行

3.支持写入文件根据大小和写入日志行数进行切割或按天切割，可以自定如果不自定会使用默认参数

4.支持websocket实时跟踪某个用户或者设备日志，web客户端需要自己实现，如有需求可以找本人 email:310508138@qq.com  
如果是本机测试，可以直接访问：http://127.0.0.1:12345/logWebsocket?sn="admin"  
sn为设备号或用户名，此处根据自己需要是否决定做校验，默认不做校验，可以直接访问。需要校验可以在file.go文件中修改1075行修改校验逻辑,跟踪有效期24个小时。  

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

linux终端：

![linux终端](http://150.158.34.122/image/linux.png)

windows终端：

![windows终端](http://150.158.34.122/image/windows.png)

IDEA终端：

![IDEA终端](http://150.158.34.122/image/idea.png)

Websocket实时跟踪某个用户或者设备日志，web客户端需要自己实现，如有需求可以找本人：email:310508138@qq.com

![Websocket](http://150.158.34.122/image/websocket.png)

go get github.com/qinjinze/alog  
或  
go install github.com/qinjinze/alog@latest  
以上命令如果使用代理下载不一定获取最新版本，建议使用git clone git@github.com:qinjinze/alog.git  
cd alog/cmd

go run main.go              

<br>
                 
**English**:  
There have always been pain points in the past, so I wrote a log library myself, hoping to help everyone.

1. Supports printing color logs on Linux, Windows, and IDEA terminals, and has not been tested on other vscodes

2. Support log writing to the database. Most companies use MySQL, which is supported by default without any special requirements. It is recommended to use PostgreSQL or ES databases. If using a database other than the default, please modify the model. go and file. go files, file. go file at line 693 or 750  

3. Implement file write splitting based on size and number of log lines written or daily splitting, which can be customized. If not customized, default parameters will be used  

4. Support websocket for real-time tracking of user or device logs. The web client needs to implement it themselves. If there is a need, you can contact me via email: 310508138@qq.com  
   If it is a local test, you can directly access: http://127.0.0.1:12345/logWebsocket?sn= "Admin"  
   SN is the device number or username, and it is determined whether to perform verification based on personal needs. By default, verification is not performed and can be accessed directly. To verify, you can modify the verification logic in line 1075 of the file. go file. Set the tracking validity period to 24 hours     

If you need to generate a log file or write data to a database, file, or websocket to send to the client, you need to copy the configuration file to your own project. The configuration file is currently located in the directory/conf/app. conf

5. Log level configuration

| Level | Configuration | Definition                                                                                           | Color      |
| ----- | ------------- |------------------------------------------------------------------------------------------------------| ---------- |
| 0     | UNKNOWN       | unknown                                                                                              | Green blue |
| 1     | TRACE         | User level basic output                                                                              | Grey white |
| 2     | DEBG          | User level debugging                                                                                 | Grey white |
| 3     | INFO          | User level importance                                                                                | Sky blue   |
| 4     | WARN          | User level warning                                                                                   | yellow     |
| 5     | EROR          | User level error                                                                                     | red        |
| 6     | FATAL         | User level basic output                                                                              | Pink       |
| 7     | CRIT          | System level hazards, such as permission errors, abnormal access, etc                            | Pink       |
| 8     | ALRT          | System level warnings, such as abnormal database access, configuration file errors, etc          | Pink       |
| 9     | EMER          | System level  emergencies, such as disk errors, memory anomalies, network unavailability,etc | Pink       |
| 10    | INVADE        | Hacker intrusion                                                                                     | green      |

go get github.com/qinjinze/alog   
or  
go install github.com/qinjinze/alog@latest  
If using a proxy to download the above command may not necessarily obtain the latest version, it is recommended to use git clone git@github.com : qinjinze/logs. git  

cd alog/cmd

go run main.go
