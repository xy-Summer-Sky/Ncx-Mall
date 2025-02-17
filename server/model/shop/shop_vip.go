package shop

import (
	"time"

	"gorm.io/gorm"
)

type Viptype 		int
type Traffictype 	int

const (
	Free Viptype = iota
	Vip  
	SuperVip
)

const (
	Tcp Traffictype = iota
	Udp
)

type ShopVip struct {
    ID          uint        `json:"id" gorm:"primarykey"`
    Type        Viptype     `json:"type" gorm:"uniqueIndex;comment:会员类型"`
    Price       float64     `json:"price" gorm:"comment:会员价格"`
    Speed       int32       `json:"speed" gorm:"comment:允许最大隧道带宽"`
    TunnelNum   int32       `json:"tunnelNum" gorm:"comment:允许最大隧道数量"`
    Traffic     Traffictype `json:"traffic" gorm:"comment:隧道类型"`
    CreatedAt   time.Time   `json:"createdAt"`
    UpdatedAt   time.Time   `json:"updatedAt"`
}


// TableName 指定表名
func (s *ShopVip) TableName() string {
    return "shop_vips"
}

// CreateShopVip 创建会员套餐
func CreateShopVip(db *gorm.DB, vip *ShopVip) error {
    return db.Create(vip).Error
}

// GetShopVipByType 根据会员类型查询套餐
func GetShopVipByType(db *gorm.DB, vipType Viptype) (*ShopVip, error) {
    var vip ShopVip
    err := db.Where("type = ?", vipType).First(&vip).Error
    if err != nil {
        return nil, err
    }
    return &vip, nil
}

// ListShopVips 获取所有会员套餐列表
func ListShopVips(db *gorm.DB) ([]ShopVip, error) {
    var vips []ShopVip
    err := db.Find(&vips).Error
    if err != nil {
        return nil, err
    }
    return vips, nil
}

// UpdateShopVip 更新会员套餐信息
func UpdateShopVip(db *gorm.DB, vip *ShopVip) error {
    return db.Model(vip).Updates(map[string]interface{}{
        "price":      vip.Price,
        "speed":      vip.Speed,
        "tunnel_num": vip.TunnelNum,
        "traffic":    vip.Traffic,
    }).Error
}

// DeleteShopVip 删除会员套餐
func DeleteShopVip(db *gorm.DB, vipType Viptype) error {
    return db.Where("type = ?", vipType).Delete(&ShopVip{}).Error
}

// String 会员类型字符串表示
func (v Viptype) String() string {
    switch v {
    case Free:
        return "免费会员"
    case Vip:
        return "普通会员"
    case SuperVip:
        return "超级会员"
    default:
        return "未知类型"
    }
}

// String 流量类型字符串表示
func (t Traffictype) String() string {
    switch t {
    case Tcp:
        return "TCP"
    case Udp:
        return "UDP"
    default:
        return "未知类型"
    }
}