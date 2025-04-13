package handler

import (
	"auth/internal/models"
	"auth/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(srv *service.AuthService) *AuthHandler {
	return &AuthHandler{service: srv}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input models.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	user := models.User{
		Phone:    input.Phone,
		Password: input.Password,
		Email:    input.Email,
		FullName: input.FullName,
		Address:  input.Address,
	}

	if err := h.service.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка регистрации"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Регистрация успешна"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	token, err := h.service.Login(input.Phone, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный пароль"})
		return
	}

	c.SetCookie("jwt", token, 3600*24, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Вход успешен"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	token, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен не найден"})
		return
	}

	if err := h.service.Logout(token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Ошибка при выходе",
			"message": err.Error(),
		})
		return
	}

	c.SetCookie("jwt", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Выход выполнен"})
}
