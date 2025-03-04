package server

import (
	"context"

	pb "github.com/flipped-aurora/gin-vue-admin/server/grpc/protocol/ncx_client"
	tunnel "github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ConfigInfoServer struct {
	pb.UnimplementedConfigInfoServer
}

func (n *ConfigInfoServer) AskConfigInfo(cx context.Context, cfg *pb.ConfigInfoReq) (*pb.ConfigInfoResp, error) {
	var resp pb.ConfigInfoResp
	services := make([]*pb.ServiceConfig, 0)

	// 从数据库中获取配置信息
	tunnels, err := tunnel.ServiceGroupApp.TunnelService.FindUserAllTunnelsByToken(cfg.GetToken())
	if err != nil {
		resp.Code = 1;
	}

	for _, t := range tunnels {
		service := &pb.ServiceConfig{}
		service.ServiceName = t.Tunnelname
		service.ServiceIp = t.Localip
		service.ServicePort = int32(t.Localport)
		service.ProxyPort = int32(t.Remoteport)
		services = append(services, service)
	}

	resp.Code = 0;
	resp.ServiceConfigs = services
	return &resp, nil
}
