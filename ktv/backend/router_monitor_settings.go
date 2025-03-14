package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterMonitorSettingsSetup(engine *gin.Engine) {
    r := engine.Group("/v1")

    r.GET("/monitor_settings", RouterMonitorSettingsList)
    r.GET("/monitor_settings/:id", RouterMonitorSettingsGet)
    r.POST("/monitor_settings", RouterMonitorSettingsCreate)
    r.PUT("/monitor_settings/:id", RouterMonitorSettingsUpdate)
    r.DELETE("/monitor_settings/:id", RouterMonitorSettingsDelete)
}

func RouterMonitorSettingsList(c *gin.Context) {
	var monitorSettings []MonitorSettings

	monitorSettings, err := DBList[MonitorSettings]()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, monitorSettings)
}

func RouterMonitorSettingsGet(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	monitorSetting, err := DBGet[MonitorSettings](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, monitorSetting)
}

func RouterMonitorSettingsCreate(c *gin.Context) {
	var monitorSetting MonitorSettings
	if err := c.ShouldBindJSON(&monitorSetting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := DBCreate[MonitorSettings](monitorSetting); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, monitorSetting)
}

func RouterMonitorSettingsUpdate(c *gin.Context) {
	var monitorSetting MonitorSettings

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	origin, err := DBGet[MonitorSettings](idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "MonitorSettings not found"})

		return
	}

	if err := c.ShouldBindJSON(&monitorSetting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err = DBUpdate[MonitorSettings](origin, monitorSetting)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "MonitorSettings not found"})

		return
	}

	c.JSON(http.StatusOK, monitorSetting)
}

func RouterMonitorSettingsDelete(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "MonitorSettings not found"})

		return
	}

	err = DBDelete[MonitorSettings](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)

		return
	}

	c.Status(http.StatusNoContent)
}