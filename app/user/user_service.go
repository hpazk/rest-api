package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Services interface {
	CreateUser(req RequestUser) (User, error)
	CheckExistEmail(req RequestUser) error
	AuthUser(req RequestUserLogin) (User, error)
}

type services struct {
	repository Repository
}

func NewServices(repository Repository) *services {
	return &services{repository}
}

func (s *services) CreateUser(req RequestUser) (User, error) {
	user := User{}
	user.Name = req.Name
	user.Email = req.Email

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	user.Password = string(hashedPassword)

	newUser, err := s.repository.InsertUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *services) CheckExistEmail(req RequestUser) error {
	email := req.Email

	if user := s.repository.FindEmail(email); user != nil {
		return errors.New("email already registered")
	}

	return nil
}

func (s *services) AuthUser(req RequestUserLogin) (User, error) {
	email := req.Email
	password := req.Password

	user, err := s.repository.FindUserByEmail(email)
	if err != nil {
		return user, errors.New("email not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid email or password")
	}

	return user, nil
}
