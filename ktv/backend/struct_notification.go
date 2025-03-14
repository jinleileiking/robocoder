package main

import (
	"gorm.io/gorm"
	"time"
)

type Notification struct {
	gorm.Model
	RoomName      string    `gorm:"type:varchar(100);not null;comment:包房名称" json:"room_name"`
	DeviceName    string    `gorm:"type:varchar(100);not null;comment:设备名称" json:"device_name"`
	NotificationType string `gorm:"type:varchar(50);not null;comment:通知类型" json:"notification_type"`
	NotificationTime time.Time `gorm:"type:datetime;not null;comment:通知时间" json:"notification_time"`
	Link          string    `gorm:"type:varchar(255);not null;comment:跳转链接" json:"link"`
	Recipient     string    `gorm:"type:varchar(100);not null;comment:接收人" json:"recipient"`
	Message       string    `gorm:"type:text;not null;comment:通知内容" json:"message"`
}