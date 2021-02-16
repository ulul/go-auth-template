package user

import (
	"golang.org/x/crypto/bcrypt"
)

// Service interface
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

// NewService function
func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	user := User{}
	if err != nil {
		return user, err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Ocupation = input.Ocupation
	user.Password = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
