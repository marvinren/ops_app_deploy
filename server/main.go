package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	_ "ops-app-deploy/apps/deploy/api"
	"ops-app-deploy/apps/deploy/model/app"
	"ops-app-deploy/frame/config"
	"ops-app-deploy/frame/db"
	"ops-app-deploy/frame/server"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

var g errgroup.Group


func main() {
	gin.SetMode("debug")

	//配置文件是写在config的init()里直接执行的，这里直接获取配置信息
	//配置文件默认读取./config.yml文件，具体配置信息可查看config/config.go文件定义
	//这里缺少对配置信息的检查
	appConfig := config.Configuration
	if appConfig == nil {
		fmt.Printf("参数错误")
		return
	}

	dbIns := db.Instance()
	dbIns.Sync2(app.DeployApp{})

	deployApp := server.New(appConfig.App.Name, appConfig.Server.Address)
	//server的start里使用的了errgroup来实现多个服务同时运行的支持，目前这里只运行了一个服务
	//具体方法可查看https://blog.csdn.net/ma2595162349/article/details/109435100
	deployApp.Start(g)

	if err := g.Wait(); err != nil {
		fmt.Println(err.Error())
	}

	//因为gin已经进入协程来运行一个或多个服务
	//所以这里使用了系统signal捕获来决定整个程序是否持续运行
	//这也是一种优雅退出的方式
	var state int32 = 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

EXIT:
	for {
		fmt.Println("Server正在运行...")
		sig := <-sc
		fmt.Printf("获取到信号[%s]\n", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			atomic.StoreInt32(&state, 0)
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	fmt.Println("服务退出")
	time.Sleep(time.Second)
	os.Exit(int(atomic.LoadInt32(&state)))
}