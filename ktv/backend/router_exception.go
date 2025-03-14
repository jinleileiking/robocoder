package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterExceptionSetup(engine *gin.Engine) {
	r := engine.Group("/v1")

	r.GET("/exceptions", RouterExceptionList)
	r.GET("/exceptions/:id", RouterExceptionGet)
	r.POST("/exceptions", RouterExceptionCreate)
	r.PUT("/exceptions/:id", RouterExceptionUpdate)
	r.DELETE("/exceptions/:id", RouterExceptionDelete)
}

func RouterExceptionList(c *gin.Context) {
	var exceptions []Exception

	exceptions, err := DBList[Exception]()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, exceptions)
}

func RouterExceptionGet(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	exception, err := DBGet[Exception](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, exception)
}

func RouterExceptionCreate(c *gin.Context) {
	var exception Exception
	if err := c.ShouldBindJSON(&exception); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := DBCreate[Exception](exception); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, exception)
}

func RouterExceptionUpdate(c *gin.Context) {
	var exception Exception

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	origin, err := DBGet[Exception](idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exception not found"})

		return
	}

	if err := c.ShouldBindJSON(&exception); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err = DBUpdate[Exception](origin, exception)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Exception not found"})

		return
	}

	c.JSON(http.StatusOK, exception)
}

func RouterExceptionDelete(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Exception not found"})

		return
	}

	err = DBDelete[Exception](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)

		return
	}

	c.Status(http.StatusNoContent)
}