package main

import (
	"gorm.io/gorm"
)

type NotificationTemplate struct {
	gorm.Model
	Name       string `gorm:"type:varchar(100);not null;comment:通知模板名称" json:"name"`
	Content    string `gorm:"type:text;not null;comment:通知模板内容" json:"content"`
	ReceiverID uint   `gorm:"not null;comment:接收人ID" json:"receiver_id"`
}