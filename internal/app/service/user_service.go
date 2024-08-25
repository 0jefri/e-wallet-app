package service

import (
	"e-wallet-app/internal/app/repository"
	"e-wallet-app/internal/model"
	"fmt"
)

type UserService interface {
	RegisterNewUser(payload model.User) error
}

type userService struct {
	repo repository.UserRepository
}

func (e *userService) RegisterNewUser(payload model.User) error {
	if payload.Username == "" || payload.Email == "" || payload.Password == "" || payload.FirstName == "" || payload.LastName == "" || payload.PhoneNumber == "" {
		return fmt.Errorf("Username, Email, Password, FirstName, LastName, PhoneNumber Penulis is required")
	}
	err := e.repo.Create(&payload)
	if err != nil {
		return fmt.Errorf("Failed to create novel: %s", err.Error())
	}
	return nil
}
