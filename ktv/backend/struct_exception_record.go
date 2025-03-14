package main

import (
	"time"
	"gorm.io/gorm"
)

// ExceptionRecord represents a record of an exception that occurred in a device.
type ExceptionRecord struct {
	gorm.Model
	RoomName       string    `gorm:"type:varchar(100);not null;comment:包房名称" json:"room_name"`
	DeviceName     string    `gorm:"type:varchar(100);not null;comment:设备名称" json:"device_name"`
	ExceptionTime  time.Time `gorm:"type:datetime;not null;comment:异常时间" json:"exception_time"`
	Description    string    `gorm:"type:text;comment:异常描述" json:"description"`
	Handler        string    `gorm:"type:varchar(100);comment:处理人" json:"handler"`
	HandledTime    time.Time `gorm:"type:datetime;comment:处理时间" json:"handled_time"`
	HandleResult   string    `gorm:"type:text;comment:处理结果" json:"handle_result"`
}