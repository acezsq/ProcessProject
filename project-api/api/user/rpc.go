package user

import (
	"github.com/acezsq/project-api/config"
	"github.com/acezsq/project-common/discovery"
	"github.com/acezsq/project-common/logs"
	loginServiceV1 "github.com/acezsq/project-user/pkg/service/login.service.v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
)

var LoginServiceClient loginServiceV1.LoginServiceClient

func InitRpcUserClient() {
	etcdRegister := discovery.NewResolver(config.C.EtcdConfig.Addrs, logs.LG)
	resolver.Register(etcdRegister)
	conn, err := grpc.Dial("etcd:///user", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	LoginServiceClient = loginServiceV1.NewLoginServiceClient(conn)
}
