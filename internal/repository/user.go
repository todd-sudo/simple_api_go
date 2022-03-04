package repository

import (
	"context"

	"github.com/todd-sudo/todo/internal/model"
	log "github.com/todd-sudo/todo/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type UserRepository interface {
	InsertUser(ctx context.Context, user model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user model.User) (*model.User, error)
	VerifyCredential(ctx context.Context, email string) (*model.User, error)
	IsDuplicateEmail(ctx context.Context, email string) (bool, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	ProfileUser(ctx context.Context, userID string) (*model.User, error)
}

type userConnection struct {
	ctx        context.Context
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(ctx context.Context, db *gorm.DB) UserRepository {
	return &userConnection{
		ctx:        ctx,
		connection: db,
	}
}

// Добавление пользователя
func (db *userConnection) InsertUser(ctx context.Context, user model.User) (*model.User, error) {
	tx := db.connection.WithContext(ctx)
	user.Password = hashAndSalt([]byte(user.Password))
	res := tx.Save(&user)
	if res.Error != nil {
		log.Errorf("insert user error %v", res.Error)
		return nil, res.Error
	}
	return &user, nil
}

// Обновление пользователя
func (db *userConnection) UpdateUser(ctx context.Context, user model.User) (*model.User, error) {
	tx := db.connection.WithContext(ctx)
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser model.User
		tx.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	res := tx.Save(&user)
	if res.Error != nil {
		log.Errorf("update user error %v", res.Error)
		return nil, res.Error
	}
	return &user, nil
}

func (db *userConnection) VerifyCredential(ctx context.Context, email string) (*model.User, error) {
	tx := db.connection.WithContext(ctx)
	var user model.User
	res := tx.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		log.Errorf("verify credential error %v", res.Error)
		return nil, res.Error
	}
	return &user, nil
}

// Проверка на наличие одинаковых email
func (db *userConnection) IsDuplicateEmail(ctx context.Context, email string) (bool, error) {
	var user model.User
	res := db.connection.WithContext(ctx).Where("email = ?", email).Take(&user)
	if res.Error != nil {
		log.Errorf("is duplicate email error %v", res.Error)
		return false, res.Error
	}
	return true, nil
}

// Поиск пользователя по email
func (db *userConnection) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	tx := db.connection.WithContext(ctx)
	var user model.User
	res := tx.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		log.Errorf("find by email user error %v", res.Error)
		return nil, res.Error
	}
	return &user, nil
}

// Вывод профиля пользователя
func (db *userConnection) ProfileUser(ctx context.Context, userID string) (*model.User, error) {
	tx := db.connection.WithContext(ctx)
	var user model.User
	res := tx.Preload("Items").Preload("Items.User").Find(&user, userID)
	if res.Error != nil {
		log.Errorf("profile user error %v", res.Error)
		return nil, res.Error
	}
	return &user, nil
}

// Хеширование пароля при сохранении
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Errorf("Failed to hash a password %v", err)
	}
	return string(hash)
}
