package main

import (
	"auth/config"
	"auth/internal/handler"
	"auth/internal/repository"
	"auth/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db := repository.NewPostgres(cfg.Postgres)
	redis := repository.NewRedis(cfg.Redis)

	repo := repository.NewAuthRepository(db, redis)
	srv := service.NewAuthService(repo, cfg.JWTSecret)
	h := handler.NewAuthHandler(srv)

	r := gin.Default()
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.POST("/logout", h.Logout)

	if err := r.Run(":8081"); err != nil {
		log.Fatal("Server failed: ", err)
	}
}
