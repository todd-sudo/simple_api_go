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
	InsertUser(ctx context.Context, user model.User) model.User
	UpdateUser(ctx context.Context, user model.User) model.User
	VerifyCredential(ctx context.Context, email string) interface{}
	IsDuplicateEmail(ctx context.Context, email string) (tx *gorm.DB)
	FindByEmail(ctx context.Context, email string) model.User
	ProfileUser(ctx context.Context, userID string) model.User
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
func (db *userConnection) InsertUser(ctx context.Context, user model.User) model.User {
	tx := db.connection.WithContext(ctx)
	user.Password = hashAndSalt([]byte(user.Password))
	tx.Save(&user)
	return user
}

// Обновление пользователя
func (db *userConnection) UpdateUser(ctx context.Context, user model.User) model.User {
	tx := db.connection.WithContext(ctx)
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser model.User
		tx.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(ctx context.Context, email string) interface{} {
	tx := db.connection.WithContext(ctx)
	var user model.User
	res := tx.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

// Проверка на наличие одинаковых email
func (db *userConnection) IsDuplicateEmail(ctx context.Context, email string) (tx *gorm.DB) {
	var user model.User
	return db.connection.WithContext(ctx).Where("email = ?", email).Take(&user)
}

// Поиск пользователя по email
func (db *userConnection) FindByEmail(ctx context.Context, email string) model.User {
	tx := db.connection.WithContext(ctx)
	var user model.User
	tx.Where("email = ?", email).Take(&user)
	return user
}

// Вывод профиля пользователя
func (db *userConnection) ProfileUser(ctx context.Context, userID string) model.User {
	tx := db.connection.WithContext(ctx)
	var user model.User
	tx.Preload("Items").Preload("Items.User").Find(&user, userID)
	return user
}

// Хеширование пароля при сохранении
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Errorf("Failed to hash a password %v", err)
	}
	return string(hash)
}
