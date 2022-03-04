package service

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/todd-sudo/todo/internal/dto"
	"github.com/todd-sudo/todo/internal/model"
	"github.com/todd-sudo/todo/internal/repository"
)

type Auth interface {
	VerifyCredential(ctx context.Context, email string, password string) (*model.User, error)
	CreateUser(ctx context.Context, user dto.RegisterDTO) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	IsDuplicateEmail(ctx context.Context, email string) (bool, error)
}

type JWT interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type Item interface {
	Insert(ctx context.Context, b dto.ItemCreateDTO) (*model.Item, error)
	Update(ctx context.Context, b dto.ItemUpdateDTO) (*model.Item, error)
	Delete(ctx context.Context, b model.Item) error
	All(ctx context.Context) ([]*model.Item, error)
	FindByID(ctx context.Context, itemID uint64) (*model.Item, error)
	IsAllowedToEdit(ctx context.Context, userID string, itemID uint64) (bool, error)
}

type User interface {
	Update(ctx context.Context, user dto.UserUpdateDTO) (*model.User, error)
	Profile(ctx context.Context, userID string) (*model.User, error)
}

type Service struct {
	Auth
	JWT
	Item
	User
}

func NewService(ctx context.Context, r repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(ctx, r.User),
		JWT:  NewJWTService(),
		Item: NewItemService(ctx, r.Item),
		User: NewUserService(ctx, r.User),
	}
}
