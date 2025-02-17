package tunnel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tunnel struct {
    global.GVA_MODEL
	Tunnelid  	uuid.UUID   `json:"tunnelid" gorm:"primarykey;autoIncrement;comment:通道ID"`
	Tunnelname 	string    	`json:"tunnelname" gorm:"comment:通道名称"`
	Tunneltype 	string    	`json:"tunneltype" gorm:"comment:通道类型"`
	Tunnelpoint string   	`json:"tunnelpoint" gorm:"comment:通道点"`
	Localip 	string   	`json:"localip" gorm:"comment:本地IP"`
	Localport 	uint16   	`json:"localport" gorm:"comment:本地端口"`
	Remtoeip 	string   	`json:"remoteip" gorm:"comment:远程IP"`
	Remoteport 	uint16   	`json:"remoteport" gorm:"comment:远程端口"`
	Token 		string   	`json:"token" gorm:"comment:通道令牌"`
    Userid      uint        `json:"userid" gorm:"foreignkey;comment:隧道所属用户ID"`
}


// TableName 指定表名
func (t *Tunnel) TableName() string {
    return "tunnels"
}

// GetTunnelById 根据ID查询通道
func GetTunnelById(db *gorm.DB, tunnelId string) (*Tunnel, error) {
    var tunnel Tunnel
    err := db.Where("tunnelid = ?", tunnelId).First(&tunnel).Error
    if err != nil {
        return nil, err
    }
    return &tunnel, nil
}

// ListTunnels 获取通道列表
func ListTunnels(db *gorm.DB) ([]Tunnel, error) {
    var tunnels []Tunnel
    err := db.Find(&tunnels).Error
    if err != nil {
        return nil, err
    }
    return tunnels, nil
}

// UpdateTunnel 更新通道信息
func UpdateTunnel(db *gorm.DB, tunnel *Tunnel) error {
    return db.Model(tunnel).Updates(map[string]interface{}{
        "tunnelname":  tunnel.Tunnelname,
        "tunneltype":  tunnel.Tunneltype,
        "tunnelpoint": tunnel.Tunnelpoint,
        "localip":     tunnel.Localip,
        "localport":   tunnel.Localport,
        "remoteip":    tunnel.Remtoeip,
        "remoteport":  tunnel.Remoteport,
        "token":       tunnel.Token,
    }).Error
}

// DeleteTunnel 删除通道
func DeleteTunnel(db *gorm.DB, tunnelId string) error {
    return db.Where("tunnelid = ?", tunnelId).Delete(&Tunnel{}).Error
}

// GetTunnelsByType 根据通道类型查询
func GetTunnelsByType(db *gorm.DB, tunnelType string) ([]Tunnel, error) {
    var tunnels []Tunnel
    err := db.Where("tunneltype = ?", tunnelType).Find(&tunnels).Error
    if err != nil {
        return nil, err
    }
    return tunnels, nil
}

// GetTunnelsByPoint 根据通道点查询
func GetTunnelsByPoint(db *gorm.DB, point string) ([]Tunnel, error) {
    var tunnels []Tunnel
    err := db.Where("tunnelpoint = ?", point).Find(&tunnels).Error
    if err != nil {
        return nil, err
    }
    return tunnels, nil
}