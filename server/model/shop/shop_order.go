package shop

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status uint
const (
    Unpaid Status = iota
    Paid
)

type ShopOrder struct {
	global.GVA_MODEL
	OrderID     uuid.UUID	`json:"orderId" gorm:"primarykey;autoIncrement;comment:订单ID"`
	UserID      uint    	`json:"userId" gorm:"foreignkey;comment:订单所属用户ID"`
	ServiceType uint		`json:"serviceType" gorm:"comment:服务类型"`
	Price       float64 	`json:"price" gorm:"comment:订单价格"`
	Status      uint    	`json:"status" gorm:"comment:订单状态"`
	CreateTime  int64   	`json:"createTime" gorm:"comment:订单创建时间"`
	ExpireTime  int64   	`json:"expireTime" gorm:"comment:订单过期时间"`
}

func(order *ShopOrder) TableName() string {
	return "shop_orders"
}

// GetOrderByID 通过ID查询订单
func GetOrderByID(db *gorm.DB, id uint) (*ShopOrder, error) {
    var order ShopOrder
    if err := db.Where("id = ?", id).First(&order).Error; err != nil {
        return nil, err
    }
    return &order, nil
}

// GetOrdersByUserID 查询用户的所有订单
func GetOrdersByUserID(db *gorm.DB, userID uint) ([]ShopOrder, error) {
    var orders []ShopOrder
    if err := db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
        return nil, err
    }
    return orders, nil
}

// GetOrdersByStatus 根据状态查询订单
func GetOrdersByStatus(db *gorm.DB, status uint) ([]ShopOrder, error) {
    var orders []ShopOrder
    if err := db.Where("status = ?", status).Find(&orders).Error; err != nil {
        return nil, err
    }
    return orders, nil
}

// GetOrdersByTimeRange 根据时间范围查询订单
func GetOrdersByTimeRange(db *gorm.DB, startTime, endTime int64) ([]ShopOrder, error) {
    var orders []ShopOrder
    if err := db.Where("create_time BETWEEN ? AND ?", startTime, endTime).Find(&orders).Error; err != nil {
        return nil, err
    }
    return orders, nil
}

// GetExpiredOrders 查询已过期订单
func GetExpiredOrders(db *gorm.DB) ([]ShopOrder, error) {
    var orders []ShopOrder
    currentTime := time.Now().Unix()
    if err := db.Where("expire_time < ?", currentTime).Find(&orders).Error; err != nil {
        return nil, err
    }
    return orders, nil
}

// GetOrdersWithPagination 分页查询订单
func GetOrdersWithPagination(db *gorm.DB, page, pageSize int) ([]ShopOrder, int64, error) {
    var orders []ShopOrder
    var total int64
    
    offset := (page - 1) * pageSize
    if err := db.Model(&ShopOrder{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }
    
    if err := db.Limit(pageSize).Offset(offset).Find(&orders).Error; err != nil {
        return nil, 0, err
    }
    
    return orders, total, nil
}