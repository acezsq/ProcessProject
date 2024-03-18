package main

import (
	_ "github.com/acezsq/project-api/api"
	"github.com/acezsq/project-api/config"
	"github.com/acezsq/project-api/router"
	srv "github.com/acezsq/project-common"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)

	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, nil)

}
