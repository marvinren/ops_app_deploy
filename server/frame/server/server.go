package server

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"net/http"
	"ops-app-deploy/frame/log"
	"ops-app-deploy/frame/router"
	"syscall"
	"time"
)

// Server: HTTP服务的结构体
type Server struct {
	ServerName     string        //服务名称
	Addr           string        //监听地址端口
	ServerRoot     string        //静态资源文件夹
	Handler        *gin.Engine   //HTTP Handler
	ReadTimeout    time.Duration //读取超时时间
	WriteTimeout   time.Duration //写入超时时间
	MaxHeaderBytes int           //http头大小设置
}

// 已注册HTTP服务列表
var servers = make(map[string]*Server, 0)


func GetServer(name string) *Server {
	return servers[name]
}

// 创建服务
func New(name, addr string, middleware ...gin.HandlerFunc) *Server {
	var s Server
	s.WriteTimeout = 60 * time.Second
	s.ReadTimeout = 60 * time.Second
	s.Addr = addr
	s.ServerName = name
	s.MaxHeaderBytes = 1 << 20
	s.Handler = gin.New()
	s.Handler.Use(middleware...)
	//注册路由
	logHandler := log.LoggerToFile(s.ServerName)
	if len(router.GroupList) > 0 {
		for _, group := range router.GroupList {
			if group.ServerName == name || group.ServerName == "default" {
				//增加日志handler
				handlerFunctions := append(group.Handlers, logHandler)
				grp := s.Handler.Group(group.RelativePath, handlerFunctions...)
				for _, r := range group.Router {
					//增加router加是日志
					if r.Method == "ANY" {
						grp.Any(r.RelativePath, r.HandlerFunc...)
					} else {
						grp.Handle(r.Method, r.RelativePath, r.HandlerFunc...)
					}

				}
			}
		}
	}
	//日志
	servers[name] = &s
	return &s
}


//启动服务
func (s *Server) Start(g errgroup.Group) {

	server := &http.Server{
		Addr:           s.Addr,
		Handler:        s.Handler,
		ReadTimeout:    s.ReadTimeout,
		WriteTimeout:   s.WriteTimeout,
		MaxHeaderBytes: s.MaxHeaderBytes,
	}

	//log.Printf("[%v]Server listen: %v Actual pid is %d", s.ServerName, s.Addr, syscall.Getpid())
	log.Log.Infof("[%v]Server listen: %v Actual pid is %d", s.ServerName, s.Addr, syscall.Getpid())

	g.Go(func() error {
		return server.ListenAndServe()
	})
}