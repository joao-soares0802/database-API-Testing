package controllers

import (
	"awesomeProject/database"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ShowGame(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var game models.Games
	err = db.First(&game, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find game: " + err.Error(),
		})
		return
	}
	c.JSON(200, game)
}

func CreateGame(c *gin.Context){
	db := database.GetDatabase()
	var game models.Games
	err := c.ShouldBindJSON(&game)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err=db.Save(&game).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create game in database: " + err.Error(),
		})
		return
	}
	c.JSON(200, game)
}

func ShowGames(c *gin.Context){
	db := database.GetDatabase()
	var games []models.Games
	err := db.Find(games).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list games: " + err.Error(),
		})
		return
	}
	c.JSON(200, games)
}

func UpdateGames(c *gin.Context){
	db := database.GetDatabase()
	var game models.Games
	err := c.ShouldBindJSON(&game)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err=db.Save(&game).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update game in database: " + err.Error(),
		})
		return
	}
	c.JSON(200, game)
}
func DeleteGame(c *gin.Context){
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	db := database.GetDatabase()
	err = db.Delete(&models.Games{}, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete game in database: " + err.Error(),
		})
		return
	}
	c.Status(204)
}