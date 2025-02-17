package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/tunnel"

type CreateTunnelResponse struct {
	Tunnel tunnel.Tunnel `json:"tunnel"`
}

type FindUserAllTunnelsResponse struct {
	Tunnels []tunnel.Tunnel `json:"tunnels"`
	Total   int64             `json:"total"`
}