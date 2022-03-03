package repository

import (
	"github.com/todd-sudo/todo/internal/model"
	"gorm.io/gorm"
)

type User interface {
	InsertUser(user model.User) model.User
	UpdateUser(user model.User) model.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) model.User
	ProfileUser(userID string) model.User
}

type Item interface {
	InsertItem(b model.Item) model.Item
	UpdateItem(b model.Item) model.Item
	DeleteItem(b model.Item)
	AllItem() []model.Item
	FindItemByID(bookID uint64) model.Item
}

type Repository struct {
	User
	Item
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
		Item: NewItemRepository(db),
	}
}
