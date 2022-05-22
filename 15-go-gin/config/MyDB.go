package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var DbHelper *gorm.DB
var err error

func init() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/test_go?charset=utf8mb4&parseTime=True&loc=Local"
	DbHelper, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "t_",                              // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // 1、默认gorm会在表名后面加复数，设置这个参数后，即可关闭
			//NoLowerCase:   true,                              // skip the snake_casing of names
			//NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		), //加SQL的执行Log
	})
	if DbHelper == nil || err != nil {
		fmt.Println("数据库连接错误: ", err.Error())
		//log.Fatal("DB初始化错误:", err)
		ShutdownServer(err)
		return
	}
	sqlDB, _ := DbHelper.DB()
	sqlDB.SetMaxIdleConns(10)           // 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxOpenConns(100)          // 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置了连接可复用的最大时间。
	//defer sqlDB.Close()
}
