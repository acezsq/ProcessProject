package user

import (
	"github.com/acezsq/project-api/router"
	"github.com/gin-gonic/gin"
	"log"
)

type RouterUser struct {
}

func init() {
	log.Println("init user router")
	ru := &RouterUser{}
	router.Register(ru)
}

func (*RouterUser) Route(r *gin.Engine) {
	// 初始化grpc客户端连接
	h := New()
	r.POST("/project/login/getCaptcha", h.getCaptcha)
}
