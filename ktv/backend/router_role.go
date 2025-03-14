package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterRoleSetup(engine *gin.Engine) {
    r := engine.Group("/v1")

    r.GET("/roles", RouterRoleList)
    r.GET("/roles/:id", RouterRoleGet)
    r.POST("/roles", RouterRoleCreate)
    r.PUT("/roles/:id", RouterRoleUpdate)
    r.DELETE("/roles/:id", RouterRoleDelete)
}

func RouterRoleList(c *gin.Context) {
	var roles []Role

	roles, err := DBList[Role]()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, roles)
}

func RouterRoleGet(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	role, err := DBGet[Role](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, role)
}

func RouterRoleCreate(c *gin.Context) {
	var role Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := DBCreate[Role](role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, role)
}

func RouterRoleUpdate(c *gin.Context) {
	var role Role

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	origin, err := DBGet[Role](idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})

		return
	}

	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err = DBUpdate[Role](origin, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Role not found"})

		return
	}

	c.JSON(http.StatusOK, role)
}

func RouterRoleDelete(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Role not found"})

		return
	}

	err = DBDelete[Role](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)

		return
	}

	c.Status(http.StatusNoContent)
}