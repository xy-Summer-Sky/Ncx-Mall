package utils

type VipType int 
type Traffictype int

const (
	Free VipType = iota
	Vip  
	SuperVip
)
const (
	Tcp Traffictype = iota
	Udp
)

// String 会员类型字符串表示
func (v VipType) String() string {
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

// 根据VIP类型得到服务配置信息
type VipServiceConfig struct {
	ID          uint        `json:"id" gorm:"primarykey"`
    Type        VipType     `json:"type" gorm:"uniqueIndex;comment:会员类型"`
    Price       float64     `json:"price" gorm:"comment:会员价格"`
    Speed       int32       `json:"speed" gorm:"comment:允许最大隧道带宽"`
    TunnelNum   int32       `json:"tunnelNum" gorm:"comment:允许最大隧道数量"`
    Traffic     Traffictype `json:"traffic" gorm:"comment:隧道类型"`
}

func GetVipServiceConfig(vip VipType) VipServiceConfig {
	switch vip {
	case Free:
		return VipServiceConfig{
			ID: 1,
			Type: Free,
			Price: 0,
			Speed: 10,
			TunnelNum: 1,
			Traffic: Tcp,
		}
	case Vip:
		return VipServiceConfig{
			ID: 2,
			Type: Vip,
			Price: 10,
			Speed: 100,
			TunnelNum: 5,
			Traffic: Tcp,
		}
	case SuperVip:
		return VipServiceConfig{
			ID: 3,
			Type: SuperVip,
			Price: 20,
			Speed: 1000,
			TunnelNum: 10,
			Traffic: Udp,
		}
	default:
		return VipServiceConfig{}
	}
}