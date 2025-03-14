package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterExceptionRecordSetup(engine *gin.Engine) {
    r := engine.Group("/v1")

    r.GET("/exception-records", RouterExceptionRecordList)
    r.GET("/exception-records/:id", RouterExceptionRecordGet)
    r.POST("/exception-records", RouterExceptionRecordCreate)
    r.PUT("/exception-records/:id", RouterExceptionRecordUpdate)
    r.DELETE("/exception-records/:id", RouterExceptionRecordDelete)
}

func RouterExceptionRecordList(c *gin.Context) {
	var exceptionRecords []ExceptionRecord

	exceptionRecords, err := DBList[ExceptionRecord]()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, exceptionRecords)
}

func RouterExceptionRecordGet(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	exceptionRecord, err := DBGet[ExceptionRecord](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, exceptionRecord)
}

func RouterExceptionRecordCreate(c *gin.Context) {
	var exceptionRecord ExceptionRecord
	if err := c.ShouldBindJSON(&exceptionRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := DBCreate[ExceptionRecord](exceptionRecord); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, exceptionRecord)
}

func RouterExceptionRecordUpdate(c *gin.Context) {
	var exceptionRecord ExceptionRecord

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	origin, err := DBGet[ExceptionRecord](idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ExceptionRecord not found"})

		return
	}

	if err := c.ShouldBindJSON(&exceptionRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err = DBUpdate[ExceptionRecord](origin, exceptionRecord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ExceptionRecord not found"})

		return
	}

	c.JSON(http.StatusOK, exceptionRecord)
}

func RouterExceptionRecordDelete(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ExceptionRecord not found"})

		return
	}

	err = DBDelete[ExceptionRecord](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)

		return
	}

	c.Status(http.StatusNoContent)
}