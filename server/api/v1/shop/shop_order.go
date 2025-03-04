package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	shopRes "github.com/flipped-aurora/gin-vue-admin/server/model/shop/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ShopOrderApi struct {}

// CreateShopOrder
// @Tags     Shop
// @Summary  用户申请创建订单
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data body     shopReq.CreateOrder true "订单信息"
// @Success  200  {object} response.Response{data=shopRes.CreateShopOrderResponse} "订单创建成功"
// @Router   /shop/CreateShopOrder [post]
func (s *ShopOrderApi) CreateShopOrder(c *gin.Context) {
	var req shopReq.CreateOrder
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(req, utils.CreateShopOrderVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 创建订单
	order := &shop.ShopOrder {
		OrderID: uuid.New(),
		ServiceType: req.ServiceType,
		Price: req.Price,
		Status: req.Status,
		CreateTime: req.CreateTime,
		ExpireTime: req.ExpireTime,
	}

	order.UserID = utils.GetUserID(c)
	tunnelReturn, err := shopOrderService.ShopOrderService.CreateShopOrder(*order)

	if err != nil {
		global.GVA_LOG.Error("创建订单失败！", zap.Error(err))
		response.FailWithDetailed(shopRes.CreateShopOrderResponse{ShopOrder: tunnelReturn}, "创建订单失败", c)
		return
	}

	response.OkWithDetailed(shopRes.CreateShopOrderResponse{ShopOrder: tunnelReturn}, "创建订单成功", c)
}

// DeleteShopOrder
// @Tags      Shop
// @Summary   用户申请取消订单
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data body shopReq.DeleteOrder true "订单取消请求体"
// @Success   200 {object} response.Response{msg=string} "订单删除成功"
// @Router    /shop/DeleteShopOrder [post]
func (s *ShopOrderApi) DeleteShopOrder(c *gin.Context) {
	var req shopReq.DeleteOrder
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(req, utils.DeleteShopOrderVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = shopOrderService.ShopOrderService.DeleteShopOrder(int(req.OrderID))
	if err != nil {
		global.GVA_LOG.Error("删除订单失败！", zap.Error(err))
		response.FailWithMessage("删除订单失败", c)
		return
	}

	response.OkWithMessage("删除订单成功", c)
}

// TODO: 一次性返回所有订单性能开销大，后续可以考虑分页查询实现流式传输

// GetUserOrders
// @Tags      Shop
// @Summary   用户查询其拥有的所有订单
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data body shopReq.GetUserOrders true "查询用户订单请求体"
// @Success   200 {object} response.Response{data=shopRes.GetUserOrdersResponse} "订单查询成功"
// @Router    /shop/GetUserOrders [post]
func (s *ShopOrderApi) GetUserOrders(c *gin.Context) {
	var req shopReq.GetUserOrders
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(req, utils.GetUserOrdersVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	orders, err := shopOrderService.ShopOrderService.GetUserOrders(int(req.UserID))
	if err != nil {
		global.GVA_LOG.Error("查询订单失败！", zap.Error(err))
		response.FailWithMessage("查询订单失败", c)
		return
	}

	response.OkWithDetailed(shopRes.GetUserOrdersResponse{ShopOrders: orders}, "查询订单成功", c)

}


// 微信回调函数的url
func (w * ShopOrderApi)NotifyWeChatPay(c * gin.Context) {
   
shopOrderService.WXPayService.NotifyWeChatPay(c)
   
}