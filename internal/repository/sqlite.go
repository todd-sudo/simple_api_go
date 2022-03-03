package repository

import (
	"github.com/todd-sudo/todo/internal/model"
	log "github.com/todd-sudo/todo/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewPostgresDB() (*gorm.DB, error) {
	db, err := gorm.Open(
		sqlite.Open("gorm.db"),
		&gorm.Config{},
	)
	migrations(db)
	log.Info("Migration Successfully")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{}, &model.Item{})
	if err != nil {
		log.Errorf("Migrate error: %v", err)
	}
}
