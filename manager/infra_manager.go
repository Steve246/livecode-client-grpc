package manager

import (
	"enigmacamp.com/lopei_grpc_cnlt/config"
	"enigmacamp.com/lopei_grpc_cnlt/service"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type InfraManager interface {
	LopeiClientConn() service.LopeiPaymentClient
}

type infraManager struct {
	lopeiClient service.LopeiPaymentClient
	cfg         config.Config
}

func (i *infraManager) initGrpcClient() {
	conn, err := grpc.Dial(i.cfg.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	client := service.NewLopeiPaymentClient(conn)
	i.lopeiClient = client
	fmt.Println("Grpc Client Connected")
}
func (i *infraManager) LopeiClientConn() service.LopeiPaymentClient {
	return i.lopeiClient
}
func NewInfraManager(config config.Config) InfraManager {
	infra := infraManager{
		cfg: config,
	}
	infra.initGrpcClient()
	return &infra
}
