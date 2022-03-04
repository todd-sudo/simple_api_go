package service

import (
	"context"

	"github.com/mashingan/smapping"
	"github.com/todd-sudo/todo/internal/dto"
	"github.com/todd-sudo/todo/internal/model"
	"github.com/todd-sudo/todo/internal/repository"
	log "github.com/todd-sudo/todo/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

//AuthService is a contract about something that this service can do
type AuthService interface {
	VerifyCredential(ctx context.Context, email, password string) interface{}
	CreateUser(ctx context.Context, user dto.RegisterDTO) model.User
	FindByEmail(ctx context.Context, email string) model.User
	IsDuplicateEmail(ctx context.Context, email string) bool
}

type authService struct {
	ctx            context.Context
	userRepository repository.UserRepository
}

//NewAuthService creates a new instance of AuthService
func NewAuthService(ctx context.Context, userRep repository.UserRepository) AuthService {
	return &authService{
		ctx:            ctx,
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(ctx context.Context, email, password string) interface{} {
	res := service.userRepository.VerifyCredential(ctx, email)
	if v, ok := res.(model.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(ctx context.Context, user dto.RegisterDTO) model.User {
	userToCreate := model.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Errorf("Failed map %v", err)
	}
	res := service.userRepository.InsertUser(ctx, userToCreate)
	return res
}

func (service *authService) FindByEmail(ctx context.Context, email string) model.User {
	return service.userRepository.FindByEmail(ctx, email)
}

func (service *authService) IsDuplicateEmail(ctx context.Context, email string) bool {
	res := service.userRepository.IsDuplicateEmail(ctx, email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}
