package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func createItem(c *gin.Context) {
	var newItem Item

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	newItem.ID = nextID
	nextID++
	items = append(items, newItem)

	c.IndentedJSON(http.StatusCreated, newItem)
}

func getItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid item ID"})
		return
	}

	for _, a := range items {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func updateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid item ID"})
		return
	}

	var updatedItem Item
	if err := c.BindJSON(&updatedItem); err != nil {
		return
	}

	for i, a := range items {
		if a.ID == id {
			updatedItem.ID = a.ID
			items[i] = updatedItem
			c.IndentedJSON(http.StatusOK, updatedItem)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func deleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid item ID"})
		return
	}

	for i, a := range items {
		if a.ID == id {
			items = append(items[:i], items[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "item deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}
