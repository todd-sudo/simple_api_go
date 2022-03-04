package repository

import (
	"context"

	"github.com/todd-sudo/todo/internal/model"

	"gorm.io/gorm"
)

//BookRepository is a ....
type ItemRepository interface {
	InsertItem(ctx context.Context, b model.Item) model.Item
	UpdateItem(ctx context.Context, b model.Item) model.Item
	DeleteItem(ctx context.Context, b model.Item)
	AllItem(ctx context.Context) []model.Item
	FindItemByID(ctx context.Context, bookID uint64) model.Item
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
func (db *itemConnection) InsertItem(ctx context.Context, i model.Item) model.Item {
	tx := db.connection.WithContext(ctx)
	tx.Save(&i)
	tx.Preload("User").Find(&i)
	return i
}

// Обновление item
func (db *itemConnection) UpdateItem(ctx context.Context, i model.Item) model.Item {
	tx := db.connection.WithContext(ctx)
	tx.Save(&i)
	tx.Preload("User").Find(&i)
	return i
}

// Удаление item
func (db *itemConnection) DeleteItem(ctx context.Context, i model.Item) {
	tx := db.connection.WithContext(ctx)
	tx.Delete(&i)
}

// Поиск item по id
func (db *itemConnection) FindItemByID(ctx context.Context, itemID uint64) model.Item {
	tx := db.connection.WithContext(ctx)
	var item model.Item
	tx.Preload("User").Find(&item, itemID)
	return item
}

// Все item
func (db *itemConnection) AllItem(ctx context.Context) []model.Item {
	tx := db.connection.WithContext(ctx)
	var items []model.Item
	tx.Preload("User").Find(&items)
	return items
}
