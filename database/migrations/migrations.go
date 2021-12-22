package migrations

import (
	"awesomeProject/models"
	"gorm.io/gorm"
)

type Teste struct {
	Testando int `json:"testando"`
}

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Books{}, Teste{}, models.Games{})
}
