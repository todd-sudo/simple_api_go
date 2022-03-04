package service

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/todd-sudo/todo/internal/dto"
	"github.com/todd-sudo/todo/internal/model"
	"github.com/todd-sudo/todo/internal/repository"
)

type Auth interface {
	VerifyCredential(ctx context.Context, email string, password string) interface{}
	CreateUser(ctx context.Context, user dto.RegisterDTO) model.User
	FindByEmail(ctx context.Context, email string) model.User
	IsDuplicateEmail(ctx context.Context, email string) bool
}

type JWT interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type Item interface {
	Insert(ctx context.Context, b dto.ItemCreateDTO) model.Item
	Update(ctx context.Context, b dto.ItemUpdateDTO) model.Item
	Delete(ctx context.Context, b model.Item)
	All(ctx context.Context) []model.Item
	FindByID(ctx context.Context, itemID uint64) model.Item
	IsAllowedToEdit(ctx context.Context, userID string, itemID uint64) bool
}

type User interface {
	Update(ctx context.Context, user dto.UserUpdateDTO) model.User
	Profile(ctx context.Context, userID string) model.User
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
