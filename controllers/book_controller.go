package controllers

import (
	"awesomeProject/database"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var book models.Books
	err = db.First(&book, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find book: " + err.Error(),
		})
		return
	}
	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()
	var book models.Books
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = db.Save(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book in database: " + err.Error(),
		})
		return
	}
	c.JSON(200, book)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()
	var books []models.Books
	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list books: " + err.Error(),
		})
		return
	}
	c.JSON(200, books)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()
	var book models.Books
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = db.Save(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update book in database: " + err.Error(),
		})
		return
	}
	c.JSON(200, book)
}
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	db := database.GetDatabase()
	err = db.Delete(&models.Books{}, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book in database: " + err.Error(),
		})
		return
	}
	c.Status(204)
}
