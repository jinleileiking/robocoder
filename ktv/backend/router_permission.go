package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterPermissionSetup(engine *gin.Engine) {
    r := engine.Group("/v1")

    r.GET("/permissions", RouterPermissionList)
    r.GET("/permissions/:id", RouterPermissionGet)
    r.POST("/permissions", RouterPermissionCreate)
    r.PUT("/permissions/:id", RouterPermissionUpdate)
    r.DELETE("/permissions/:id", RouterPermissionDelete)
}

func RouterPermissionList(c *gin.Context) {
	var permissions []Permission

	permissions, err := DBList[Permission]()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, permissions)
}

func RouterPermissionGet(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	permission, err := DBGet[Permission](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, permission)
}

func RouterPermissionCreate(c *gin.Context) {
	var permission Permission
	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := DBCreate[Permission](permission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, permission)
}

func RouterPermissionUpdate(c *gin.Context) {
	var permission Permission

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	origin, err := DBGet[Permission](idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})

		return
	}

	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err = DBUpdate[Permission](origin, permission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Permission not found"})

		return
	}

	c.JSON(http.StatusOK, permission)
}

func RouterPermissionDelete(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Permission not found"})

		return
	}

	err = DBDelete[Permission](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)

		return
	}

	c.Status(http.StatusNoContent)
}