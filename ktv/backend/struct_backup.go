package main

import (
	"gorm.io/gorm"
	"time"
)

// Backup represents the data backup and restore functionality in the system.
type Backup struct {
	gorm.Model
	BackupID   string    `gorm:"type:varchar(100);not null;unique;comment:备份ID" json:"backup_id"`
	BackupName string    `gorm:"type:varchar(255);not null;comment:备份名称" json:"backup_name"`
	Status     string    `gorm:"type:varchar(50);not null;comment:备份状态" json:"status"`
	CreatedBy  string    `gorm:"type:varchar(100);not null;comment:创建者" json:"created_by"`
	CreatedAt  time.Time `gorm:"type:datetime;not null;comment:创建时间" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null;comment:更新时间" json:"updated_at"`
}