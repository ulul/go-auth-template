package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Service interface
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
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

func (s *service) Login(loginInput LoginInput) (User, error) {
	email := loginInput.Email
	password := loginInput.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, error := s.repository.FindByEmail(email)

	if error != nil {
		return false, error
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil

}

func (s *service) SaveAvatar(ID int, avatarFileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}

	user.Avatar = avatarFileLocation
	updatedUser, err := s.repository.Update(user)

	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil

}
