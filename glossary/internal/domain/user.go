package domain

import (
	"secondhand_glossary/internal/model"
)

type UserRepository interface {
	GetUser(u model.UserLogin) (user model.User, err error)
	GetUsers() (users []model.User, err error) // Admin
	GetUserProfile(userId uint) (user model.User, err error)
	RegisterUser(u model.UserRegister) (user model.User, err error)
	UpdateUserProfile(u model.User) (user model.User, err error)
	DeleteUser(userId uint) (err error) // Admin
}

type UserService interface {
	GetUsers() (users []model.User, err error)
	DeleteUser(userId uint) (err error)
	Register(r model.UserRegister) (user model.User, err error)
	Login(l model.UserLogin) (user model.User, err error)
	Logout(userId uint) (err error)
	UpdateProfile(u model.User) (user model.User, err error)
	GetProfileDetails(userId uint) (user model.User, err error)
}
