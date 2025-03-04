package initialize

import (
	"fmt"
	"net"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	pb "github.com/flipped-aurora/gin-vue-admin/server/grpc/protocol/ncx_client"
	server "github.com/flipped-aurora/gin-vue-admin/server/grpc/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func initGrpcServer() (*grpc.Server, error) {
	grpcsrv := grpc.NewServer()

	// 注册grpc服务群
	configinfoserver := &server.GrpcServerGroupApp.ConfigInfoServer
	pb.RegisterConfigInfoServer(grpcsrv, configinfoserver)

	return grpcsrv, nil
}

func StartGrpcServer() error {
	grpcServerCfg := global.GVA_CONFIG.GrpcServer
	listen, err := net.Listen(grpcServerCfg.Type, fmt.Sprintf(":%d", grpcServerCfg.Port))
	if err != nil {
		global.GVA_LOG.Error("grpc server listen failed", zap.Error(err))
		return err
	}

	go func() {
		if err := global.GVA_GRPCSERVER.Serve(listen); err != nil {
			global.GVA_LOG.Error("grpc server start failed", zap.Error(err))
		}
	}()

	return err
}

func GrpcServer() error {
	grpcServer, err := initGrpcServer()
	if err != nil {
		return err
	}

	global.GVA_GRPCSERVER = grpcServer

	return err
}
