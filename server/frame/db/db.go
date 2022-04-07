package db

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"ops-app-deploy/frame/config"
	"ops-app-deploy/frame/log"
	"sync"
	"xorm.io/xorm"
)

var (
	instance *xorm.Engine
	once     sync.Once
)

//初始化数据操作 driver为数据库类型
func Instance() *xorm.Engine {
	once.Do(func() {
		configuration := config.Configuration
		driverName := configuration.Database.Driver
		logger := log.Log

		if len(driverName) <= 0 {
			driverName = "mysql"
		}

		// 建立数据库连接
		engine, err := xorm.NewEngine(driverName, configuration.Database.DbUrl)
		if err != nil {
			logger.Errorf("数据库连接错误:%v", err.Error())
			panic("数据库连接错误")
			return
		}

		// 验证数据库是否连通
		err = engine.Ping()
		if err != nil {
			logger.Errorf("数据库连接错误:%v", err.Error())
			panic("数据库连接错误")
			return
		}

		// 设置日志级别
		if configuration.Database.Debug {
			engine.SetLogger(log.LoggerForXorm())
			engine.ShowSQL(true)
			engine.SetLogLevel(0)
		}

		logger.Info("数据库连接已经建立.....")

		instance = engine

	})
	return instance
}
