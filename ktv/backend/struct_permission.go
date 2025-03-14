package main

import (
	"gorm.io/gorm"
)

// Permission represents the permissions assigned to different roles in the system
type Permission struct {
	gorm.Model
	RoleName      string `gorm:"type:varchar(100);not null;comment:角色名称" json:"role_name"`
	Resource      string `gorm:"type:varchar(100);not null;comment:资源名称" json:"resource"`
	Action        string `gorm:"type:varchar(50);not null;comment:操作类型" json:"action"`
	Description   string `gorm:"type:varchar(255);comment:权限描述" json:"description"`
}