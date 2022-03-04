package repository

import (
	"context"

	"github.com/todd-sudo/todo/internal/model"
	"gorm.io/gorm"
)

type User interface {
	InsertUser(ctx context.Context, user model.User) model.User
	UpdateUser(ctx context.Context, user model.User) model.User
	VerifyCredential(ctx context.Context, email string) interface{}
	IsDuplicateEmail(ctx context.Context, email string) (tx *gorm.DB)
	FindByEmail(ctx context.Context, email string) model.User
	ProfileUser(ctx context.Context, userID string) model.User
}

type Item interface {
	InsertItem(ctx context.Context, b model.Item) model.Item
	UpdateItem(ctx context.Context, b model.Item) model.Item
	DeleteItem(ctx context.Context, b model.Item)
	AllItem(ctx context.Context) []model.Item
	FindItemByID(ctx context.Context, bookID uint64) model.Item
}

type Repository struct {
	User
	Item
}

func NewRepository(ctx context.Context, db *gorm.DB) *Repository {
	return &Repository{
		User: NewUserRepository(ctx, db),
		Item: NewItemRepository(ctx, db),
	}
}
