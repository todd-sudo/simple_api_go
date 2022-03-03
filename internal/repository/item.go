package repository

import (
	"github.com/todd-sudo/todo/internal/model"

	"gorm.io/gorm"
)

//BookRepository is a ....
type ItemRepository interface {
	InsertItem(b model.Item) model.Item
	UpdateItem(b model.Item) model.Item
	DeleteItem(b model.Item)
	AllItem() []model.Item
	FindItemByID(bookID uint64) model.Item
}

type itemConnection struct {
	connection *gorm.DB
}

func NewItemRepository(dbConn *gorm.DB) ItemRepository {
	return &itemConnection{
		connection: dbConn,
	}
}

// Добавление item
func (db *itemConnection) InsertItem(i model.Item) model.Item {
	db.connection.Save(&i)
	db.connection.Preload("User").Find(&i)
	return i
}

// Обновление item
func (db *itemConnection) UpdateItem(i model.Item) model.Item {
	db.connection.Save(&i)
	db.connection.Preload("User").Find(&i)
	return i
}

// Удаление item
func (db *itemConnection) DeleteItem(i model.Item) {
	db.connection.Delete(&i)
}

// Поиск item по id
func (db *itemConnection) FindItemByID(itemID uint64) model.Item {
	var item model.Item
	db.connection.Preload("User").Find(&item, itemID)
	return item
}

// Все item
func (db *itemConnection) AllItem() []model.Item {
	var items []model.Item
	db.connection.Preload("User").Find(&items)
	return items
}
