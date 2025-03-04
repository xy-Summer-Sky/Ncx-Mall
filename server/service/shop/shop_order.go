// 订单db增删查改
package shop

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	"gorm.io/gorm"
)

type ShopOrderService struct {}

//@author: [smallcjy](https://github.com/smallcjy)
//@function: CreateShopOrder
//@description: 创建用户订单
//@param: s model.ShopOrder
//@return: err error, shopOrderInter model.ShopOrder
func (shopOrderService *ShopOrderService) CreateShopOrder(s shop.ShopOrder) (shopOrderInter shop.ShopOrder, err error) {
	var shopOrderModel shop.ShopOrder
	if !errors.Is(global.GVA_DB.Where("id = ?", s.ID).First(&shopOrderModel).Error, gorm.ErrRecordNotFound) {
		return shopOrderInter, errors.New("存在相同订单号")
	}
	// 否则 创建订单
	err = global.GVA_DB.Create(&s).Error
	if err != nil {
		return shopOrderInter, err
	}

	return s, err
}

//@author: [smallcjy](https://github.com/smallcjy)
//@function: DeleteShopOrder
//@description: 取消用户订单
//@param: id int
//@return: err error
func (shopOrderService *ShopOrderService) DeleteShopOrder(id int) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&shop.ShopOrder{}).Error; err != nil {
			return err
		}
		return nil
	})
}

//@author: [smallcjy](https://github.com/smallcjy)
//@function: PayShopOrder
//@description: 支付用户订单
//@param: id int
//@return: err error
func (shopOrderService *ShopOrderService) PayShopOrder(id int) (err error) {
    var order shop.ShopOrder
    // 查询订单信息
    if err := global.GVA_DB.Where("id = ?", id).First(&order).Error; err != nil {
        return err
    }

    // 判断订单是否已经支付
    if order.Status == 1 {
        return errors.New("订单已支付")
    }

    // 调用微信支付接口下单，生成预支付订单
    var wxPayment WXPayService
    wxResp, err := wxPayment.WechatPayOrder(order)
    if err != nil {
        return err
    }

    // 更新订单状态为预支付，并保存微信返回的预支付ID，方便后续生成支付二维码等操作
    if err := global.GVA_DB.Model(&order).Updates(map[string]interface{}{
        "Status":   "prepaid",
        "PrepayID": wxResp.PrepayID,
    }).Error; err != nil {
        return err
    }

    // 可以选择将wxResp返回给上层，由前端生成支付二维码等处理
    return nil
}

//@author: [smallcjy](https://github.com/smallcjy)
//@function: GetUserOrders
//@description: 查询用户所有订单
//@param: id int
//@return: err error
func (shopOrderService *ShopOrderService) GetUserOrders(id int) (orders []shop.ShopOrder, err error) {
	err = global.GVA_DB.Where("user_id = ?", id).Find(&orders).Error
	return orders, err
}