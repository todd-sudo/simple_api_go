package service

import (
	"github.com/mashingan/smapping"
	"github.com/todd-sudo/todo/internal/dto"
	"github.com/todd-sudo/todo/internal/model"
	"github.com/todd-sudo/todo/internal/repository"
	log "github.com/todd-sudo/todo/pkg/logger"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) model.User
	Profile(userID string) model.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

// Обновить пользователя
func (service *userService) Update(user dto.UserUpdateDTO) model.User {
	userToUpdate := model.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Errorf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

// Профиль пользователя
func (service *userService) Profile(userID string) model.User {
	return service.userRepository.ProfileUser(userID)
}
