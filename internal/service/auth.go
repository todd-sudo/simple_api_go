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
	VerifyCredential(ctx context.Context, email, password string) (*model.User, error)
	CreateUser(ctx context.Context, user dto.RegisterDTO) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	IsDuplicateEmail(ctx context.Context, email string) (bool, error)
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

func (service *authService) VerifyCredential(ctx context.Context, email, password string) (*model.User, error) {
	user, err := service.userRepository.VerifyCredential(ctx, email)

	comparedPassword := comparePassword(user.Password, []byte(password))
	if user.Email == email && comparedPassword {
		return user, nil
	}
	return nil, err
}

func (service *authService) CreateUser(ctx context.Context, user dto.RegisterDTO) (*model.User, error) {
	userToCreate := model.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Errorf("Failed map %v", err)
	}
	userModel, errU := service.userRepository.InsertUser(ctx, userToCreate)
	if errU != nil {
		log.Errorf("create user error %v", errU)
		return nil, errU
	}
	return userModel, nil
}

func (service *authService) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := service.userRepository.FindByEmail(ctx, email)
	if err != nil {
		log.Errorf("find by email user error %v", err)
		return nil, err
	}
	return user, nil
}

func (service *authService) IsDuplicateEmail(ctx context.Context, email string) (bool, error) {
	res, err := service.userRepository.IsDuplicateEmail(ctx, email)
	if err != nil {
		return false, err
	}
	return res, nil
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
