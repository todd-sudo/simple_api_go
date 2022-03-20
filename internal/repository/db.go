package repository

import (
	"fmt"

	"github.com/todd-sudo/todo/internal/config"
	"github.com/todd-sudo/todo/internal/model"
	log "github.com/todd-sudo/todo/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connection database: %v", err)
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
