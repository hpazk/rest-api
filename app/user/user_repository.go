package user

import (
	"github.com/jinzhu/gorm"
)

// type UsersStorage []interface{}

// var users UsersStorage

type Repository interface {
	InsertUser(user User) (User, error)
	FindEmail(email string) *User
	FindUserByEmail(email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) InsertUser(user User) (User, error) {
	// users = append(users, user)

	// fmt.Println(users)
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindEmail(email string) *User {
	var user User

	err := r.db.First(&user, "email = ?", email).Error
	if err == nil {
		return &user
	}

	return nil
}

func (r *repository) FindUserByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
