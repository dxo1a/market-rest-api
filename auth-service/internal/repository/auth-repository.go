package repository

import (
	"auth/internal/models"
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *models.User) error
	GetByPhone(phone string) (*models.User, error)
	BlacklistToken(token string, exp time.Duration) error
}

type authRepository struct {
	db          *gorm.DB
	redisClient *redis.Client
	ctx         context.Context
}

func NewAuthRepository(db *gorm.DB, redisClient *redis.Client) AuthRepository {
	return &authRepository{
		db:          db,
		redisClient: redisClient,
		ctx:         context.Background(),
	}
}

func (r *authRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *authRepository) GetByPhone(phone string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("phone = ?", phone).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("пользователь не найден")
		}
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) BlacklistToken(token string, exp time.Duration) error {
	return r.redisClient.Set(r.ctx, token, "blacklisted", exp).Err()
}
