package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime"`// 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime"`// 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
