package store

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/pkg/conf"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/url"
	"os"
	"time"
)

var db *gorm.DB

// Setup 初始化数据库连接
func Setup() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&time_zone=%s",
		conf.C.Database.Username,
		conf.C.Database.Password,
		conf.C.Database.Host,
		conf.C.Database.Port,
		conf.C.Database.Name,
		url.QueryEscape("'Asia/Shanghai'"))

	// gorm的日志记录器
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
		Logger: newLogger.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Panicf("Open DB connect failed %s", err.Error())
	}
	logrus.Info("Init DB successfully")
}

//GetConn 获取db connect
func GetConn() *gorm.DB {
	return db
}

