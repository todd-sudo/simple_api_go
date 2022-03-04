package service

import (
	"context"

	"github.com/mashingan/smapping"
	"github.com/todd-sudo/todo/internal/dto"
	"github.com/todd-sudo/todo/internal/model"
	"github.com/todd-sudo/todo/internal/repository"
	log "github.com/todd-sudo/todo/pkg/logger"
)

type UserService interface {
	Update(ctx context.Context, user dto.UserUpdateDTO) (*model.User, error)
	Profile(ctx context.Context, userID string) (*model.User, error)
}

type userService struct {
	ctx            context.Context
	userRepository repository.UserRepository
}

func NewUserService(ctx context.Context, userRepo repository.UserRepository) UserService {
	return &userService{
		ctx:            ctx,
		userRepository: userRepo,
	}
}

// Обновить пользователя
func (service *userService) Update(ctx context.Context, user dto.UserUpdateDTO) (*model.User, error) {
	userToUpdate := model.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Errorf("Failed map %v:", err)
	}
	updatedUser, err := service.userRepository.UpdateUser(ctx, userToUpdate)
	if err != nil {
		log.Errorf("update user error: %v", err)
	}
	return updatedUser, nil
}

// Профиль пользователя
func (service *userService) Profile(ctx context.Context, userID string) (*model.User, error) {
	userProfile, err := service.userRepository.ProfileUser(ctx, userID)
	if err != nil {
		log.Errorf("profile user error: %v", err)
		return nil, err
	}
	return userProfile, nil
}
