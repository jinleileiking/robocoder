package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterBackupSetup(engine *gin.Engine) {
	r := engine.Group("/v1")

	r.GET("/backups", RouterBackupList)
	r.GET("/backups/:id", RouterBackupGet)
	r.POST("/backups", RouterBackupCreate)
	r.PUT("/backups/:id", RouterBackupUpdate)
	r.DELETE("/backups/:id", RouterBackupDelete)
}

func RouterBackupList(c *gin.Context) {
	var backups []Backup

	backups, err := DBList[Backup]()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, backups)
}

func RouterBackupGet(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	backup, err := DBGet[Backup](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, backup)
}

func RouterBackupCreate(c *gin.Context) {
	var backup Backup
	if err := c.ShouldBindJSON(&backup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DBCreate[Backup](backup); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, backup)
}

func RouterBackupUpdate(c *gin.Context) {
	var backup Backup

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	origin, err := DBGet[Backup](idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Backup not found"})
		return
	}

	if err := c.ShouldBindJSON(&backup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = DBUpdate[Backup](origin, backup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Backup not found"})
		return
	}

	c.JSON(http.StatusOK, backup)
}

func RouterBackupDelete(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Backup not found"})
		return
	}

	err = DBDelete[Backup](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Status(http.StatusNoContent)
}