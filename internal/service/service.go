package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/todd-sudo/todo/internal/dto"
	"github.com/todd-sudo/todo/internal/model"
	"github.com/todd-sudo/todo/internal/repository"
)

type Auth interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) model.User
	FindByEmail(email string) model.User
	IsDuplicateEmail(email string) bool
}

type JWT interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type Item interface {
	Insert(b dto.ItemCreateDTO) model.Item
	Update(b dto.ItemUpdateDTO) model.Item
	Delete(b model.Item)
	All() []model.Item
	FindByID(itemID uint64) model.Item
	IsAllowedToEdit(userID string, itemID uint64) bool
}

type User interface {
	Update(user dto.UserUpdateDTO) model.User
	Profile(userID string) model.User
}

type Service struct {
	Auth
	JWT
	Item
	User
}

func NewService(r repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(r.User),
		JWT:  NewJWTService(),
		Item: NewItemService(r.Item),
		User: NewUserService(r.User),
	}
}
