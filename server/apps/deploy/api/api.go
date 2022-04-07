package api

import (
	"github.com/gin-gonic/gin"
	"ops-app-deploy/frame/log"
	"ops-app-deploy/frame/model"
	"ops-app-deploy/frame/response"
	"ops-app-deploy/frame/router"
)

func init() {
	//g1 := router.New("admin", "/demo/form", jwt.JWTAuthMiddleware())
	log.Log.Infof("加载router: %s", "/deploy/app")
	gd := router.New("deploy", "/deploy/app")
	gd.GET("/ping", "", func(c *gin.Context) {
		response.SucessResp(c).SetCode(200).SetBtype(model.Buniss_Other).SetMsg("测试使用连接").WriteJsonExit()
	})

}
