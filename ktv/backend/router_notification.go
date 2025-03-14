package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterNotificationSetup(engine *gin.Engine) {
    r := engine.Group("/v1")

    r.GET("/notifications", RouterNotificationList)
    r.GET("/notifications/:id", RouterNotificationGet)
    r.POST("/notifications", RouterNotificationCreate)
    r.PUT("/notifications/:id", RouterNotificationUpdate)
    r.DELETE("/notifications/:id", RouterNotificationDelete)
}

func RouterNotificationList(c *gin.Context) {
	var notifications []Notification

	notifications, err := DBList[Notification]()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, notifications)
}

func RouterNotificationGet(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	notification, err := DBGet[Notification](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, notification)
}

func RouterNotificationCreate(c *gin.Context) {
	var notification Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := DBCreate[Notification](notification); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, notification)
}

func RouterNotificationUpdate(c *gin.Context) {
	var notification Notification

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	origin, err := DBGet[Notification](idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})

		return
	}

	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err = DBUpdate[Notification](origin, notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Notification not found"})

		return
	}

	c.JSON(http.StatusOK, notification)
}

func RouterNotificationDelete(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Notification not found"})

		return
	}

	err = DBDelete[Notification](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)

		return
	}

	c.Status(http.StatusNoContent)
}