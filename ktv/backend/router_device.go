package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterDeviceSetup(engine *gin.Engine) {
	r := engine.Group("/v1")

	r.GET("/devices", RouterDeviceList)
	r.GET("/devices/:id", RouterDeviceGet)
	r.POST("/devices", RouterDeviceCreate)
	r.PUT("/devices/:id", RouterDeviceUpdate)
	r.DELETE("/devices/:id", RouterDeviceDelete)
}

func RouterDeviceList(c *gin.Context) {
	var devices []Device

	devices, err := DBList[Device]()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, devices)
}

func RouterDeviceGet(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	device, err := DBGet[Device](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, device)
}

func RouterDeviceCreate(c *gin.Context) {
	var device Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := DBCreate[Device](device); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, device)
}

func RouterDeviceUpdate(c *gin.Context) {
	var device Device

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	origin, err := DBGet[Device](idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})

		return
	}

	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err = DBUpdate[Device](origin, device)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Device not found"})

		return
	}

	c.JSON(http.StatusOK, device)
}

func RouterDeviceDelete(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Device not found"})

		return
	}

	err = DBDelete[Device](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)

		return
	}

	c.Status(http.StatusNoContent)
}