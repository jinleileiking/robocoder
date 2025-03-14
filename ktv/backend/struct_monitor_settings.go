package main

import (
	"gorm.io/gorm"
	"time"
)

type MonitorSettings struct {
	gorm.Model
	Frequency       int       `gorm:"type:int;not null;comment:监控频率（秒）" json:"frequency"`
	LastCheckedAt   time.Time `gorm:"type:datetime;comment:最后一次检查时间" json:"last_checked_at"`
	NotificationIDs string    `gorm:"type:varchar(255);comment:接收通知的人员ID列表，逗号分隔" json:"notification_ids"`
	Enabled         bool      `gorm:"type:bool;not null;default:true;comment:是否启用监控" json:"enabled"`
}