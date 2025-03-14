package main

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);not null;comment:门店名称"`
	Address string `gorm:"type:varchar(255);comment:门店地址"`
	Phone   string `gorm:"type:varchar(20);comment:联系电话"`
}
