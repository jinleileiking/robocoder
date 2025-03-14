package main

import "github.com/gin-gonic/gin"

func main() {
	Init()
	router := gin.Default()

	RouterRoomSetup(router)
	RouterDeviceSetup(router)
	RouterExceptionSetup(router)
	RouterExceptionRecordSetup(router)
	RouterNotificationTemplateSetup(router)
	RouterNotificationRecordSetup(router)
	
	router.Run(":8080")
}

func migrate() {
	db.AutoMigrate(&Room{}, &Device{}, &Exception{}, &ExceptionRecord{}, &NotificationTemplate{}, &NotificationRecord{})
}