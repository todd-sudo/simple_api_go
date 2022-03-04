package repository

import (
	"context"

	"github.com/todd-sudo/todo/internal/model"
	log "github.com/todd-sudo/todo/pkg/logger"
	"gorm.io/gorm"
)

//BookRepository is a ....
type ItemRepository interface {
	InsertItem(ctx context.Context, b model.Item) (*model.Item, error)
	UpdateItem(ctx context.Context, b model.Item) (*model.Item, error)
	DeleteItem(ctx context.Context, b model.Item) error
	AllItem(ctx context.Context) ([]*model.Item, error)
	FindItemByID(ctx context.Context, bookID uint64) (*model.Item, error)
}

type itemConnection struct {
	ctx        context.Context
	connection *gorm.DB
}

func NewItemRepository(ctx context.Context, dbConn *gorm.DB) ItemRepository {
	return &itemConnection{
		ctx:        ctx,
		connection: dbConn,
	}
}

// Добавление item
func (db *itemConnection) InsertItem(ctx context.Context, i model.Item) (*model.Item, error) {
	tx := db.connection.WithContext(ctx)
	tx.Save(&i)
	res := tx.Preload("User").Find(&i)
	if res.Error != nil {
		log.Errorf("inset item error: %v", res.Error)
		return nil, res.Error
	}
	return &i, nil
}

// Обновление item
func (db *itemConnection) UpdateItem(ctx context.Context, i model.Item) (*model.Item, error) {
	tx := db.connection.WithContext(ctx)
	tx.Save(&i)
	res := tx.Preload("User").Find(&i)
	if res.Error != nil {
		log.Errorf("update item error %v", res.Error)
		return nil, res.Error
	}
	return &i, nil
}

// Удаление item
func (db *itemConnection) DeleteItem(ctx context.Context, i model.Item) error {
	tx := db.connection.WithContext(ctx)
	res := tx.Delete(&i)
	if res.Error != nil {
		log.Errorf("delete item error %v", res.Error)
		return res.Error
	}
	return nil
}

// Поиск item по id
func (db *itemConnection) FindItemByID(ctx context.Context, itemID uint64) (*model.Item, error) {
	tx := db.connection.WithContext(ctx)
	var item model.Item
	res := tx.Preload("User").Find(&item, itemID)
	if res.Error != nil {
		log.Errorf("find item by id error %v", res.Error)
		return nil, res.Error
	}
	return &item, nil
}

// Все item
func (db *itemConnection) AllItem(ctx context.Context) ([]*model.Item, error) {
	tx := db.connection.WithContext(ctx)
	var items []*model.Item
	res := tx.Preload("User").Find(&items)
	if res.Error != nil {
		log.Errorf("get all items error %v", res.Error)
		return nil, res.Error
	}
	return items, nil
}
