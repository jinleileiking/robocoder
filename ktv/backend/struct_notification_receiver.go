package main

import (
	"gorm.io/gorm"
)

// NotificationReceiver represents the configuration for receiving notifications via Feishu.
type NotificationReceiver struct {
	gorm.Model
	ReceiverName string `gorm:"type:varchar(100);not null;comment:接收人名称" json:"receiver_name"`
	FeishuID     string `gorm:"type:varchar(100);not null;comment:飞书ID" json:"feishu_id"`
	RoomID       uint   `gorm:"not null;comment:包房ID" json:"room_id"`
}
