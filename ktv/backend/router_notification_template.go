package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterNotificationTemplateSetup(engine *gin.Engine) {
	r := engine.Group("/v1")

	r.GET("/notification_templates", RouterNotificationTemplateList)
	r.GET("/notification_templates/:id", RouterNotificationTemplateGet)
	r.POST("/notification_templates", RouterNotificationTemplateCreate)
	r.PUT("/notification_templates/:id", RouterNotificationTemplateUpdate)
	r.DELETE("/notification_templates/:id", RouterNotificationTemplateDelete)
}

func RouterNotificationTemplateList(c *gin.Context) {
	var notificationTemplates []NotificationTemplate

	notificationTemplates, err := DBList[NotificationTemplate]()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, notificationTemplates)
}

func RouterNotificationTemplateGet(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	notificationTemplate, err := DBGet[NotificationTemplate](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, notificationTemplate)
}

func RouterNotificationTemplateCreate(c *gin.Context) {
	var notificationTemplate NotificationTemplate
	if err := c.ShouldBindJSON(&notificationTemplate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := DBCreate[NotificationTemplate](notificationTemplate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, notificationTemplate)
}

func RouterNotificationTemplateUpdate(c *gin.Context) {
	var notificationTemplate NotificationTemplate

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	origin, err := DBGet[NotificationTemplate](idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "NotificationTemplate not found"})

		return
	}

	if err := c.ShouldBindJSON(&notificationTemplate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err = DBUpdate[NotificationTemplate](origin, notificationTemplate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "NotificationTemplate not found"})

		return
	}

	c.JSON(http.StatusOK, notificationTemplate)
}

func RouterNotificationTemplateDelete(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "NotificationTemplate not found"})

		return
	}

	err = DBDelete[NotificationTemplate](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)

		return
	}

	c.Status(http.StatusNoContent)
}