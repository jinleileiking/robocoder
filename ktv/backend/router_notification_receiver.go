package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterNotificationReceiverSetup(engine *gin.Engine) {
    r := engine.Group("/v1")

    r.GET("/notificationreceivers", RouterNotificationReceiverList)
    r.GET("/notificationreceivers/:id", RouterNotificationReceiverGet)
    r.POST("/notificationreceivers", RouterNotificationReceiverCreate)
    r.PUT("/notificationreceivers/:id", RouterNotificationReceiverUpdate)
    r.DELETE("/notificationreceivers/:id", RouterNotificationReceiverDelete)
}

func RouterNotificationReceiverList(c *gin.Context) {
	var notificationReceivers []NotificationReceiver

	notificationReceivers, err := DBList[NotificationReceiver]()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, notificationReceivers)
}

func RouterNotificationReceiverGet(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	notificationReceiver, err := DBGet[NotificationReceiver](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, notificationReceiver)
}

func RouterNotificationReceiverCreate(c *gin.Context) {
	var notificationReceiver NotificationReceiver
	if err := c.ShouldBindJSON(&notificationReceiver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := DBCreate[NotificationReceiver](notificationReceiver); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, notificationReceiver)
}

func RouterNotificationReceiverUpdate(c *gin.Context) {
	var notificationReceiver NotificationReceiver

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	origin, err := DBGet[NotificationReceiver](idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "NotificationReceiver not found"})

		return
	}

	if err := c.ShouldBindJSON(&notificationReceiver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err = DBUpdate[NotificationReceiver](origin, notificationReceiver)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "NotificationReceiver not found"})

		return
	}

	c.JSON(http.StatusOK, notificationReceiver)
}

func RouterNotificationReceiverDelete(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "NotificationReceiver not found"})

		return
	}

	err = DBDelete[NotificationReceiver](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)

		return
	}

	c.Status(http.StatusNoContent)
}