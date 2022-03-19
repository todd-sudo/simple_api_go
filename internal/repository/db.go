package repository

import (
	"fmt"

	"github.com/todd-sudo/todo/internal/config"
	"github.com/todd-sudo/todo/internal/model"
	log "github.com/todd-sudo/todo/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *config.ConfigDatabase) (*gorm.DB, error) {

	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
	// 	"postgres",
	// 	cfg.User,
	// 	cfg.Password,
	// 	cfg.DBName,
	// 	cfg.Port,
	// 	cfg.SslMode,
	// )
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Password,
		cfg.DBName,
	)

	log.Info(dsn)
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
