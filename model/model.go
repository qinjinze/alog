package model

import (
	"github.com/alog/utils"
	_ "github.com/alog/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wonderivan/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 设备日志
type DeviceLog struct {
	Id int64 `json:"DeviceLogId" gorm:"primaryKey"` //Id
	Level      string    `orm:"size(10);NULL;description(日志等级名称)" json:"Level" gorm:"comment:日志等级名称"`           //
	LevelInt   int       `orm:"size(2);NULL;description(日志等级)" json:"LevelInt" gorm:"comment:日志等级"`             //
	Function   string    `orm:"size(100);NULL;description(功能模块)" json:"Function" gorm:"comment:功能模块"`           //
	RequestId  string    `orm:"size(30);NULL;description(每次请求的唯一id)" json:"RequestId" gorm:"comment:每次请求的唯一id"` //
	Token      string    `orm:"size(255);NULL;description(设备登录Token)" json:"Token" gorm:"comment:设备登录Token"`      //
	UserName   string    `orm:"size(255);NULL;description(设备唯一id或登录名)" json:"UserName" gorm:"comment:设备唯一id或登录名"`         //
	Content    string    `orm:"size(255);NULL;description(日志内容)" json:"Content" gorm:"comment:日志内容"`        //
	LogTime    time.Time `orm:"size(15);NULL;description(日志生成时间)" json:"LogTime" gorm:"comment:日志生成时间"`
	CreateTime time.Time `orm:"size(6);auto_now_add;type(datetime);description(插入表中时间)" json:"CreateTime" gorm:"autoCreateTime;comment:插入表中时间"`
}

// 管理后台日志
type PlatformLog struct {
	Id         int64     `json:"PlatformLogId" gorm:"primaryKey"`                                               //Id
	Level      string    `orm:"size(10);NULL;description(日志等级名称)" json:"Level" gorm:"comment:日志等级名称"`           //
	LevelInt   int       `orm:"size(2);NULL;description(日志等级)" json:"LevelInt" gorm:"comment:日志等级"`             //
	Function   string    `orm:"size(100);NULL;description(功能模块)" json:"Function" gorm:"comment:功能模块"`           //
	RequestId  string    `orm:"size(30);NULL;description(每次请求的唯一id)" json:"RequestId" gorm:"comment:每次请求的唯一id"` //
	Token      string    `orm:"size(255);NULL;description(用户登录Token)" json:"Token" gorm:"comment:用户登录Token"`      //
	UserName   string    `orm:"size(255);NULL;description(账户)" json:"UserName" gorm:"comment:用户登录唯一id"`         //
	Content    string    `orm:"size(255);NULL;description(日志内容)" json:"Content" gorm:"comment:日志内容"`        //
	LogTime    time.Time `orm:"size(15);NULL;description(日志生成时间)" json:"LogTime" gorm:"comment:日志生成时间"`
	CreateTime time.Time `orm:"size(6);auto_now_add;type(datetime);description(插入表中时间)" json:"CreateTime" gorm:"autoCreateTime;comment:插入表中时间"`
}

// C端用户日志
type UserLog struct {
	Id         int64     `json:"UserLogId" gorm:"primaryKey"`                                                   //Id
	Level      string    `orm:"size(10);NULL;description(日志等级名称)" json:"Level" gorm:"comment:日志等级名称"`           //
	LevelInt   int       `orm:"size(2);NULL;description(日志等级)" json:"LevelInt" gorm:"comment:日志等级"`             //
	Function   string    `orm:"size(100);NULL;description(功能模块)" json:"Function" gorm:"comment:功能模块"`           //
	RequestId  string    `orm:"size(30);NULL;description(每次请求的唯一id)" json:"RequestId" gorm:"comment:每次请求的唯一id"` //
	Token      string    `orm:"size(255);NULL;description(用户登录Token)" json:"Token" gorm:"comment:用户登录Token"`      //
	UserName   string    `orm:"size(255);NULL;description(账户)" json:"UserName" gorm:"comment:用户登录唯一id"`         //
	Content    string    `orm:"size(255);NULL;description(日志内容)" json:"Content" gorm:"comment:日志内容"`        //
	LogTime    time.Time `orm:"size(15);NULL;description(日志生成时间)" json:"LogTime" gorm:"comment:日志生成时间"`
	CreateTime time.Time `orm:"size(6);auto_now_add;type(datetime);description(插入表中时间)" json:"CreateTime" gorm:"autoCreateTime;comment:插入表中时间"`
}

var Db *gorm.DB

func InitModel(dbUserName, dbPassword, dbAddr, dbPort, dbName string) {
	utils.InitinitConfigFile()
	if utils.G_mysql_dbUserName == "" {
		utils.G_mysql_dbUserName = dbUserName
	}
	if utils.G_mysql_dbPassword == "" {
		utils.G_mysql_dbPassword = dbPassword
	}
	if utils.G_mysql_addr == "" {
		utils.G_mysql_addr = dbAddr
	}
	if utils.G_mysql_port == "" {
		utils.G_mysql_port = dbPort
	}
	if utils.DbName == "" {
		utils.DbName = dbName
	}
	logger.Info("数据库连接信息：", utils.G_mysql_dbUserName, utils.G_mysql_dbPassword, utils.G_mysql_addr, utils.G_mysql_port, utils.DbName)
	var err error
	dsn := utils.G_mysql_dbUserName + ":" + utils.G_mysql_dbPassword + "@tcp(" + utils.G_mysql_addr + ":" + utils.G_mysql_port + ")/" + utils.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败，err=" + err.Error())
		return
	}

	//连接成功
	sqlDB, err := Db.DB()
	if err != nil {
		logger.Fatal("数据库关闭失败", err)
		panic("数据库连接失败，err=" + err.Error())
		return
	}
	sqlDB.SetMaxIdleConns(30)    //设置连接池，空闲
	sqlDB.SetMaxOpenConns(10000) //设置打开最大连接

	// 迁移 schema ，自动创建表结构
	err = Db.AutoMigrate(&DeviceLog{}, &PlatformLog{}, &UserLog{})
	if err != nil {
		logger.Error("数据库迁移失败", err)
	}

}
