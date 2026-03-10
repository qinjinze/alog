package model

import (
	pgDb "alog/init/common"
	"github.com/wonderivan/logger"
	"gorm.io/gorm"
	logorm "gorm.io/gorm/logger"
)

var ErrorLogDb *gorm.DB
var PlatformLogDb *gorm.DB
var UserLogDb *gorm.DB
var DeviceLogDb *gorm.DB

func initDeviceLog() {

	DeviceLogDb = pgDb.GetDB("device_log")

	// 初始化时关闭日志
	DeviceLogDb.Logger = logorm.Default.LogMode(logorm.Silent)
	logger.Info("连接数据库成功", DeviceLogDb.Name())

	var tables = []interface{}{
		&DeviceLog{},
	}

	// 创建表，如果表已存在，则会抛出错误（除非你使用 db.Set("gorm:table_options", "IF NOT EXISTS")）
	DeviceLogDb.Set("gorm:table_options", "IF NOT EXISTS")
	for i, _ := range tables {
		if !DeviceLogDb.Migrator().HasTable(tables[i]) {
			DeviceLogDb.Migrator().CreateTable(tables[i])
		}
		DeviceLogDb.AutoMigrate(
			tables[i],
		)
	}
	//initDeviceLog()
	//initErrorLog()
	// 恢复日志级别
	DeviceLogDb.Logger = logorm.Default.LogMode(logorm.Info)
}
func initErrorLog() {

	ErrorLogDb = pgDb.GetDB("error_log")

	// 初始化时关闭日志
	ErrorLogDb.Logger = logorm.Default.LogMode(logorm.Silent)
	logger.Info("连接数据库成功", ErrorLogDb.Name())

	var tables = []interface{}{
		&ErrorLog{},
	}

	// 创建表，如果表已存在，则会抛出错误（除非你使用 db.Set("gorm:table_options", "IF NOT EXISTS")）
	ErrorLogDb.Set("gorm:table_options", "IF NOT EXISTS")
	for i, _ := range tables {
		if !ErrorLogDb.Migrator().HasTable(tables[i]) {
			ErrorLogDb.Migrator().CreateTable(tables[i])
		}
		ErrorLogDb.AutoMigrate(
			tables[i],
		)
	}
	//initDeviceLog()
	//initErrorLog()
	// 恢复日志级别
	ErrorLogDb.Logger = logorm.Default.LogMode(logorm.Info)
}

func initPlatformLog() {

	PlatformLogDb = pgDb.GetDB("platform_log")

	// 初始化时关闭日志
	PlatformLogDb.Logger = logorm.Default.LogMode(logorm.Silent)
	logger.Info("连接数据库成功", PlatformLogDb.Name())

	var tables = []interface{}{
		&PlatformLog{},
	}

	// 创建表，如果表已存在，则会抛出错误（除非你使用 db.Set("gorm:table_options", "IF NOT EXISTS")）
	PlatformLogDb.Set("gorm:table_options", "IF NOT EXISTS")
	for i, _ := range tables {
		if !PlatformLogDb.Migrator().HasTable(tables[i]) {
			PlatformLogDb.Migrator().CreateTable(tables[i])
		}
		PlatformLogDb.AutoMigrate(
			tables[i],
		)
	}

	// 恢复日志级别
	PlatformLogDb.Logger = logorm.Default.LogMode(logorm.Info)
}

func initUserLog() {

	UserLogDb = pgDb.GetDB("api_log")

	// 初始化时关闭日志
	UserLogDb.Logger = logorm.Default.LogMode(logorm.Silent)
	logger.Info("连接数据库成功", UserLogDb.Name())

	var tables = []interface{}{
		&UserLog{},
	}

	// 创建表，如果表已存在，则会抛出错误（除非你使用 db.Set("gorm:table_options", "IF NOT EXISTS")）
	UserLogDb.Set("gorm:table_options", "IF NOT EXISTS")
	for i, _ := range tables {
		if !UserLogDb.Migrator().HasTable(tables[i]) {
			UserLogDb.Migrator().CreateTable(tables[i])
		}
		UserLogDb.AutoMigrate(
			tables[i],
		)
	}
	//initDeviceLog()
	//initErrorLog()
	// 恢复日志级别
	UserLogDb.Logger = logorm.Default.LogMode(logorm.Info)
}

// 初始化数据库连接
func init() {

	logger.Info("连接数据库，注册model 建表")
	initPlatformLog()
	initDeviceLog()
	initErrorLog()
	initUserLog()

	logger.Info("init model success")

}
