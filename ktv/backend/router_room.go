package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterRoomSetup(engine *gin.Engine) {
    r := engine.Group("/v1")

    r.GET("/rooms", RouterRoomList)
    r.GET("/rooms/:id", RouterRoomGet)
    r.POST("/rooms", RouterRoomCreate)
    r.PUT("/rooms/:id", RouterRoomUpdate)
    r.DELETE("/rooms/:id", RouterRoomDelete)
}

func RouterRoomList(c *gin.Context) {
	var rooms []Room

	rooms, err := DBList[Room]()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, rooms)
}

func RouterRoomGet(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	room, err := DBGet[Room](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, room)
}

func RouterRoomCreate(c *gin.Context) {
	var room Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := DBCreate[Room](room); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, room)
}

func RouterRoomUpdate(c *gin.Context) {
	var room Room

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	origin, err := DBGet[Room](idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})

		return
	}

	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err = DBUpdate[Room](origin, room)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Room not found"})

		return
	}

	c.JSON(http.StatusOK, room)
}

func RouterRoomDelete(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Room not found"})

		return
	}

	err = DBDelete[Room](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)

		return
	}

	c.Status(http.StatusNoContent)
}