package services

import (
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Siginup(emial string, password string) error
	Login(emial string, psssword string) (*string, error)
}

type authService struct {
	repository repositories.IAuthRepository
}

func NewAuthService(repository repositories.IAuthRepository) IAuthService {
	return &authService{repository: repository}
}

func (s *authService) Siginup(emial string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := models.User{
		Email: emial,
		Password: string(hashedPassword),
	}
	return  s.repository.CreateUser(newUser)
}

func (s *authService) Login(email string, password string) (*string, error) {
	foundUser, err := s.repository.FindUser(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &foundUser.Email, nil
}