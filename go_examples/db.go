package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func DBCreate[T any](entity T) error {
	return db.Create(&entity).Error
}

func DBGet[T any](id int) (T, error) {
	var entity T
	if err := db.First(&entity, id).Error; err != nil {
		return entity, err
	}

	return entity, nil
}

func DBList[T any]() ([]T, error) {
	var entities []T
	if err := db.Find(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}

func DBUpdate[T any](origin, updated T) error {
	return db.Model(&origin).Updates(updated).Error
}

func DBDelete[T any](id int) error {
	var entity T

	return db.Delete(entity, id).Error
}

func Init() {
	var err error
	dsn := "user:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	migrate()
}
