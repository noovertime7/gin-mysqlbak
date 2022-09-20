package database

import (
	"fmt"
	"github.com/noovertime7/gin-mysqlbak/conf"
	infoLog "github.com/noovertime7/mysqlbak/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

var (
	gormDB *gorm.DB
	dbOnce sync.Once
	err    error
)

func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		initDB()
	})
	return gormDB
}

func initDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.GetStringConf("mysql", "user"),
		conf.GetStringConf("mysql", "password"),
		conf.GetStringConf("mysql", "host"),
		conf.GetStringConf("mysql", "port"),
		conf.GetStringConf("mysql", "dbname"),
	)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         255,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	gormDB, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	//连接池最大允许的空闲连接数
	sqlDB, err := gormDB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(100)
	//设置最大连接数
	sqlDB.SetMaxIdleConns(20)
	//设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(60 * time.Second)
	infoLog.Logger.Info("初始化数据库成功")
}
