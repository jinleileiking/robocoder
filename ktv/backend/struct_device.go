package main

import (
	"gorm.io/gorm"
	"time"
)

type Device struct {
	gorm.Model
	Name             string    `gorm:"type:varchar(100);not null;comment:设备名称" json:"name"`
	Type             string    `gorm:"type:varchar(50);not null;comment:设备种类" json:"type"`
	HealthStatus     string    `gorm:"type:varchar(50);not null;comment:设备健康状态" json:"health_status"`
	ModelType        string    `gorm:"type:varchar(100);not null;comment:设备型号" json:"model_type"`
	Uptime           int64     `gorm:"not null;comment:开机时长（秒）" json:"uptime"`
	CPUUsage         float64   `gorm:"not null;comment:CPU占用百分比" json:"cpu_usage"`
	MemoryUsage      float64   `gorm:"not null;comment:内存占用百分比" json:"memory_usage"`
	RoomID           uint      `gorm:"not null;comment:所属包房ID" json:"room_id"`
	LastCheckedAt    time.Time `gorm:"not null;comment:最后检查时间" json:"last_checked_at"`
}