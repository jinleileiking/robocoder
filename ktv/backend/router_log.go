package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterLogSetup(engine *gin.Engine) {
    r := engine.Group("/v1")

    r.GET("/logs", RouterLogList)
    r.GET("/logs/:id", RouterLogGet)
    r.POST("/logs", RouterLogCreate)
    r.PUT("/logs/:id", RouterLogUpdate)
    r.DELETE("/logs/:id", RouterLogDelete)
}

func RouterLogList(c *gin.Context) {
	var logs []Log

	logs, err := DBList[Log]()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, logs)
}

func RouterLogGet(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	log, err := DBGet[Log](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, log)
}

func RouterLogCreate(c *gin.Context) {
	var log Log
	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := DBCreate[Log](log); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, log)
}

func RouterLogUpdate(c *gin.Context) {
	var log Log

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	origin, err := DBGet[Log](idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Log not found"})

		return
	}

	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err = DBUpdate[Log](origin, log)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Log not found"})

		return
	}

	c.JSON(http.StatusOK, log)
}

func RouterLogDelete(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Log not found"})

		return
	}

	err = DBDelete[Log](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)

		return
	}

	c.Status(http.StatusNoContent)
}