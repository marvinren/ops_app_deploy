package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"io"
	"ops-app-deploy/frame/config"
	"os"
	"path"
	"sync"
	"time"
	xlog "xorm.io/xorm/log"
)

var Log *log.Logger
var once sync.Once

func init() {
	config := config.Configuration
	Logger(config.App.Name)
}

func Logger(serverName string) *log.Logger {
	once.Do(func() {
		if Log != nil {
			return
		}

		configuration := config.Configuration

		//log的实例化
		logClient := log.New()
		// 日志打印到指定的目录
		var logPath = configuration.Logger.Path

		if len(logPath) == 0 {
			logPath = "/tmp/log/app"
		}

		if len(serverName) == 0 {
			serverName = "default"
		}

		if err := os.MkdirAll(logPath, 0777); err != nil {
			fmt.Println(err.Error())
		}
		fileName := path.Join(logPath, serverName+".log")
		//日志文件
		if _, err := os.Stat(fileName); err != nil {
			if _, err := os.Create(fileName); err != nil {
				fmt.Println(err.Error())
			}
		}

		//禁止logrus的输出
		src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("err", err)
		}
		// 设置日志输出方式
		if configuration.Logger.Stdout {
			writers := []io.Writer{
				src,
				os.Stdout}
			//同时写文件和屏幕
			fileAndStdoutWriter := io.MultiWriter(writers...)
			if err == nil {
				log.SetOutput(fileAndStdoutWriter)
			} else {
				log.Info("failed to log to file.")
			}
		} else {
			logClient.Out = src
		}

		//设置日志级别
		tmp := configuration.Logger.Level
		logClient.SetLevel(log.Level(tmp))

		//设置滚动日志策略
		//apiLogPath := "gin-api.log"
		logWriter, err := rotatelogs.New(
			fileName+".%Y-%m-%d-%H-%M.log",
			rotatelogs.WithLinkName(fileName),         // 生成软链，指向最新日志文件
			rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
			rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
		)
		writeMap := lfshook.WriterMap{
			log.InfoLevel:  logWriter,
			log.FatalLevel: logWriter,
			log.DebugLevel: logWriter, // 为不同级别设置不同的输出目的
			log.WarnLevel:  logWriter,
			log.ErrorLevel: logWriter,
			log.PanicLevel: logWriter,
		}
		lfHook := lfshook.NewHook(writeMap, &log.JSONFormatter{})
		logClient.AddHook(lfHook)

		//设置日志格式
		logClient.SetFormatter(&log.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})

		Log = logClient
	})

	return Log
}

func LoggerToFile(serverName string) gin.HandlerFunc {

	logClient := Logger(serverName)

	return func(c *gin.Context) {

		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		// 这里是指定日志打印出来的格式。分别是状态码，执行时间,请求ip,请求方法,请求路由(等下我会截图)
		logClient.Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
	}
}

func LoggerForXorm() *xlog.SimpleLogger {
	var logger *xlog.SimpleLogger

	configuration := config.Configuration

	var logPath = configuration.Logger.Path

	if len(logPath) == 0 {
		logPath = "/tmp/log/app"
	}
	if err := os.MkdirAll(logPath, 0777); err != nil {
		Log.Error("日志文件创建错误", err.Error())
	}
	fileName := path.Join(logPath, "sql.log")
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			Log.Error("日志文件创建错误", err.Error())
		}
	}



	//获取日志文件路径
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	writers := []io.Writer{
		src,
		os.Stdout}
	//同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if err == nil {
		logger = xlog.NewSimpleLogger(fileAndStdoutWriter)
	} else {
		log.Info("failed to log to file.")
	}

	return logger


}