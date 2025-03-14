package main

import (
	"gorm.io/gorm"
	"time"
)

type NotificationRecord struct {
	gorm.Model
	Recipient       string    `gorm:"type:varchar(100);not null;comment:接收人" json:"recipient"`                        // 接收人
	ContentTemplate string    `gorm:"type:text;not null;comment:通知内容模板" json:"content_template"`                     // 通知内容模板
	RoomName        string    `gorm:"type:varchar(100);not null;comment:包房名称" json:"room_name"`                        // 包房名称
	DeviceName      string    `gorm:"type:varchar(100);not null;comment:设备名称" json:"device_name"`                      // 设备名称
	Link            string    `gorm:"type:varchar(255);comment:跳转链接" json:"link"`                                     // 跳转链接
	SentAt          time.Time `gorm:"type:datetime;not null;comment:发送时间" json:"sent_at"`                              // 发送时间
	Status          string    `gorm:"type:varchar(50);not null;comment:通知状态" json:"status"`                            // 通知状态
}