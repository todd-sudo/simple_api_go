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
	if err != nil {
		log.Errorf("Error connection database: %v", err)
		return nil, err
	}

	err = migrations(db)
	if err != nil {
		return nil, err
	}
	log.Info("Migration Successfully")

	return db, nil
}

func migrations(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{}, &model.Item{})
	if err != nil {
		log.Errorf("Migrate error: %v", err)
		return err
	}
	return nil
}
