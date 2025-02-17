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