package service

import (
	"auth/internal/models"
	"auth/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo      repository.AuthRepository
	jwtSecret string
}

func NewAuthService(r repository.AuthRepository, secret string) *AuthService {
	return &AuthService{repo: r, jwtSecret: secret}
}

func (s *AuthService) Register(user *models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return s.repo.CreateUser(user)
}

func (s *AuthService) Login(phone, password string) (string, error) {
	user, err := s.repo.GetByPhone(phone)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(s.jwtSecret))
}

func (s *AuthService) Logout(token string) error {
	return s.repo.BlacklistToken(token, 24*time.Hour)
}
