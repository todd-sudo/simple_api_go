package repository

import (
	"context"

	"github.com/todd-sudo/todo/internal/model"
	"gorm.io/gorm"
)

type User interface {
	InsertUser(ctx context.Context, user model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user model.User) (*model.User, error)
	VerifyCredential(ctx context.Context, email string) (*model.User, error)
	IsDuplicateEmail(ctx context.Context, email string) (bool, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	ProfileUser(ctx context.Context, userID string) (*model.User, error)
}

type Item interface {
	InsertItem(ctx context.Context, b model.Item) (*model.Item, error)
	UpdateItem(ctx context.Context, b model.Item) (*model.Item, error)
	DeleteItem(ctx context.Context, b model.Item) error
	AllItem(ctx context.Context) ([]*model.Item, error)
	FindItemByID(ctx context.Context, bookID uint64) (*model.Item, error)
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
