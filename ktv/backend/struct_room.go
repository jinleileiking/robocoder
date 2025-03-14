package main

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null;comment:包房名称" json:"name"`
	Description string `gorm:"type:varchar(255);comment:包房描述" json:"description"`
	Status      string `gorm:"type:varchar(50);comment:包房状态" json:"status"`
	DeviceCount int    `gorm:"type:int;comment:包房内设备数量" json:"device_count"`
}