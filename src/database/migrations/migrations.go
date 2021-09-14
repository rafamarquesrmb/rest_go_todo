package migrations

import (
	"github.com/rafamarquesrmb/rest_go_todo/src/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.Task{})
}
