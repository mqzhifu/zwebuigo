package model

import "gorm.io/gorm"

type DIY_MODEL struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt int            // 创建时间
	UpdatedAt int            // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
