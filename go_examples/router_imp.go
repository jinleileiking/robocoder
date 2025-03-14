package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct{}

func RouterBookList(c *gin.Context) {
	var books []Book

	books, err := DBList[Book]()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, books)
}

func RouterBookGet(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	book, err := DBGet[Book](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, book)
}

func RouterBookCreate(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := DBCreate[Book](book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, book)
}

func RouterBookUpdate(c *gin.Context) {
	var book Book

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	origin, err := DBGet[Book](idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})

		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err = DBUpdate[Book](origin, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Book not found"})

		return
	}

	c.JSON(http.StatusOK, book)
}

func RouterBookDelete(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Book not found"})

		return
	}

	err = DBDelete[Book](idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)

		return
	}

	c.Status(http.StatusNoContent)
}
