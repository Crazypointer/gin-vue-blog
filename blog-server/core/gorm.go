package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"server/global"
	"time"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		log.Println("未配grom，取消mysql链接！")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info) //打印所有
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) //只打印错误
	}
	global.MysqlLog = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatalf(fmt.Sprintf("[%s] mysql链接失败！", dsn))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               //最大空闲链接数
	sqlDB.SetMaxOpenConns(100)              //最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //链接最大复用时间，不能超过mysql的wait_timeout
	return db
}
