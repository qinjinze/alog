package common

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wonderivan/logger"
	"gorm.io/driver/mysql"
	//"net/url"

	//logger "alogalog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	logorm "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// 定义配置结构体
type DbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type RedisConfig struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

type MqttConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	Postgres map[string]DbConfig `yaml:"postgres"` // 支持多数据库配置
	Mysql    map[string]DbConfig `yaml:"mysql"`    // 支持多数据库配置
	Redis    RedisConfig         `yaml:"redis"`    // 支持多redis配置
	Mqtt     MqttConfig          `yaml:"mqtt"`     // 支持多mqtt配置
}

var (
	dbs  = make(map[string]*gorm.DB) // 使用map存储多个数据库实例
	Conf *Config
	err  error
)

func LoadConfig(path string) error {
	v := viper.New()
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&Conf); err != nil {
		return err
	}
	return nil
}

func InitLogs() {
	// 设置日志配置
	logger.SetLogger(`{"File":{"filename":"./logs/app.log","level":"TRAC","daily":true,"maxlines":1000000,"maxsize":104857600,"maxdays":-1,"append":true,"permit":"0777"}}`)
	//log := logger.LogConfig{
	//	IsConsole:    true,
	//	Level:        "debug",
	//	Color:        true,
	//	IsFile:       true,
	//	LogName:      "log.log",
	//	LogPath:      "./",
	//	IsError:      false,
	//	ErrorLogPath: "errorlog.log",
	//	ErrorLogName: "./errorlog.log",
	//	SaveDbType:   "mysql",
	//	LogType:      "platform", //user,platform、device
	//	UserName:     "",         //用户名
	//	RequestId:    "",         //每次API请求生成唯一ID
	//	Page:         "",         //页面
	//	Api:          "",         //api名称
	//	Function:     "",         //方法函数
	//	Seller:       "",
	//	SellerId:     "",
	//	Token:        "",
	//	DbHost:       "127.0.0.1",
	//	DbPort:       "5432",
	//	DbUserName:   "postgres",
	//	DbPassword:   "Ktz&*1217",
	//	DbName:       "error_log",
	//}
	//
	//log.InitLogConfig()

}

// 加载配置文件
func init() {

	// 初始化日志
	InitLogs()

	err = LoadConfig("../conf/config.yaml")
	if err != nil {
		logger.Error("配置文件加载失败: %v", err)
		//os.Exit(1)
		panic(err)
	}

	// 初始化mysql
	//mysqlInit()
	postgresInit()

}

func mysqlInit() {
	logger.Info("初始化mysql数据库连接...%+v", Conf)
	for key, dbConfig := range Conf.Mysql {
		logger.Info("key: %s, dbConfig: %v", key, dbConfig)
		//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		//	dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Database, dbConfig.Port)
		// 密码特殊字符转义
		//encodedPassword := url.QueryEscape(dbConfig.Password)
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s",
			dbConfig.Username,
			dbConfig.Password, // 使用转义后的密码
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Database,
		)
		newLogger := logorm.New(
			//log.New(log.Writer(), "\r\n", log.LstdFlags), // io writer
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logorm.Config{
				//SlowThreshold: time.Second, // 慢 SQL 阈值
				SlowThreshold: 50 * time.Millisecond, // 慢 SQL 阈值
				LogLevel:      logorm.Info,           // 日志级别
				//LogLevel:                  logorm.Warn, // 日志级别
				IgnoreRecordNotFoundError: false, // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,  // 启用彩色打印
			},
		)
		logger.Info("dsn: %s", dsn)
		// 创建数据库连接[1](@ref)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			PrepareStmt: true, // 开启预编译语句缓存
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //全局禁用复数表名：
			},
			Logger: newLogger,
		})
		if err != nil {
			logger.Error("mysql数据库[%s]连接失败: %v", key, err)
			continue // 单个数据库失败不影响其他连接
		}

		// 配置连接池参数[5](@ref)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)                  // 默认空闲连接数
		sqlDB.SetMaxOpenConns(100)                 // 默认最大连接数
		sqlDB.SetConnMaxLifetime(time.Hour)        // 连接最大存活时间
		sqlDB.SetConnMaxIdleTime(10 * time.Minute) // 空闲连接超时

		// 测试连通性
		if err := sqlDB.Ping(); err != nil {
			logger.Error("mysql数据库[%s]心跳检测失败: %v", key, err)
			continue
		}

		dbs[key] = db // 存入全局map
		logger.Info("mysql数据库[%s]连接成功 | 地址:%s:%d", key, dbConfig.Host, dbConfig.Port)
	}

}
func postgresInit() {
	logger.Info("初始化postgres数据库连接...%+v", Conf)
	for key, dbConfig := range Conf.Postgres {
		logger.Info("key: %s, dbConfig: %v", key, dbConfig)
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Database, dbConfig.Port)
		// 密码特殊字符转义
		//encodedPassword := url.QueryEscape(dbConfig.Password)
		newLogger := logorm.New(
			//log.New(log.Writer(), "\r\n", log.LstdFlags), // io writer
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logorm.Config{
				//SlowThreshold: time.Second, // 慢 SQL 阈值
				SlowThreshold: 50 * time.Millisecond, // 慢 SQL 阈值
				LogLevel:      logorm.Info,           // 日志级别
				//LogLevel:                  logorm.Warn, // 日志级别
				IgnoreRecordNotFoundError: false, // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,  // 启用彩色打印
			},
		)
		logger.Info("dsn: %s", dsn)
		// 创建数据库连接
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			PrepareStmt: true, // 开启预编译语句缓存
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //全局禁用复数表名：
			},
			Logger: newLogger,
		})
		if err != nil {
			logger.Error("postgres数据库[%s]连接失败: %v", key, err)
			continue // 单个数据库失败不影响其他连接
		}

		// 配置连接池参数
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)                  // 默认空闲连接数
		sqlDB.SetMaxOpenConns(100)                 // 默认最大连接数
		sqlDB.SetConnMaxLifetime(time.Hour)        // 连接最大存活时间
		sqlDB.SetConnMaxIdleTime(10 * time.Minute) // 空闲连接超时

		// 测试连通性
		if err := sqlDB.Ping(); err != nil {
			logger.Error("postgres数据库[%s]心跳检测失败: %v", key, err)
			continue
		}

		dbs[key] = db // 存入全局map
		logger.Info("postgres数据库[%s]连接成功 | 地址:%s:%d", key, dbConfig.Host, dbConfig.Port)
	}

}

// 获取数据库实例
func GetDB(name string) *gorm.DB {
	if db, ok := dbs[name]; ok {
		return db
	} else {
		logger.Error("数据库[%s]未初始化", name)
		return nil
	}
}
