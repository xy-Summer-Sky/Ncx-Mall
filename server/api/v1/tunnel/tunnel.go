package tunnel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tunnel"
	tunnelReq "github.com/flipped-aurora/gin-vue-admin/server/model/tunnel/request"
	tunnelRes "github.com/flipped-aurora/gin-vue-admin/server/model/tunnel/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TunnelApi struct {}

// CreateTunnel
// @Tags      Tunnel
// @Summary   用户申请创建通道
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body    tunnelReq.CreateTunnel "用户ID，隧道参数"
// @Success   200   {object}  response.Response{msg=string}  "隧道创建成功"
// @Router    /user/changePassword [post]
func (t *TunnelApi) CreateTunnel(c *gin.Context) {
	var req tunnelReq.CreateTunnel
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return 
	}

	err = utils.Verify(req, utils.CreateTunnelVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	// 创建通道
	tunnel := &tunnel.Tunnel{
		Tunnelname: req.Tunnelname,
		Tunneltype: req.Tunneltype,
		Tunnelpoint: req.Tunnelpoint,
		Localip: req.Localip,
		Localport: req.Localport,
		Remtoeip: req.Remoteip,
		Remoteport: req.Remoteport,
		Token: req.Token,
	}

	tunnel.Userid = utils.GetUserID(c)
	tunnelReturn, err := tunnelService.CreateTunnel(*tunnel)

	if err != nil {
		global.GVA_LOG.Error("创建隧道失败！", zap.Error(err))
		response.FailWithDetailed(tunnelRes.CreateTunnelResponse{Tunnel: tunnelReturn}, "创建隧道失败", c)
		return
	}

	response.OkWithDetailed(tunnelRes.CreateTunnelResponse{Tunnel: tunnelReturn}, "创建隧道成功", c)	
}

// DeleteTunnel
// @Tags      Tunnel
// @Summary   用户申请删除通道
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body    tunnelReq.DeleteTunnel "隧道ID"
// @Success   200   {object}  response.Response{msg=string}  "隧道删除成功"
// @Router    /user/changePassword [post]
func (t *TunnelApi) DeleteTunnel(c *gin.Context) {
	var req tunnelReq.DeleteTunnel
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return 
	}

	err = utils.Verify(req, utils.DeleteTunnelVerify)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = tunnelService.DeleteTunnel(req.Tunnelid)
	if err != nil {
		global.GVA_LOG.Error("删除隧道失败！", zap.Error(err))
		response.FailWithMessage("删除隧道失败", c)
		return
	}

	response.OkWithMessage("删除隧道成功", c)
}

// FindUserAllTunnels
// @Tags      Tunnel
// @Summary   用户查询其拥有的所有通道
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body    tunnelReq.FindUserAllTunnel "用户ID"
// @Success   200   {object}  response.Response{msg=string}  "隧道查询成功"
// @Router    /user/changePassword [post]

func (t *TunnelApi) FindUserAllTunnels(c *gin.Context) {
	var req tunnelReq.FindUserAllTunnels
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return 
	}

	err = utils.Verify(req, utils.FindUserAllTunnelsVerify)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	tunnels, err := tunnelService.FindUserAllTunnels(req.Userid)
	if err != nil {
		global.GVA_LOG.Error("查询隧道失败！", zap.Error(err))
		response.FailWithMessage("查询隧道失败", c)
		return
	}

	response.OkWithDetailed(tunnelRes.FindUserAllTunnelsResponse{Tunnels: tunnels}, "查询隧道成功", c)
}