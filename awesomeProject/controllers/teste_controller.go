package controllers

import (
	"awesomeProject/database"
	"awesomeProject/database/migrations"
	"github.com/gin-gonic/gin"
)

func ShowTestes(c *gin.Context){
	db := database.GetDatabase()
	var testes []migrations.Teste
	err := db.Find(&testes).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list books: " + err.Error(),
		})
		return
	}
	c.JSON(200, testes)
}

func CreateTeste(c *gin.Context){
	db := database.GetDatabase()
	var teste migrations.Teste
	err := c.ShouldBindJSON(&teste)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err=db.Create(&teste).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book in database: " + err.Error(),
		})
		return
	}
	c.JSON(200, teste)
}
