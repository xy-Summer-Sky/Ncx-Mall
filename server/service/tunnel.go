package service

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tunnel" // correct import path
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TunnelService struct{}

//@author: [smallcjy](https://github.com/smallcjy)
//@function: CreateTunnel
//@description: 创建隧道
//@param: tunnel model.Tunnel
//@return: err error, tunnelInter model.Tunnel
func (tunnelService *TunnelService) CreateTunnel(t tunnel.Tunnel) (tunnelInter tunnel.Tunnel, err error) {
	var tunnelModel tunnel.Tunnel
	uid := t.Userid
	if !errors.Is(global.GVA_DB.Where("remoteip = ?", t.Remtoeip).First(&tunnelModel).Error, gorm.ErrRecordNotFound) { // 判断隧道的远程端口和地址是否已存在
		if !errors.Is(global.GVA_DB.Where("remoteport = ?", t.Remoteport).First(&tunnelModel).Error, gorm.ErrRecordNotFound) {
			return tunnelInter, errors.New("远程隧道已被占用")
		}
	}
	// 否则 创建隧道
	t.Tunnelid = uuid.New()
	t.Userid = uid
	err = global.GVA_DB.Create(&t).Error
	if err != nil {
		return tunnelInter, err
	}

	return t, err
}

//@author: [smallcjy](https://github.com/smallcjy)
//@function: DeleteTunnel
//@description: 删除隧道
//@param: id
//@return: err error

func (tunnelService *TunnelService) DeleteTunnel(id uint) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("tunnelid = ?", id).Delete(&tunnel.Tunnel{}).Error; err != nil {
			return err
		}
		return nil
	})
}

//@author: [smallcjy](https://github.com/smallcjy)
//@function: FindUserAllTunnels
//@description: 查找某个用户的所有隧道
//@param: id
//@return: err error 	Tunnels []model.Tunnel

func (tunnelService *TunnelService) FindUserAllTunnels(id uint) (Tunnels []tunnel.Tunnel, err error) {
    err = global.GVA_DB.Where("userid = ?", id).Find(&Tunnels).Error
    return Tunnels, err
}