package main

import (
	srv "github.com/acezsq/project-common"
	"github.com/acezsq/project-user/config"
	"github.com/acezsq/project-user/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)
	// grpc服务注册
	gc := router.RegisterGrpc()
	// grpc服务注册到etcd中
	router.RegisterEtcdServer()

	stop := func() {
		gc.Stop()
	}
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, stop)

}
