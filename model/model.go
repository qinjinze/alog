package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"github.com/qinjinze/alog/utils"
	//"alog/utils"
	_ "github.com/qinjinze/alog/utils"
	//_ "alog/utils"
	"github.com/beego/beego/orm"
	"github.com/wonderivan/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormloger "gorm.io/gorm/logger"
	"time"
)

// 实时监控的历史轨迹日志
type DeviceLog struct {
	Id         int64     `json:"DeviceLogId" gorm:"primaryKey"`                                                   //Id
	Sn         string    `orm:"size(30);description(编号)" json:"Sn"`                                               //
	Level      string    `orm:"size(10);NULL;description(日志等级名称)" json:"Level" gorm:"comment:日志等级名称"`             //
	LevelInt   int       `orm:"size(2);NULL;description(日志等级)" json:"LevelInt" gorm:"comment:日志等级"`               //
	Function   string    `orm:"size(100);NULL;description(功能模块)" json:"Function" gorm:"comment:功能模块"`             //
	RequestId  string    `orm:"size(30);NULL;description(每次请求的唯一id)" json:"RequestId" gorm:"comment:每次请求的唯一id"`   //
	Token      string    `orm:"size(255);NULL;description(设备登录Token)" json:"Token" gorm:"comment:设备登录Token"`      //
	UserName   string    `orm:"size(255);NULL;description(设备唯一id或登录名)" json:"UserName" gorm:"comment:设备唯一id或登录名"` //
	Content    string    `orm:"size(1024);NULL;description(日志内容)" json:"Content" gorm:"comment:日志内容;type:text"`   //
	LogTime    time.Time `orm:"size(6);NULL;description(日志生成时间)" json:"LogTime" gorm:"comment:日志生成时间"`
	CreateTime time.Time `orm:"size(6);auto_now_add;type(datetime);description(插入表中时间)" json:"CreateTime" gorm:"autoCreateTime;comment:插入表中时间"`
	Seller     string    `orm:"size(255);NULL;description(商家简称：网店名称)" json:"Seller" gorm:"comment:商家简称：网店名称"` //
	SellerId   string    `orm:"size(100);NULL;description(商家编号)" json:"SellerId" gorm:"comment:商家编号"`         //
}

// 管理后台日志
type PlatformLog struct {
	Id                int64     `json:"PlatformLogId" gorm:"primaryKey"`                                               //Id
	Level             string    `orm:"size(10);NULL;description(日志等级名称)" json:"Level" gorm:"comment:日志等级名称"`           //
	LevelInt          int       `orm:"size(2);NULL;description(日志等级)" json:"LevelInt" gorm:"comment:日志等级"`             //
	Function          string    `orm:"size(100);NULL;description(功能模块)" json:"Function" gorm:"comment:功能模块"`           //
	RequestId         string    `orm:"size(30);NULL;description(每次请求的唯一id)" json:"RequestId" gorm:"comment:每次请求的唯一id"` //
	Token             string    `orm:"size(255);NULL;description(用户登录Token)" json:"Token" gorm:"comment:用户登录Token"`    //
	UserName          string    `orm:"size(255);NULL;description(账户)" json:"UserName" gorm:"comment:用户登录唯一id"`         //
	Content           string    `orm:"size(1024);NULL;description(日志内容)" json:"Content" gorm:"comment:日志内容;type:text"` //
	LogTime           time.Time `orm:"size(6);NULL;description(日志生成时间)" json:"LogTime" gorm:"comment:日志生成时间"`
	CreateTime        time.Time `orm:"size(6);auto_now_add;type(datetime);description(插入表中时间)" json:"CreateTime" gorm:"autoCreateTime;comment:插入表中时间"`
	Seller            string    `orm:"size(255);NULL;description(商家简称：网店名称)" json:"Seller" gorm:"comment:商家简称：网店名称"` //
	SellerId          string    `orm:"size(100);NULL;description(商家编号)" json:"SellerId" gorm:"comment:商家编号"`         //
	Page              string    `orm:"size(10);NULL;description(页面名称)" json:"Page" gorm:"comment:页面名称"`              //
	Api               string    `orm:"size(10);NULL;description(请求接口)" json:"Api" gorm:"comment:请求接口"`
	Ip                string    `orm:"size(150);NULL;description(登录IP)" json:"Ip"`                //登录IP
	Browser           string    `orm:"size(30);NULL;description(浏览器)" json:"Browser"`             //浏览器，小程序，APP
	BrowserVersion    string    `orm:"size(30);NULL;description(浏览器版本)" json:"BrowserVersion"`    // 浏览器，小程序，APP 版本
	System            string    `orm:"size(30);NULL;description(操作系统)" json:"System"`             //操作系统
	SystemVersion     string    `orm:"size(30);NULL;description(系统版本)" json:"SystemVersion"`      //系统版本
	SystemLanguage    string    `orm:"size(30);NULL;description(系统语言)" json:"SystemLanguage"`     //系统语言
	Model             string    `orm:"size(30);NULL;description(硬件型号)" json:"Model"`              // 手机、平板、电脑型号
	Width             int       `orm:"size(30);NULL;description(窗口宽度)" json:"Width"`              // 窗口宽度
	Height            int       `orm:"size(30);NULL;description(窗口高度)" json:"Height"`             // 窗口高度
	MobileePixelRatio string    `orm:"size(30);NULL;description(窗口像素比)" json:"MobileePixelRatio"` // 手机、平板、电脑像素比，可要获取物理像素分辨率，我们还需要使用devicePixelRatio。 devicePixelRatio表示表示实际像素与逻辑像素的比例

}

// C端用户日志
type UserLog struct {
	Id               int64     `json:"UserLogId" gorm:"primaryKey"`                                                   //Id
	Level            string    `orm:"size(10);NULL;description(日志等级名称)" json:"Level" gorm:"comment:日志等级名称"`           //
	LevelInt         int       `orm:"size(2);NULL;description(日志等级)" json:"LevelInt" gorm:"comment:日志等级"`             //
	Function         string    `orm:"size(100);NULL;description(功能模块)" json:"Function" gorm:"comment:功能模块"`           //
	RequestId        string    `orm:"size(30);NULL;description(每次请求的唯一id)" json:"RequestId" gorm:"comment:每次请求的唯一id"` //
	Token            string    `orm:"size(255);NULL;description(用户登录Token)" json:"Token" gorm:"comment:用户登录Token"`    //
	UserName         string    `orm:"size(255);NULL;description(账户)" json:"UserName" gorm:"comment:用户登录唯一id"`         //
	Content          string    `orm:"size(1024);NULL;description(日志内容)" json:"Content" gorm:"comment:日志内容;type:text"` //
	LogTime          time.Time `orm:"size(6);NULL;description(日志生成时间)" json:"LogTime" gorm:"comment:日志生成时间"`
	CreateTime       time.Time `orm:"size(6);auto_now_add;type(datetime);description(插入表中时间)" json:"CreateTime" gorm:"autoCreateTime;comment:插入表中时间"`
	Seller           string    `orm:"size(255);NULL;description(商家简称：网店名称)" json:"Seller" gorm:"comment:商家简称：网店名称"` //
	SellerId         string    `orm:"size(100);NULL;description(商家编号)" json:"SellerId" gorm:"comment:商家编号"`         //
	Page             string    `orm:"size(10);NULL;description(页面名称)" json:"Page" gorm:"comment:页面名称"`              //
	Api              string    `orm:"size(10);NULL;description(请求接口)" json:"Api" gorm:"comment:请求接口"`
	Latitude         string    `orm:"size(30);NULL;description(用户使用程序时所在经度)" json:"Latitude"`   // 用户经度
	Longitude        string    `orm:"size(30);NULL;description(用户使用程序时所在纬度)" json:"Longitude"`  // 用户纬度
	Ip               string    `orm:"size(150);NULL;description(登录IP)" json:"Ip"`               //登录IP
	Browser          string    `orm:"size(100);NULL;description(浏览器)" json:"Browser"`           //浏览器，小程序，APP
	BrowserVersion   string    `orm:"size(100);NULL;description(浏览器版本)" json:"BrowserVersion"`  // 浏览器，小程序，APP 版本
	System           string    `orm:"size(100);NULL;description(操作系统)" json:"System"`           //操作系统
	SystemVersion    string    `orm:"size(100);NULL;description(系统版本)" json:"SystemVersion"`    //系统版本
	SystemLanguage   string    `orm:"size(100);NULL;description(系统语言)" json:"SystemLanguage"`   //系统语言
	Model            string    `orm:"size(30);NULL;description(硬件型号)" json:"Model"`             // 手机、平板、电脑型
	Width            int       `orm:"size(30);NULL;description(窗口宽度)" json:"Width"`             // 窗口宽度
	Height           int       `orm:"size(30);NULL;description(窗口高度)" json:"Height"`            // 窗口高度
	DevicePixelRatio string    `orm:"size(30);NULL;description(窗口像素比)" json:"DevicePixelRatio"` // 手机、平板、电脑像素比
}

// 用户信息
type User struct {
	Id       int64  `json:"UserLogId" gorm:"primaryKey"`                                           //Id
	UserName string `orm:"size(255);NULL;description(账户)" json:"UserName" gorm:"comment:用户登录唯一id"` //
	Password string `orm:"size(255);NULL;description(密码)" json:"Password" gorm:"comment:用户登录密码"`   //
}

var Db *gorm.DB
var err error

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
		if dbName == "" {
			dbName = "log" + time.Now().Format("20060102")
		}
		utils.DbName = dbName
	}
	logger.Info("数据库连接信息：", utils.G_mysql_dbUserName, utils.G_mysql_dbPassword, utils.G_mysql_addr, utils.G_mysql_port, utils.DbName)
	// 注册数据库驱动和数据库DSN
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", utils.G_mysql_dbUserName+":"+utils.G_mysql_dbPassword+"@tcp("+utils.G_mysql_addr+")/")

	// 创建数据库
	sql := "CREATE DATABASE IF NOT EXISTS " + utils.DbName + " CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci"
	_, err = orm.NewOrm().Raw(sql).Exec()
	if err != nil {
		fmt.Println("Error creating database:", err)
		return
	}

	//newLogger := gormloger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	gormloger.Config{
	//		SlowThreshold:              time.Second,   // Slow SQL threshold（慢速SQL阈值）
	//		LogLevel:                   gormloger.Info, // Log level（日志级别）
	//		IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
	//		ParameterizedQueries:      true,           // Don't include params in the SQL log（忽略记录器的ErrRecordNotFound错误）
	//		Colorful:                  true,          // Disable color(禁用颜色)设置彩色打印
	//	},
	//)


	dsn := utils.G_mysql_dbUserName + ":" + utils.G_mysql_dbPassword + "@tcp(" + utils.G_mysql_addr + ":" + utils.G_mysql_port + ")/" + utils.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置

	}), &gorm.Config{
		//Logger: newLogger,
		Logger: gormloger.Default.LogMode(gormloger.Info),
	})
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
	sqlDB.SetMaxIdleConns(30)  //设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) //设置打开最大连接

	// 迁移 schema ，自动创建表结构
	err = Db.AutoMigrate(&DeviceLog{}, &PlatformLog{}, &UserLog{}, &User{})
	if err != nil {
		logger.Error("数据库迁移失败", err)
	}

}
