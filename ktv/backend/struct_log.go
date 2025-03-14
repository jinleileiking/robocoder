package main

import (
	"gorm.io/gorm"
	"time"
)

type Log struct {
	gorm.Model
	UserID      uint      `gorm:"not null;comment:用户ID" json:"user_id"`
	Action      string    `gorm:"type:varchar(255);not null;comment:用户操作" json:"action"`
	Description string    `gorm:"type:varchar(255);comment:操作描述" json:"description"`
	Timestamp   time.Time `gorm:"not null;comment:操作时间" json:"timestamp"`
	ModelType   string    `gorm:"type:varchar(100);comment:操作涉及的模型类型" json:"model_type"`
	ModelID     uint      `gorm:"comment:操作涉及的模型ID" json:"model_id"`
}