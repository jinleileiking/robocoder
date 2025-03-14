package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterNotificationRecordSetup(engine *gin.Engine) {
    r := engine.Group("/v1")

    r.GET("/notification_records", RouterNotificationRecordList)
    r.GET("/notification_records/:id", RouterNotificationRecordGet)
    r.POST("/notification_records", RouterNotificationRecordCreate)
    r.PUT("/notification_records/:id", RouterNotificationRecordUpdate)
    r.DELETE("/notification_records/:id", RouterNotificationRecordDelete)
}

func RouterNotificationRecordList(c *gin.Context) {
	var notificationRecords []NotificationRecord

	notificationRecords, err := DBList[NotificationRecord]()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, notificationRecords)
}

func RouterNotificationRecordGet(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	notificationRecord, err := DBGet[NotificationRecord](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, notificationRecord)
}

func RouterNotificationRecordCreate(c *gin.Context) {
	var notificationRecord NotificationRecord
	if err := c.ShouldBindJSON(&notificationRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := DBCreate[NotificationRecord](notificationRecord); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, notificationRecord)
}

func RouterNotificationRecordUpdate(c *gin.Context) {
	var notificationRecord NotificationRecord

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	origin, err := DBGet[NotificationRecord](idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "NotificationRecord not found"})

		return
	}

	if err := c.ShouldBindJSON(&notificationRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err = DBUpdate[NotificationRecord](origin, notificationRecord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "NotificationRecord not found"})

		return
	}

	c.JSON(http.StatusOK, notificationRecord)
}

func RouterNotificationRecordDelete(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "NotificationRecord not found"})

		return
	}

	err = DBDelete[NotificationRecord](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)

		return
	}

	c.Status(http.StatusNoContent)
}