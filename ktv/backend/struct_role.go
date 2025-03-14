package main

import (
	"gorm.io/gorm"
)

// Role represents the user role in the system with different permissions.
type Role struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null;comment:角色名称" json:"name"`
	Description string `gorm:"type:varchar(255);comment:角色描述" json:"description"`
	Permissions string `gorm:"type:text;comment:角色权限列表（JSON格式）" json:"permissions"`
}